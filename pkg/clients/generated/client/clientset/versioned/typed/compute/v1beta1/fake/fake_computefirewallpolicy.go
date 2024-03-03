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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeFirewallPolicies implements ComputeFirewallPolicyInterface
type FakeComputeFirewallPolicies struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computefirewallpoliciesResource = v1beta1.SchemeGroupVersion.WithResource("computefirewallpolicies")

var computefirewallpoliciesKind = v1beta1.SchemeGroupVersion.WithKind("ComputeFirewallPolicy")

// Get takes name of the computeFirewallPolicy, and returns the corresponding computeFirewallPolicy object, and an error if there is any.
func (c *FakeComputeFirewallPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computefirewallpoliciesResource, c.ns, name), &v1beta1.ComputeFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicy), err
}

// List takes label and field selectors, and returns the list of ComputeFirewallPolicies that match those selectors.
func (c *FakeComputeFirewallPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeFirewallPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computefirewallpoliciesResource, computefirewallpoliciesKind, c.ns, opts), &v1beta1.ComputeFirewallPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeFirewallPolicyList{ListMeta: obj.(*v1beta1.ComputeFirewallPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeFirewallPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeFirewallPolicies.
func (c *FakeComputeFirewallPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computefirewallpoliciesResource, c.ns, opts))

}

// Create takes the representation of a computeFirewallPolicy and creates it.  Returns the server's representation of the computeFirewallPolicy, and an error, if there is any.
func (c *FakeComputeFirewallPolicies) Create(ctx context.Context, computeFirewallPolicy *v1beta1.ComputeFirewallPolicy, opts v1.CreateOptions) (result *v1beta1.ComputeFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computefirewallpoliciesResource, c.ns, computeFirewallPolicy), &v1beta1.ComputeFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicy), err
}

// Update takes the representation of a computeFirewallPolicy and updates it. Returns the server's representation of the computeFirewallPolicy, and an error, if there is any.
func (c *FakeComputeFirewallPolicies) Update(ctx context.Context, computeFirewallPolicy *v1beta1.ComputeFirewallPolicy, opts v1.UpdateOptions) (result *v1beta1.ComputeFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computefirewallpoliciesResource, c.ns, computeFirewallPolicy), &v1beta1.ComputeFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeFirewallPolicies) UpdateStatus(ctx context.Context, computeFirewallPolicy *v1beta1.ComputeFirewallPolicy, opts v1.UpdateOptions) (*v1beta1.ComputeFirewallPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computefirewallpoliciesResource, "status", c.ns, computeFirewallPolicy), &v1beta1.ComputeFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicy), err
}

// Delete takes name of the computeFirewallPolicy and deletes it. Returns an error if one occurs.
func (c *FakeComputeFirewallPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computefirewallpoliciesResource, c.ns, name, opts), &v1beta1.ComputeFirewallPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeFirewallPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computefirewallpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeFirewallPolicyList{})
	return err
}

// Patch applies the patch and returns the patched computeFirewallPolicy.
func (c *FakeComputeFirewallPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computefirewallpoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeFirewallPolicy), err
}
