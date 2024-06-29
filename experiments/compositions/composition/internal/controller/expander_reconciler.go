// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"golang.org/x/time/rate"
	compositionv1alpha1 "google.com/composition/api/v1alpha1"
	"google.com/composition/pkg/applier"
	"google.com/composition/pkg/containerexecutor/jobcontainerexecutor"
	pb "google.com/composition/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// ExpanderReconciler reconciles a expander object
type ExpanderReconciler struct {
	client.Client
	Scheme                    *runtime.Scheme
	Recorder                  record.EventRecorder
	RESTMapper                meta.RESTMapper
	Config                    *rest.Config
	Dynamic                   *dynamic.DynamicClient
	InputGVK                  schema.GroupVersionKind
	InputGVR                  schema.GroupVersionResource
	Composition               types.NamespacedName
	ComopsitionChangedWatcher chan event.GenericEvent
}

type EvaluateWaitError struct {
	msg string
}

func (e *EvaluateWaitError) Error() string { return e.msg }

var planGVK schema.GroupVersionKind = schema.GroupVersionKind{
	Group:   "composition.google.com",
	Version: "v1alpha1",
	Kind:    "Plan",
}

var contextGVK schema.GroupVersionKind = schema.GroupVersionKind{
	Group:   "composition.google.com",
	Version: "v1alpha1",
	Kind:    "Context",
}

func (r *ExpanderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger = logger.WithName(r.Composition.Name).WithName(r.InputGVK.Group)

	logger.Info("Got Input API for expansion", "request", req)

	inputcr := unstructured.Unstructured{}
	inputcr.SetGroupVersionKind(r.InputGVK)
	if err := r.Get(ctx, req.NamespacedName, &inputcr); err != nil {
		logger.Error(err, "unable to fetch Input API Object")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	loggerCR := logger.WithName(inputcr.GetName())

	// Grab the latest composition
	// TODO(barni@) - Decide how we want the latest composition changes are to be applied.
	var compositionCR compositionv1alpha1.Composition
	if err := r.Client.Get(ctx, r.Composition, &compositionCR); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Unable to fetch Composition Object")
			return ctrl.Result{}, err
		}
	}

	// Associate a plan object with this input CR
	var plancr compositionv1alpha1.Plan
	planNN := types.NamespacedName{Name: r.InputGVR.Resource + "-" + inputcr.GetName(), Namespace: inputcr.GetNamespace()}
	if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Unable to fetch Plan Object")
			return ctrl.Result{}, err
		}
		// create a plan object
		plancr = compositionv1alpha1.Plan{
			ObjectMeta: metav1.ObjectMeta{
				Name:      planNN.Name,
				Namespace: planNN.Namespace,
			},
			Spec: compositionv1alpha1.PlanSpec{
				Stages: map[string]compositionv1alpha1.Stage{},
			},
			Status: compositionv1alpha1.PlanStatus{
				Stages: map[string]*compositionv1alpha1.StageStatus{},
			},
		}
		if err := ctrl.SetControllerReference(&inputcr, &plancr, r.Scheme); err != nil {
			logger.Error(err, "Unable to set controller reference for Plan Object")
			return ctrl.Result{}, err
		}
		if err := r.Client.Create(ctx, &plancr); err != nil {
			logger.Error(err, "Unable to create Plan Object")
			return ctrl.Result{}, err
		}

		// Read after create to get the UID
		if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
			logger.Error(err, "Unable to fetch Plan Object")
			return ctrl.Result{}, err
		}
	}
	// Since applylib doesnt support multiple apply batches we are
	// tracking all old objects and accumulating them each apply.
	oldAppliers := []*applier.Applier{}

	// Create a new status for comparison later
	newStatus := compositionv1alpha1.PlanStatus{
		Stages: map[string]*compositionv1alpha1.StageStatus{},
		// TODO: Accumulates LastPruned.
		// Ideally we need to reset if input/composition gen changes
		LastPruned: plancr.Status.LastPruned,
	}

	// Try updating status before returning
	defer func() {
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			nn := types.NamespacedName{Namespace: plancr.Namespace, Name: plancr.Name}
			err := r.Client.Get(ctx, nn, &plancr)
			if err != nil {
				return err
			}
			plancr.Status = *newStatus.DeepCopy()
			return r.Client.Status().Update(ctx, &plancr)
		})
		if err != nil {
			logger.Error(err, "unable to update Plan status")
		}
	}()
	expanderDebugLogsEnabled := false
	_, exist := inputcr.GetAnnotations()["composition-expander-debug-logs"]
	if exist {
		expanderDebugLogsEnabled = true
		logger.Info("annotation'composition-expander-debug-logs' is turned on.")
	}
	logger.Info("annotation'composition-expander-debug-logs' is turned off.")

	// ------------------- EVALUATION SECTION -----------------------
	stagesEvaluated := []string{}
	values := map[string]interface{}{}
	planUpdated := false
	waitRequested := false
	for i, expander := range compositionCR.Spec.Expanders {
		logger = loggerCR.WithName(expander.Name).WithName("Expand")

		uri, ev, reason, err := r.getExpanderValue(ctx, expander.Version, expander.Type)
		if err != nil {
			logger.Error(err, "Error getting expander version", "expander", expander.Type,
				"version", expander.Version, "reason", reason)
			newStatus.AppendErrorCondition(expander.Name, err.Error(), reason)
			return ctrl.Result{}, err
		}

		logger.Info("Got valid expander uri", "uri", uri)

		if expanderDebugLogsEnabled {
			r.Recorder.Event(&inputcr, "Normal", fmt.Sprintf("Running expander stage %d: %s", i, expander.Name), expanderDebugLog(&inputcr))
		}
		if ev.Spec.Type == compositionv1alpha1.ExpanderTypeJob {
			reason, err := r.runJob(ctx, logger, &inputcr, expander.Name, planNN.Name, uri, ev.Spec.ImageRegistry)
			if err != nil {
				newStatus.AppendErrorCondition(expander.Name, err.Error(), reason)
				return ctrl.Result{}, err
			}
		} else {
			newValues, updated, reason, err := r.evaluateAndSavePlan(ctx, logger, &inputcr, values, expander, planNN, ev, uri, expanderDebugLogsEnabled)
			_, iswaitErr := err.(*EvaluateWaitError)
			if iswaitErr {
				newStatus.AppendWaitingCondition(expander.Name, err.Error(), reason)
				// Subsume the error
				waitRequested = true
				// continue to apply phase
				break
			}
			if err != nil {
				newStatus.AppendErrorCondition(expander.Name, err.Error(), reason)
				// Skip apply phase and return
				return ctrl.Result{}, err
			}
			planUpdated = updated
			values = newValues
		}
		stagesEvaluated = append(stagesEvaluated, expander.Name)
		// Inject plan.Ready Condition with list of expanders
		newStatus.ClearCondition(compositionv1alpha1.Ready)
		message := fmt.Sprintf("Evaluated stages: %s", strings.Join(stagesEvaluated, ", "))
		newStatus.AppendCondition(compositionv1alpha1.Ready, metav1.ConditionFalse, message, "EvaluationPending")
		if expanderDebugLogsEnabled {
			r.Recorder.Event(&inputcr, "Normal", fmt.Sprintf("Evaluated expander stage %d: %s", i, expander.Name), expanderDebugLog(&inputcr))
		}
	}

	// ------------------- APPLIER SECTION -----------------------
	stagesApplied := []string{}

	// Re-read the Plan CR to load the expanded manifests
	oldGeneration := plancr.GetGeneration()
	if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
		logger.Error(err, "unable to read Plan CR")
		return ctrl.Result{}, err
	}
	if planUpdated && oldGeneration == plancr.GetGeneration() {
		logger.Info("Did not get the latest Plan CR. Will retry.", "generation", oldGeneration)
		return ctrl.Result{RequeueAfter: 5 * time.Second}, nil
	}
	for index, expander := range stagesEvaluated {
		stage, ok := plancr.Spec.Stages[expander]
		if !ok {
			err := fmt.Errorf("plancr.spec.stages[%s] not found !!", stage)
			logger.Error(err, "error applying stage", "stage", stage)
			// This is not expected since we just processed the stage above
			// We dont want to return error. Lets retry again in 20 secs.
			return ctrl.Result{RequeueAfter: 20 * time.Second}, nil
		}
		if stage.Values != "" {
			// This looks like an Getter stage. skip it
			continue
		}
		if stage.Manifest == "" {
			err := fmt.Errorf("plancr.spec.stages[%s].manifest is empty !!", stage)
			logger.Error(err, "error applying stage", "stage", stage)
			return ctrl.Result{}, err
		}

		// Create Applier and wait for the Applier to complete
		logger = loggerCR.WithName(expander).WithName("Apply")

		ac := applier.ApplierClient{
			Client:     r.Client,
			Dynamic:    r.Dynamic,
			RESTMapper: r.RESTMapper,
		}
		namespace := ""
		if compositionCR.Spec.NamespaceMode != compositionv1alpha1.NamespaceModeExplicit {
			namespace = inputcr.GetNamespace()
		}
		applier := applier.NewApplier(ctx, logger, ac, expander, namespace, r.InputGVR.Resource, &plancr)
		err := applier.Load() // Load Manifests
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ApplyFailed", fmt.Sprintf("error loading manifests for expander, name: %s", expander))
			logger.Error(err, "Unable to Load manifests for applying")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander, err.Error(), "FailedLoadingManifestsFromPlan")
			return ctrl.Result{}, err
		}
		logger.Info("Successfully loaded manifests for applying")
		if newStatus.Stages == nil {
			newStatus.Stages = map[string]*compositionv1alpha1.StageStatus{}
		}
		newStatus.Stages[expander] = &compositionv1alpha1.StageStatus{ResourceCount: applier.Count()}

		// Prune only for the last expander section
		prune := false
		if index == len(compositionCR.Spec.Expanders)-1 {
			prune = true
		}

		err = applier.Apply(oldAppliers, prune) // Apply Manifests
		applier.UpdateStageStatus(&newStatus)
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ApplyFailed",
				fmt.Sprintf("error applying manifests for expander, name: %s", expander))
			logger.Error(err, "Unable to apply manifests")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander, err.Error(), "FailedApplyingManifests")
			return ctrl.Result{}, err
		}

		logger.Info("Successfully applied manifests")
		r.Recorder.Event(&inputcr, "Normal", "ResourcesApplied", fmt.Sprintf("All expanded resources were applied. name: %s", expander))

		success, err := applier.Wait()
		if err != nil {
			r.Recorder.Event(&inputcr, "Warning", "ReconcileFailed", fmt.Sprintf("Failed waiting for resources to be reconciled. name: %s", expander))
			logger.Error(err, "Failed waiting for applied resources to reconcile")
			// Inject plan.ERROR Condition
			newStatus.AppendErrorCondition(expander, err.Error(), "FailedWaitingForAppliedResources")
			return ctrl.Result{}, err
		}

		oldAppliers = append(oldAppliers, applier)

		if !success {
			r.Recorder.Event(&inputcr, "Warning", "ReconcileFailed", fmt.Sprintf("Some resources are not healthy. name: %s", expander))
			logger.Info("Applied succesfully but some resources did not become healthy")
			// Inject plan.Waiting Condition
			newStatus.AppendWaitingCondition(expander, "Not all resources are healthy", "WaitingForAppliedResources")
			return ctrl.Result{}, fmt.Errorf("Some applied resources are not healthy")
		}
		logger.Info("Applied resources successfully.")

		stagesApplied = append(stagesApplied, expander)
		r.Recorder.Event(&inputcr, "Normal", "ResourcesReconciled", fmt.Sprintf("All applied resources were reconciled. name: %s", expander))
		// Inject plan.Ready Condition with list of expanders
		newStatus.ClearCondition(compositionv1alpha1.Ready)
		message := fmt.Sprintf("Evaluated stages: %s \n Applied stages: %s", strings.Join(stagesEvaluated, ", "), strings.Join(stagesApplied, ", "))
		newStatus.AppendCondition(compositionv1alpha1.Ready, metav1.ConditionFalse, message, "PendingStages")

		if expanderDebugLogsEnabled {
			r.Recorder.Event(&inputcr, "Normal", fmt.Sprintf("Finished expander stage %d: %s", index, expander), expanderDebugLog(&inputcr)+fmt.Sprintf("---resource count: %d", applier.Count()))
			for i, resourceStatus := range newStatus.Stages[expander].LastApplied {
				logger.Info("Resource %d, Resource name: %s, resource namespace: %s, resource gvk: %s %s %s, resource status: %s",
					i, resourceStatus.Name, resourceStatus.Namespace, resourceStatus.Group, resourceStatus.Version, resourceStatus.Kind, resourceStatus.Health)
			}
		}
	}

	// Inject plan.Ready Condition with list of expanders
	newStatus.ClearCondition(compositionv1alpha1.Ready)
	message := fmt.Sprintf("Evaluated and Applied stages: %s", strings.Join(stagesApplied, ", "))
	newStatus.AppendCondition(compositionv1alpha1.Ready, metav1.ConditionTrue, message, "ProcessedAllStages")
	newStatus.InputGeneration = inputcr.GetGeneration()
	newStatus.Generation = plancr.GetGeneration()
	newStatus.CompositionGeneration = compositionCR.GetGeneration()
	newStatus.CompositionUID = compositionCR.GetUID()
	if waitRequested {
		return ctrl.Result{RequeueAfter: time.Second * 5}, nil
	}
	return ctrl.Result{}, nil
}

func (r *ExpanderReconciler) getExpanderValue(
	ctx context.Context, inputExpanderVersion string, expanderType string,
) (string, *compositionv1alpha1.ExpanderVersion, string, error) {
	logger := log.FromContext(ctx)

	value := ""
	var ev compositionv1alpha1.ExpanderVersion
	err := r.Client.Get(ctx,
		types.NamespacedName{
			Name:      "composition-" + expanderType,
			Namespace: "composition-system"},
		&ev)

	if err != nil {
		// The CR should be created before the specified expander can be used.
		logger.Error(err, "Failed to get the ExpanderVersionCR")
		if apierrors.IsNotFound(err) {
			return value, nil, "MissingExpanderCR", err
		} else {
			return value, nil, "ErrorGettingExpanderVersionCR", err
		}
	}

	if ev.Status.VersionMap == nil {
		return value, nil, "ErrorEmptyVersionMap", fmt.Errorf("ExpanderVersion .status.versionMap is empty")
	}

	logger.Info("input expander version", "current", inputExpanderVersion)
	value, ok := ev.Status.VersionMap[inputExpanderVersion]
	if !ok {
		return value, nil, "VersionNotFound", fmt.Errorf("%s version not found", inputExpanderVersion)
	}
	return value, &ev, "", nil
}

func (r *ExpanderReconciler) runJob(ctx context.Context, logger logr.Logger,
	cr *unstructured.Unstructured, expanderName, planName, image, registry string) (string, error) {
	jf := jobcontainerexecutor.NewJobFactory(ctx, logger, r.Client, r.InputGVK, r.InputGVR,
		r.Composition.Name, r.Composition.Namespace,
		cr, expanderName, image, planName, registry)

	// Create Expander Job and wait for the Job to complete
	logger.Info("Creating expander job")
	err := jf.Create()
	defer jf.CleanUp()
	if err != nil {
		// Reference for event API: Event(object runtime.Object, eventtype, reason, message string)
		r.Recorder.Event(cr, "Warning", "ExpansionFailed", fmt.Sprintf("Error creating job for name: %s", expanderName))
		logger.Error(err, "Unable to create expander job")
		return "ExpanderJobCreationFailed", err
	}
	r.Recorder.Event(cr, "Normal", "ExpansionStarted", fmt.Sprintf("Job Created for name: %s", expanderName))
	logger.Info("Successfully created expander job")

	success, err := jf.Wait()
	if err != nil {
		r.Recorder.Event(cr, "Warning", "ExpansionFailed", fmt.Sprintf("Error waiting for Job for name: %s", expanderName))
		logger.Error(err, "Failed waiting for expander job")
		return "WaitingForExpanderFailed", err
	}

	if !success {
		r.Recorder.Event(cr, "Warning", "ExpansionFailed", fmt.Sprintf("Job failed for name: %s", expanderName))
		logger.Info("Expander job completed but Failed")
		return "ExpansionFailed", fmt.Errorf("Expander Job Failed")
	}
	r.Recorder.Event(cr, "Normal", "ExpansionSucceded", fmt.Sprintf("Job succeded for name: %s", expanderName))
	logger.Info("Expander job Completed successfully")
	return "", nil
}

func (r *ExpanderReconciler) evaluateAndSavePlan(ctx context.Context, logger logr.Logger,
	cr *unstructured.Unstructured, values map[string]interface{}, expander compositionv1alpha1.Expander,
	planNN types.NamespacedName, ev *compositionv1alpha1.ExpanderVersion, grpcService string, expanderDebugLogEnabled bool) (map[string]interface{}, bool, string, error) {
	// Set up a connection to the server.
	updated := false

	conn, err := grpc.Dial(grpcService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error(err, "grpc dial failed: "+grpcService)
		return values, updated, "GRPCConnError", err
	}

	// read context in cr.namespace
	var contextBytes []byte
	contextcr := unstructured.Unstructured{}
	contextcr.SetGroupVersionKind(contextGVK)
	contextNN := types.NamespacedName{Namespace: cr.GetNamespace(), Name: "context"}
	if err := r.Get(ctx, contextNN, &contextcr); err != nil {
		logger.Error(err, "unable to fetch Context CR", "context", contextNN)
		if !apierrors.IsNotFound(err) {
			return values, updated, "ErrorGettingContext", err
		}
		// If context doesnt exist ignore it. If a composition uses context,
		//  it will fail evaluation
	} else {
		contextBytes, err = json.Marshal(contextcr.Object)
		if err != nil {
			logger.Error(err, "failed to marshal Context Object")
			return values, updated, "MarshallContextFailed", err
		}
	}

	// marshall facade cr
	facadeBytes, err := json.Marshal(cr.Object)
	if err != nil {
		logger.Error(err, "failed to marshall Facade Object")
		return values, updated, "MarshallFacadeFailed", err
	}

	// marshall expander config
	// read Expander config from  in cr.namespace
	var configBytes []byte
	if expander.Reference != nil {
		expanderconfigcr := unstructured.Unstructured{}
		expanderconfigcr.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   ev.Spec.Config.Group,
			Version: ev.Spec.Config.Version,
			Kind:    ev.Spec.Config.Kind,
		})
		expanderconfigNN := types.NamespacedName{Namespace: expander.Reference.Namespace, Name: expander.Reference.Name}
		if err := r.Get(ctx, expanderconfigNN, &expanderconfigcr); err != nil {
			logger.Error(err, "unable to fetch ExpanderConfig CR", "expander config", expanderconfigNN)
			return values, updated, "GetExpanderConfigFailed", err
		}
		configBytes, err = json.Marshal(expanderconfigcr.Object)
		if err != nil {
			logger.Error(err, "failed to marshal ExpanderConfig Object")
			return values, updated, "MarshallExpanderConfigFailed", err
		}
	} else {
		// TODO check if json.Marshall is escaping quotes
		// Also causes > to be replaced unicode 'if loop.index \u003e 1'
		err = nil
		//configBytes, err = json.Marshal(expander.Template)
		configBytes = []byte(expander.Template)
		if err != nil {
			logger.Error(err, "failed to marshall Expander template")
			return values, updated, "MarshallExpanderTemplateFailed", err
		}
	}

	// marshall getter values
	valuesBytes, err := json.Marshal(values)
	if err != nil {
		logger.Error(err, "failed to marshall Getter Values")
		return values, updated, "MarshallValuesFailed", err
	}
	evaluateRequest := &pb.EvaluateRequest{
		Config:   configBytes,
		Context:  contextBytes,
		Facade:   facadeBytes,
		Resource: r.InputGVR.Resource,
		Value:    valuesBytes,
	}
	expanderClient := pb.NewExpanderClient(conn)
	if expanderDebugLogEnabled {
		logger.Info(expanderDebugLog(cr) + fmt.Sprintf("---sending expander request: %v", evaluateRequest))
	}
	result, err := expanderClient.Evaluate(ctx, evaluateRequest)
	if err != nil {
		logger.Error(err, "expander.Evaluate() Failed", "expander", expander.Name)
		return values, updated, "EvaluateError", err
	}
	if result.Status == pb.Status_EVALUATE_WAIT {
		logger.Error(nil, "expander.Evaluate() returned WAIT", "expander", expander.Name, "status", result.Status, "msg", result.Error.Message)
		err = &EvaluateWaitError{msg: fmt.Sprintf("Expander returned WAIT: %s", result.Error.Message)}
		return values, updated, "EvaluateStatusWait", err
	}
	if result.Status != pb.Status_SUCCESS {
		logger.Error(nil, "expander.Evaluate() Status is not Success", "expander", expander.Name, "status", result.Status)
		err = fmt.Errorf("Evaluate Failed: %s", result.Error.Message)
		return values, updated, "EvaluateStatusFailed", err
	}
	if expanderDebugLogEnabled {
		logger.Info(expanderDebugLog(cr) + fmt.Sprintf("---sent expander request: %v, received results: %v", evaluateRequest, result))
		r.Recorder.Event(cr, "Normal", fmt.Sprintf("Expander stage %s evaluation completed", expander.Name), expanderDebugLog(cr)+fmt.Sprintf("---request: %v, result: %v", evaluateRequest, result))
	}

	// Write to Plan object
	// Re-read the Plan CR to load the expanded manifests
	plancr := compositionv1alpha1.Plan{}
	if err := r.Client.Get(ctx, planNN, &plancr); err != nil {
		logger.Error(err, "unable to read Plan CR", "plan", planNN)
		return values, updated, "GetPlanFailed", err
	}

	if plancr.Spec.Stages == nil {
		plancr.Spec.Stages = map[string]compositionv1alpha1.Stage{}
		updated = true
	}
	if _, ok := plancr.Spec.Stages[expander.Name]; !ok {
		updated = true
	}

	if result.Type == pb.ResultType_MANIFESTS {
		err = nil
		//s, err := strconv.Unquote(string(result.Manifests))
		s := string(result.Manifests)
		if err != nil {
			logger.Error(err, "unable to unquote grpc response")
			return values, updated, "UnquoteResponseFailed", err
		}
		if s != plancr.Spec.Stages[expander.Name].Manifest {
			plancr.Spec.Stages[expander.Name] = compositionv1alpha1.Stage{
				Manifest: s,
			}
			updated = true
		}
	} else {
		//s, err := strconv.Unquote(string(result.Values))
		//if err != nil {
		//	logger.Error(err, "unable to unquote grpc response")
		//	return values, updated, "UnquoteResponseFailed", err
		//}
		s := string(result.Values)
		if s != plancr.Spec.Stages[expander.Name].Values {
			plancr.Spec.Stages[expander.Name] = compositionv1alpha1.Stage{
				Values: s,
			}
			updated = true
		}

		// Grab values for donwstream stages
		stageValues := map[string]any{}
		err = json.Unmarshal([]byte(plancr.Spec.Stages[expander.Name].Values), &stageValues)
		if err != nil {
			logger.Error(err, "Failed unmarshalling response.Values field")
			return values, updated, "UnmarshallValuesFailed", err
		}
		for k := range stageValues {
			_, ok := values[k]
			if ok {
				err := fmt.Errorf("values[%s] already exists from one of the previous stages.", k)
				logger.Error(err, "Duplicate Value Key")
				return values, updated, "DuplicateValueKey", err
			}
			values[k] = stageValues[k]
		}
	}

	err = r.Client.Update(ctx, &plancr)
	if err != nil {
		logger.Error(err, "error updating plan", "plan", planNN)
		return values, updated, "UpdatePlanFailed", err
	}

	return values, updated, "", nil
}

func expanderDebugLog(cr *unstructured.Unstructured) string {
	// TODO(@xiaoweim): add the UID in the future if possible
	return fmt.Sprintf("expanderDebugLog---%s/%s/%s---version: %d", cr.GetKind(), cr.GetNamespace(), cr.GetName(), cr.GetGeneration())
}

func (r *ExpanderReconciler) enqueueAllFromGVK(ctx context.Context, _ client.Object) []reconcile.Request {
	logger := log.FromContext(ctx)
	logger.Info("Got notification of changed CRD")
	inputcrList := &unstructured.UnstructuredList{}
	inputcrList.SetGroupVersionKind(r.InputGVK)
	if err := r.List(ctx, inputcrList); err != nil {
		logger.Error(err, "unable to fetch Input API Objects")
		return nil
	}
	if len(inputcrList.Items) == 0 {
		return nil
	}
	var reqs []reconcile.Request
	// TODO: If there are lots of objects, this will result in very many reconciles. Have not tested to see how the
	// queue copes with this. If it becomes a problem, this will need a rethink.
	for _, inputcr := range inputcrList.Items {
		nn := types.NamespacedName{Name: inputcr.GetName(), Namespace: inputcr.GetNamespace()}
		reqs = append(reqs, reconcile.Request{NamespacedName: nn})
	}
	return reqs
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExpanderReconciler) SetupWithManager(mgr ctrl.Manager, cr *unstructured.Unstructured) error {
	var err error
	// TODO(barni@): Can we setup dynamic controller at main.go for CompositionReconciler instead of 1 per ExpanderReconciler
	r.Dynamic, err = dynamic.NewForConfig(r.Config)
	if err != nil {
		return fmt.Errorf("error building dynamic client: %w", err)
	}

	ratelimiter := workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(5*time.Millisecond, 120*time.Second),
		// 40 qps, 400 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(40), 400)},
	)

	return ctrl.NewControllerManagedBy(mgr).
		For(cr).
		WatchesRawSource(&source.Channel{Source: r.ComopsitionChangedWatcher}, handler.EnqueueRequestsFromMapFunc(r.enqueueAllFromGVK)).
		WithOptions(controller.Options{RateLimiter: ratelimiter}).
		Complete(r)
}
