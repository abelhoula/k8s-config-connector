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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeOrganizationSecurityPolicyAssociations implements ComputeOrganizationSecurityPolicyAssociationInterface
type FakeComputeOrganizationSecurityPolicyAssociations struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var computeorganizationsecuritypolicyassociationsResource = v1alpha1.SchemeGroupVersion.WithResource("computeorganizationsecuritypolicyassociations")

var computeorganizationsecuritypolicyassociationsKind = v1alpha1.SchemeGroupVersion.WithKind("ComputeOrganizationSecurityPolicyAssociation")

// Get takes name of the computeOrganizationSecurityPolicyAssociation, and returns the corresponding computeOrganizationSecurityPolicyAssociation object, and an error if there is any.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeorganizationsecuritypolicyassociationsResource, c.ns, name), &v1alpha1.ComputeOrganizationSecurityPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyAssociation), err
}

// List takes label and field selectors, and returns the list of ComputeOrganizationSecurityPolicyAssociations that match those selectors.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeorganizationsecuritypolicyassociationsResource, computeorganizationsecuritypolicyassociationsKind, c.ns, opts), &v1alpha1.ComputeOrganizationSecurityPolicyAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeOrganizationSecurityPolicyAssociationList{ListMeta: obj.(*v1alpha1.ComputeOrganizationSecurityPolicyAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeOrganizationSecurityPolicyAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeOrganizationSecurityPolicyAssociations.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeorganizationsecuritypolicyassociationsResource, c.ns, opts))

}

// Create takes the representation of a computeOrganizationSecurityPolicyAssociation and creates it.  Returns the server's representation of the computeOrganizationSecurityPolicyAssociation, and an error, if there is any.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) Create(ctx context.Context, computeOrganizationSecurityPolicyAssociation *v1alpha1.ComputeOrganizationSecurityPolicyAssociation, opts v1.CreateOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeorganizationsecuritypolicyassociationsResource, c.ns, computeOrganizationSecurityPolicyAssociation), &v1alpha1.ComputeOrganizationSecurityPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyAssociation), err
}

// Update takes the representation of a computeOrganizationSecurityPolicyAssociation and updates it. Returns the server's representation of the computeOrganizationSecurityPolicyAssociation, and an error, if there is any.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) Update(ctx context.Context, computeOrganizationSecurityPolicyAssociation *v1alpha1.ComputeOrganizationSecurityPolicyAssociation, opts v1.UpdateOptions) (result *v1alpha1.ComputeOrganizationSecurityPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeorganizationsecuritypolicyassociationsResource, c.ns, computeOrganizationSecurityPolicyAssociation), &v1alpha1.ComputeOrganizationSecurityPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeOrganizationSecurityPolicyAssociations) UpdateStatus(ctx context.Context, computeOrganizationSecurityPolicyAssociation *v1alpha1.ComputeOrganizationSecurityPolicyAssociation, opts v1.UpdateOptions) (*v1alpha1.ComputeOrganizationSecurityPolicyAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeorganizationsecuritypolicyassociationsResource, "status", c.ns, computeOrganizationSecurityPolicyAssociation), &v1alpha1.ComputeOrganizationSecurityPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyAssociation), err
}

// Delete takes name of the computeOrganizationSecurityPolicyAssociation and deletes it. Returns an error if one occurs.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computeorganizationsecuritypolicyassociationsResource, c.ns, name, opts), &v1alpha1.ComputeOrganizationSecurityPolicyAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeorganizationsecuritypolicyassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeOrganizationSecurityPolicyAssociationList{})
	return err
}

// Patch applies the patch and returns the patched computeOrganizationSecurityPolicyAssociation.
func (c *FakeComputeOrganizationSecurityPolicyAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeOrganizationSecurityPolicyAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeorganizationsecuritypolicyassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeOrganizationSecurityPolicyAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeOrganizationSecurityPolicyAssociation), err
}
