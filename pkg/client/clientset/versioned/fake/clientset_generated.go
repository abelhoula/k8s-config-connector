// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	clientset "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned"
	accesscontextmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/accesscontextmanager/v1beta1"
	fakeaccesscontextmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/accesscontextmanager/v1beta1/fake"
	artifactregistryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/artifactregistry/v1beta1"
	fakeartifactregistryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/artifactregistry/v1beta1/fake"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/bigquery/v1beta1"
	fakebigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/bigquery/v1beta1/fake"
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/bigtable/v1beta1"
	fakebigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/bigtable/v1beta1/fake"
	binaryauthorizationv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/binaryauthorization/v1beta1"
	fakebinaryauthorizationv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/binaryauthorization/v1beta1/fake"
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudbuild/v1beta1"
	fakecloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudbuild/v1beta1/fake"
	cloudidentityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudidentity/v1beta1"
	fakecloudidentityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudidentity/v1beta1/fake"
	cloudschedulerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudscheduler/v1beta1"
	fakecloudschedulerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/cloudscheduler/v1beta1/fake"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/compute/v1beta1"
	fakecomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/compute/v1beta1/fake"
	configcontrollerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/configcontroller/v1beta1"
	fakeconfigcontrollerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/configcontroller/v1beta1/fake"
	containerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/container/v1beta1"
	fakecontainerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/container/v1beta1/fake"
	containeranalysisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/containeranalysis/v1beta1"
	fakecontaineranalysisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/containeranalysis/v1beta1/fake"
	dataflowv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dataflow/v1beta1"
	fakedataflowv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dataflow/v1beta1/fake"
	datafusionv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/datafusion/v1beta1"
	fakedatafusionv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/datafusion/v1beta1/fake"
	dataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dataproc/v1beta1"
	fakedataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dataproc/v1beta1/fake"
	dnsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dns/v1beta1"
	fakednsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/dns/v1beta1/fake"
	filestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/filestore/v1beta1"
	fakefilestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/filestore/v1beta1/fake"
	firestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/firestore/v1beta1"
	fakefirestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/firestore/v1beta1/fake"
	gameservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/gameservices/v1beta1"
	fakegameservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/gameservices/v1beta1/fake"
	gkehubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/gkehub/v1beta1"
	fakegkehubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/gkehub/v1beta1/fake"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/iam/v1beta1"
	fakeiamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/iam/v1beta1/fake"
	iapv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/iap/v1beta1"
	fakeiapv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/iap/v1beta1/fake"
	identityplatformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/identityplatform/v1beta1"
	fakeidentityplatformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/identityplatform/v1beta1/fake"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/k8s/v1alpha1"
	fakek8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/k8s/v1alpha1/fake"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/kms/v1beta1"
	fakekmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/kms/v1beta1/fake"
	loggingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/logging/v1beta1"
	fakeloggingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/logging/v1beta1/fake"
	memcachev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/memcache/v1beta1"
	fakememcachev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/memcache/v1beta1/fake"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/monitoring/v1beta1"
	fakemonitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/monitoring/v1beta1/fake"
	networkconnectivityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/networkconnectivity/v1beta1"
	fakenetworkconnectivityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/networkconnectivity/v1beta1/fake"
	networksecurityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/networksecurity/v1beta1"
	fakenetworksecurityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/networksecurity/v1beta1/fake"
	networkservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/networkservices/v1beta1"
	fakenetworkservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/networkservices/v1beta1/fake"
	osconfigv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/osconfig/v1beta1"
	fakeosconfigv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/osconfig/v1beta1/fake"
	privatecav1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/privateca/v1beta1"
	fakeprivatecav1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/privateca/v1beta1/fake"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/pubsub/v1beta1"
	fakepubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/pubsub/v1beta1/fake"
	recaptchaenterprisev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/recaptchaenterprise/v1beta1"
	fakerecaptchaenterprisev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/recaptchaenterprise/v1beta1/fake"
	redisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/redis/v1beta1"
	fakeredisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/redis/v1beta1/fake"
	resourcemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/resourcemanager/v1beta1"
	fakeresourcemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/resourcemanager/v1beta1/fake"
	runv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/run/v1beta1"
	fakerunv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/run/v1beta1/fake"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/secretmanager/v1beta1"
	fakesecretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/secretmanager/v1beta1/fake"
	servicenetworkingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/servicenetworking/v1beta1"
	fakeservicenetworkingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/servicenetworking/v1beta1/fake"
	serviceusagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/serviceusage/v1beta1"
	fakeserviceusagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/serviceusage/v1beta1/fake"
	sourcerepov1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/sourcerepo/v1beta1"
	fakesourcerepov1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/sourcerepo/v1beta1/fake"
	spannerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/spanner/v1beta1"
	fakespannerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/spanner/v1beta1/fake"
	sqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/sql/v1beta1"
	fakesqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/sql/v1beta1/fake"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/storage/v1beta1"
	fakestoragev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/storage/v1beta1/fake"
	storagetransferv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/storagetransfer/v1beta1"
	fakestoragetransferv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/storagetransfer/v1beta1/fake"
	vpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/vpcaccess/v1beta1"
	fakevpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/typed/vpcaccess/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// AccesscontextmanagerV1beta1 retrieves the AccesscontextmanagerV1beta1Client
func (c *Clientset) AccesscontextmanagerV1beta1() accesscontextmanagerv1beta1.AccesscontextmanagerV1beta1Interface {
	return &fakeaccesscontextmanagerv1beta1.FakeAccesscontextmanagerV1beta1{Fake: &c.Fake}
}

// ArtifactregistryV1beta1 retrieves the ArtifactregistryV1beta1Client
func (c *Clientset) ArtifactregistryV1beta1() artifactregistryv1beta1.ArtifactregistryV1beta1Interface {
	return &fakeartifactregistryv1beta1.FakeArtifactregistryV1beta1{Fake: &c.Fake}
}

// BigqueryV1beta1 retrieves the BigqueryV1beta1Client
func (c *Clientset) BigqueryV1beta1() bigqueryv1beta1.BigqueryV1beta1Interface {
	return &fakebigqueryv1beta1.FakeBigqueryV1beta1{Fake: &c.Fake}
}

// BigtableV1beta1 retrieves the BigtableV1beta1Client
func (c *Clientset) BigtableV1beta1() bigtablev1beta1.BigtableV1beta1Interface {
	return &fakebigtablev1beta1.FakeBigtableV1beta1{Fake: &c.Fake}
}

// BinaryauthorizationV1beta1 retrieves the BinaryauthorizationV1beta1Client
func (c *Clientset) BinaryauthorizationV1beta1() binaryauthorizationv1beta1.BinaryauthorizationV1beta1Interface {
	return &fakebinaryauthorizationv1beta1.FakeBinaryauthorizationV1beta1{Fake: &c.Fake}
}

// CloudbuildV1beta1 retrieves the CloudbuildV1beta1Client
func (c *Clientset) CloudbuildV1beta1() cloudbuildv1beta1.CloudbuildV1beta1Interface {
	return &fakecloudbuildv1beta1.FakeCloudbuildV1beta1{Fake: &c.Fake}
}

// CloudidentityV1beta1 retrieves the CloudidentityV1beta1Client
func (c *Clientset) CloudidentityV1beta1() cloudidentityv1beta1.CloudidentityV1beta1Interface {
	return &fakecloudidentityv1beta1.FakeCloudidentityV1beta1{Fake: &c.Fake}
}

// CloudschedulerV1beta1 retrieves the CloudschedulerV1beta1Client
func (c *Clientset) CloudschedulerV1beta1() cloudschedulerv1beta1.CloudschedulerV1beta1Interface {
	return &fakecloudschedulerv1beta1.FakeCloudschedulerV1beta1{Fake: &c.Fake}
}

// ComputeV1beta1 retrieves the ComputeV1beta1Client
func (c *Clientset) ComputeV1beta1() computev1beta1.ComputeV1beta1Interface {
	return &fakecomputev1beta1.FakeComputeV1beta1{Fake: &c.Fake}
}

// ConfigcontrollerV1beta1 retrieves the ConfigcontrollerV1beta1Client
func (c *Clientset) ConfigcontrollerV1beta1() configcontrollerv1beta1.ConfigcontrollerV1beta1Interface {
	return &fakeconfigcontrollerv1beta1.FakeConfigcontrollerV1beta1{Fake: &c.Fake}
}

// ContainerV1beta1 retrieves the ContainerV1beta1Client
func (c *Clientset) ContainerV1beta1() containerv1beta1.ContainerV1beta1Interface {
	return &fakecontainerv1beta1.FakeContainerV1beta1{Fake: &c.Fake}
}

// ContaineranalysisV1beta1 retrieves the ContaineranalysisV1beta1Client
func (c *Clientset) ContaineranalysisV1beta1() containeranalysisv1beta1.ContaineranalysisV1beta1Interface {
	return &fakecontaineranalysisv1beta1.FakeContaineranalysisV1beta1{Fake: &c.Fake}
}

// DataflowV1beta1 retrieves the DataflowV1beta1Client
func (c *Clientset) DataflowV1beta1() dataflowv1beta1.DataflowV1beta1Interface {
	return &fakedataflowv1beta1.FakeDataflowV1beta1{Fake: &c.Fake}
}

// DatafusionV1beta1 retrieves the DatafusionV1beta1Client
func (c *Clientset) DatafusionV1beta1() datafusionv1beta1.DatafusionV1beta1Interface {
	return &fakedatafusionv1beta1.FakeDatafusionV1beta1{Fake: &c.Fake}
}

// DataprocV1beta1 retrieves the DataprocV1beta1Client
func (c *Clientset) DataprocV1beta1() dataprocv1beta1.DataprocV1beta1Interface {
	return &fakedataprocv1beta1.FakeDataprocV1beta1{Fake: &c.Fake}
}

// DnsV1beta1 retrieves the DnsV1beta1Client
func (c *Clientset) DnsV1beta1() dnsv1beta1.DnsV1beta1Interface {
	return &fakednsv1beta1.FakeDnsV1beta1{Fake: &c.Fake}
}

// FilestoreV1beta1 retrieves the FilestoreV1beta1Client
func (c *Clientset) FilestoreV1beta1() filestorev1beta1.FilestoreV1beta1Interface {
	return &fakefilestorev1beta1.FakeFilestoreV1beta1{Fake: &c.Fake}
}

// FirestoreV1beta1 retrieves the FirestoreV1beta1Client
func (c *Clientset) FirestoreV1beta1() firestorev1beta1.FirestoreV1beta1Interface {
	return &fakefirestorev1beta1.FakeFirestoreV1beta1{Fake: &c.Fake}
}

// GameservicesV1beta1 retrieves the GameservicesV1beta1Client
func (c *Clientset) GameservicesV1beta1() gameservicesv1beta1.GameservicesV1beta1Interface {
	return &fakegameservicesv1beta1.FakeGameservicesV1beta1{Fake: &c.Fake}
}

// GkehubV1beta1 retrieves the GkehubV1beta1Client
func (c *Clientset) GkehubV1beta1() gkehubv1beta1.GkehubV1beta1Interface {
	return &fakegkehubv1beta1.FakeGkehubV1beta1{Fake: &c.Fake}
}

// IamV1beta1 retrieves the IamV1beta1Client
func (c *Clientset) IamV1beta1() iamv1beta1.IamV1beta1Interface {
	return &fakeiamv1beta1.FakeIamV1beta1{Fake: &c.Fake}
}

// IapV1beta1 retrieves the IapV1beta1Client
func (c *Clientset) IapV1beta1() iapv1beta1.IapV1beta1Interface {
	return &fakeiapv1beta1.FakeIapV1beta1{Fake: &c.Fake}
}

// IdentityplatformV1beta1 retrieves the IdentityplatformV1beta1Client
func (c *Clientset) IdentityplatformV1beta1() identityplatformv1beta1.IdentityplatformV1beta1Interface {
	return &fakeidentityplatformv1beta1.FakeIdentityplatformV1beta1{Fake: &c.Fake}
}

// K8sV1alpha1 retrieves the K8sV1alpha1Client
func (c *Clientset) K8sV1alpha1() k8sv1alpha1.K8sV1alpha1Interface {
	return &fakek8sv1alpha1.FakeK8sV1alpha1{Fake: &c.Fake}
}

// KmsV1beta1 retrieves the KmsV1beta1Client
func (c *Clientset) KmsV1beta1() kmsv1beta1.KmsV1beta1Interface {
	return &fakekmsv1beta1.FakeKmsV1beta1{Fake: &c.Fake}
}

// LoggingV1beta1 retrieves the LoggingV1beta1Client
func (c *Clientset) LoggingV1beta1() loggingv1beta1.LoggingV1beta1Interface {
	return &fakeloggingv1beta1.FakeLoggingV1beta1{Fake: &c.Fake}
}

// MemcacheV1beta1 retrieves the MemcacheV1beta1Client
func (c *Clientset) MemcacheV1beta1() memcachev1beta1.MemcacheV1beta1Interface {
	return &fakememcachev1beta1.FakeMemcacheV1beta1{Fake: &c.Fake}
}

// MonitoringV1beta1 retrieves the MonitoringV1beta1Client
func (c *Clientset) MonitoringV1beta1() monitoringv1beta1.MonitoringV1beta1Interface {
	return &fakemonitoringv1beta1.FakeMonitoringV1beta1{Fake: &c.Fake}
}

// NetworkconnectivityV1beta1 retrieves the NetworkconnectivityV1beta1Client
func (c *Clientset) NetworkconnectivityV1beta1() networkconnectivityv1beta1.NetworkconnectivityV1beta1Interface {
	return &fakenetworkconnectivityv1beta1.FakeNetworkconnectivityV1beta1{Fake: &c.Fake}
}

// NetworksecurityV1beta1 retrieves the NetworksecurityV1beta1Client
func (c *Clientset) NetworksecurityV1beta1() networksecurityv1beta1.NetworksecurityV1beta1Interface {
	return &fakenetworksecurityv1beta1.FakeNetworksecurityV1beta1{Fake: &c.Fake}
}

// NetworkservicesV1beta1 retrieves the NetworkservicesV1beta1Client
func (c *Clientset) NetworkservicesV1beta1() networkservicesv1beta1.NetworkservicesV1beta1Interface {
	return &fakenetworkservicesv1beta1.FakeNetworkservicesV1beta1{Fake: &c.Fake}
}

// OsconfigV1beta1 retrieves the OsconfigV1beta1Client
func (c *Clientset) OsconfigV1beta1() osconfigv1beta1.OsconfigV1beta1Interface {
	return &fakeosconfigv1beta1.FakeOsconfigV1beta1{Fake: &c.Fake}
}

// PrivatecaV1beta1 retrieves the PrivatecaV1beta1Client
func (c *Clientset) PrivatecaV1beta1() privatecav1beta1.PrivatecaV1beta1Interface {
	return &fakeprivatecav1beta1.FakePrivatecaV1beta1{Fake: &c.Fake}
}

// PubsubV1beta1 retrieves the PubsubV1beta1Client
func (c *Clientset) PubsubV1beta1() pubsubv1beta1.PubsubV1beta1Interface {
	return &fakepubsubv1beta1.FakePubsubV1beta1{Fake: &c.Fake}
}

// RecaptchaenterpriseV1beta1 retrieves the RecaptchaenterpriseV1beta1Client
func (c *Clientset) RecaptchaenterpriseV1beta1() recaptchaenterprisev1beta1.RecaptchaenterpriseV1beta1Interface {
	return &fakerecaptchaenterprisev1beta1.FakeRecaptchaenterpriseV1beta1{Fake: &c.Fake}
}

// RedisV1beta1 retrieves the RedisV1beta1Client
func (c *Clientset) RedisV1beta1() redisv1beta1.RedisV1beta1Interface {
	return &fakeredisv1beta1.FakeRedisV1beta1{Fake: &c.Fake}
}

// ResourcemanagerV1beta1 retrieves the ResourcemanagerV1beta1Client
func (c *Clientset) ResourcemanagerV1beta1() resourcemanagerv1beta1.ResourcemanagerV1beta1Interface {
	return &fakeresourcemanagerv1beta1.FakeResourcemanagerV1beta1{Fake: &c.Fake}
}

// RunV1beta1 retrieves the RunV1beta1Client
func (c *Clientset) RunV1beta1() runv1beta1.RunV1beta1Interface {
	return &fakerunv1beta1.FakeRunV1beta1{Fake: &c.Fake}
}

// SecretmanagerV1beta1 retrieves the SecretmanagerV1beta1Client
func (c *Clientset) SecretmanagerV1beta1() secretmanagerv1beta1.SecretmanagerV1beta1Interface {
	return &fakesecretmanagerv1beta1.FakeSecretmanagerV1beta1{Fake: &c.Fake}
}

// ServicenetworkingV1beta1 retrieves the ServicenetworkingV1beta1Client
func (c *Clientset) ServicenetworkingV1beta1() servicenetworkingv1beta1.ServicenetworkingV1beta1Interface {
	return &fakeservicenetworkingv1beta1.FakeServicenetworkingV1beta1{Fake: &c.Fake}
}

// ServiceusageV1beta1 retrieves the ServiceusageV1beta1Client
func (c *Clientset) ServiceusageV1beta1() serviceusagev1beta1.ServiceusageV1beta1Interface {
	return &fakeserviceusagev1beta1.FakeServiceusageV1beta1{Fake: &c.Fake}
}

// SourcerepoV1beta1 retrieves the SourcerepoV1beta1Client
func (c *Clientset) SourcerepoV1beta1() sourcerepov1beta1.SourcerepoV1beta1Interface {
	return &fakesourcerepov1beta1.FakeSourcerepoV1beta1{Fake: &c.Fake}
}

// SpannerV1beta1 retrieves the SpannerV1beta1Client
func (c *Clientset) SpannerV1beta1() spannerv1beta1.SpannerV1beta1Interface {
	return &fakespannerv1beta1.FakeSpannerV1beta1{Fake: &c.Fake}
}

// SqlV1beta1 retrieves the SqlV1beta1Client
func (c *Clientset) SqlV1beta1() sqlv1beta1.SqlV1beta1Interface {
	return &fakesqlv1beta1.FakeSqlV1beta1{Fake: &c.Fake}
}

// StorageV1beta1 retrieves the StorageV1beta1Client
func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return &fakestoragev1beta1.FakeStorageV1beta1{Fake: &c.Fake}
}

// StoragetransferV1beta1 retrieves the StoragetransferV1beta1Client
func (c *Clientset) StoragetransferV1beta1() storagetransferv1beta1.StoragetransferV1beta1Interface {
	return &fakestoragetransferv1beta1.FakeStoragetransferV1beta1{Fake: &c.Fake}
}

// VpcaccessV1beta1 retrieves the VpcaccessV1beta1Client
func (c *Clientset) VpcaccessV1beta1() vpcaccessv1beta1.VpcaccessV1beta1Interface {
	return &fakevpcaccessv1beta1.FakeVpcaccessV1beta1{Fake: &c.Fake}
}
