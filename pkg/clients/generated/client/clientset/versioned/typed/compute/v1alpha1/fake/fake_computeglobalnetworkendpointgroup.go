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

// FakeComputeGlobalNetworkEndpointGroups implements ComputeGlobalNetworkEndpointGroupInterface
type FakeComputeGlobalNetworkEndpointGroups struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var computeglobalnetworkendpointgroupsResource = v1alpha1.SchemeGroupVersion.WithResource("computeglobalnetworkendpointgroups")

var computeglobalnetworkendpointgroupsKind = v1alpha1.SchemeGroupVersion.WithKind("ComputeGlobalNetworkEndpointGroup")

// Get takes name of the computeGlobalNetworkEndpointGroup, and returns the corresponding computeGlobalNetworkEndpointGroup object, and an error if there is any.
func (c *FakeComputeGlobalNetworkEndpointGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeGlobalNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeglobalnetworkendpointgroupsResource, c.ns, name), &v1alpha1.ComputeGlobalNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalNetworkEndpointGroup), err
}

// List takes label and field selectors, and returns the list of ComputeGlobalNetworkEndpointGroups that match those selectors.
func (c *FakeComputeGlobalNetworkEndpointGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeGlobalNetworkEndpointGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeglobalnetworkendpointgroupsResource, computeglobalnetworkendpointgroupsKind, c.ns, opts), &v1alpha1.ComputeGlobalNetworkEndpointGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeGlobalNetworkEndpointGroupList{ListMeta: obj.(*v1alpha1.ComputeGlobalNetworkEndpointGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeGlobalNetworkEndpointGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeGlobalNetworkEndpointGroups.
func (c *FakeComputeGlobalNetworkEndpointGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeglobalnetworkendpointgroupsResource, c.ns, opts))

}

// Create takes the representation of a computeGlobalNetworkEndpointGroup and creates it.  Returns the server's representation of the computeGlobalNetworkEndpointGroup, and an error, if there is any.
func (c *FakeComputeGlobalNetworkEndpointGroups) Create(ctx context.Context, computeGlobalNetworkEndpointGroup *v1alpha1.ComputeGlobalNetworkEndpointGroup, opts v1.CreateOptions) (result *v1alpha1.ComputeGlobalNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeglobalnetworkendpointgroupsResource, c.ns, computeGlobalNetworkEndpointGroup), &v1alpha1.ComputeGlobalNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalNetworkEndpointGroup), err
}

// Update takes the representation of a computeGlobalNetworkEndpointGroup and updates it. Returns the server's representation of the computeGlobalNetworkEndpointGroup, and an error, if there is any.
func (c *FakeComputeGlobalNetworkEndpointGroups) Update(ctx context.Context, computeGlobalNetworkEndpointGroup *v1alpha1.ComputeGlobalNetworkEndpointGroup, opts v1.UpdateOptions) (result *v1alpha1.ComputeGlobalNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeglobalnetworkendpointgroupsResource, c.ns, computeGlobalNetworkEndpointGroup), &v1alpha1.ComputeGlobalNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalNetworkEndpointGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeGlobalNetworkEndpointGroups) UpdateStatus(ctx context.Context, computeGlobalNetworkEndpointGroup *v1alpha1.ComputeGlobalNetworkEndpointGroup, opts v1.UpdateOptions) (*v1alpha1.ComputeGlobalNetworkEndpointGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeglobalnetworkendpointgroupsResource, "status", c.ns, computeGlobalNetworkEndpointGroup), &v1alpha1.ComputeGlobalNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalNetworkEndpointGroup), err
}

// Delete takes name of the computeGlobalNetworkEndpointGroup and deletes it. Returns an error if one occurs.
func (c *FakeComputeGlobalNetworkEndpointGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computeglobalnetworkendpointgroupsResource, c.ns, name, opts), &v1alpha1.ComputeGlobalNetworkEndpointGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeGlobalNetworkEndpointGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeglobalnetworkendpointgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeGlobalNetworkEndpointGroupList{})
	return err
}

// Patch applies the patch and returns the patched computeGlobalNetworkEndpointGroup.
func (c *FakeComputeGlobalNetworkEndpointGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeGlobalNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeglobalnetworkendpointgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeGlobalNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeGlobalNetworkEndpointGroup), err
}
