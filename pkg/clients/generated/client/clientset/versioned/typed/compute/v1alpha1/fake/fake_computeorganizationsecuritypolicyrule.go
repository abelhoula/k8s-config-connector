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
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeOrganizationSecurityPolicyRules implements ComputeOrganizationSecurityPolicyRuleInterface
type FakeComputeOrganizationSecurityPolicyRules struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var computeorganizationsecuritypolicyrulesResource = v1alpha1.SchemeGroupVersion.WithResource("computeorganizationsecuritypolicyrules")

var computeorganizationsecuritypolicyrulesKind = v1alpha1.SchemeGroupVersion.WithKind("ComputeOrganizationSecurityPolicyRule")

// Get takes name of the computeOrganizationSecurityPolicyRule, and returns the corresponding computeOrganizationSecurityPolicyRule object, and an error if there is any.
func (c *FakeComputeOrganizationSecurityPolicyRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeorganizationsecuritypolicyrulesResource, c.ns, name), &v1alpha1.ComputeOrganizationSecurityPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyRule), err
}

// List takes label and field selectors, and returns the list of ComputeOrganizationSecurityPolicyRules that match those selectors.
func (c *FakeComputeOrganizationSecurityPolicyRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeorganizationsecuritypolicyrulesResource, computeorganizationsecuritypolicyrulesKind, c.ns, opts), &v1alpha1.ComputeOrganizationSecurityPolicyRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeOrganizationSecurityPolicyRuleList{ListMeta: obj.(*v1alpha1.ComputeOrganizationSecurityPolicyRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeOrganizationSecurityPolicyRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeOrganizationSecurityPolicyRules.
func (c *FakeComputeOrganizationSecurityPolicyRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeorganizationsecuritypolicyrulesResource, c.ns, opts))

}

// Create takes the representation of a computeOrganizationSecurityPolicyRule and creates it.  Returns the server's representation of the computeOrganizationSecurityPolicyRule, and an error, if there is any.
func (c *FakeComputeOrganizationSecurityPolicyRules) Create(ctx context.Context, computeOrganizationSecurityPolicyRule *v1alpha1.ComputeOrganizationSecurityPolicyRule, opts v1.CreateOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeorganizationsecuritypolicyrulesResource, c.ns, computeOrganizationSecurityPolicyRule), &v1alpha1.ComputeOrganizationSecurityPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyRule), err
}

// Update takes the representation of a computeOrganizationSecurityPolicyRule and updates it. Returns the server's representation of the computeOrganizationSecurityPolicyRule, and an error, if there is any.
func (c *FakeComputeOrganizationSecurityPolicyRules) Update(ctx context.Context, computeOrganizationSecurityPolicyRule *v1alpha1.ComputeOrganizationSecurityPolicyRule, opts v1.UpdateOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeorganizationsecuritypolicyrulesResource, c.ns, computeOrganizationSecurityPolicyRule), &v1alpha1.ComputeOrganizationSecurityPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeOrganizationSecurityPolicyRules) UpdateStatus(ctx context.Context, computeOrganizationSecurityPolicyRule *v1alpha1.ComputeOrganizationSecurityPolicyRule, opts v1.UpdateOptions) (*v1alpha1.ComputeOrganizationSecurityPolicyRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeorganizationsecuritypolicyrulesResource, "status", c.ns, computeOrganizationSecurityPolicyRule), &v1alpha1.ComputeOrganizationSecurityPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyRule), err
}

// Delete takes name of the computeOrganizationSecurityPolicyRule and deletes it. Returns an error if one occurs.
func (c *FakeComputeOrganizationSecurityPolicyRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computeorganizationsecuritypolicyrulesResource, c.ns, name, opts), &v1alpha1.ComputeOrganizationSecurityPolicyRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeOrganizationSecurityPolicyRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeorganizationsecuritypolicyrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeOrganizationSecurityPolicyRuleList{})
	return err
}

// Patch applies the patch and returns the patched computeOrganizationSecurityPolicyRule.
func (c *FakeComputeOrganizationSecurityPolicyRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeOrganizationSecurityPolicyRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeorganizationsecuritypolicyrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeOrganizationSecurityPolicyRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyRule), err
}
