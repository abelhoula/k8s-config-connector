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

// FakeComputeBackendServiceSignedURLKeys implements ComputeBackendServiceSignedURLKeyInterface
type FakeComputeBackendServiceSignedURLKeys struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var computebackendservicesignedurlkeysResource = v1alpha1.SchemeGroupVersion.WithResource("computebackendservicesignedurlkeys")

var computebackendservicesignedurlkeysKind = v1alpha1.SchemeGroupVersion.WithKind("ComputeBackendServiceSignedURLKey")

// Get takes name of the computeBackendServiceSignedURLKey, and returns the corresponding computeBackendServiceSignedURLKey object, and an error if there is any.
func (c *FakeComputeBackendServiceSignedURLKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeBackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computebackendservicesignedurlkeysResource, c.ns, name), &v1alpha1.ComputeBackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendServiceSignedURLKey), err
}

// List takes label and field selectors, and returns the list of ComputeBackendServiceSignedURLKeys that match those selectors.
func (c *FakeComputeBackendServiceSignedURLKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeBackendServiceSignedURLKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computebackendservicesignedurlkeysResource, computebackendservicesignedurlkeysKind, c.ns, opts), &v1alpha1.ComputeBackendServiceSignedURLKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeBackendServiceSignedURLKeyList{ListMeta: obj.(*v1alpha1.ComputeBackendServiceSignedURLKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeBackendServiceSignedURLKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeBackendServiceSignedURLKeys.
func (c *FakeComputeBackendServiceSignedURLKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computebackendservicesignedurlkeysResource, c.ns, opts))

}

// Create takes the representation of a computeBackendServiceSignedURLKey and creates it.  Returns the server's representation of the computeBackendServiceSignedURLKey, and an error, if there is any.
func (c *FakeComputeBackendServiceSignedURLKeys) Create(ctx context.Context, computeBackendServiceSignedURLKey *v1alpha1.ComputeBackendServiceSignedURLKey, opts v1.CreateOptions) (result *v1alpha1.ComputeBackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computebackendservicesignedurlkeysResource, c.ns, computeBackendServiceSignedURLKey), &v1alpha1.ComputeBackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendServiceSignedURLKey), err
}

// Update takes the representation of a computeBackendServiceSignedURLKey and updates it. Returns the server's representation of the computeBackendServiceSignedURLKey, and an error, if there is any.
func (c *FakeComputeBackendServiceSignedURLKeys) Update(ctx context.Context, computeBackendServiceSignedURLKey *v1alpha1.ComputeBackendServiceSignedURLKey, opts v1.UpdateOptions) (result *v1alpha1.ComputeBackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computebackendservicesignedurlkeysResource, c.ns, computeBackendServiceSignedURLKey), &v1alpha1.ComputeBackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendServiceSignedURLKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeBackendServiceSignedURLKeys) UpdateStatus(ctx context.Context, computeBackendServiceSignedURLKey *v1alpha1.ComputeBackendServiceSignedURLKey, opts v1.UpdateOptions) (*v1alpha1.ComputeBackendServiceSignedURLKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computebackendservicesignedurlkeysResource, "status", c.ns, computeBackendServiceSignedURLKey), &v1alpha1.ComputeBackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendServiceSignedURLKey), err
}

// Delete takes name of the computeBackendServiceSignedURLKey and deletes it. Returns an error if one occurs.
func (c *FakeComputeBackendServiceSignedURLKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computebackendservicesignedurlkeysResource, c.ns, name, opts), &v1alpha1.ComputeBackendServiceSignedURLKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeBackendServiceSignedURLKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computebackendservicesignedurlkeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeBackendServiceSignedURLKeyList{})
	return err
}

// Patch applies the patch and returns the patched computeBackendServiceSignedURLKey.
func (c *FakeComputeBackendServiceSignedURLKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeBackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computebackendservicesignedurlkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeBackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeBackendServiceSignedURLKey), err
}
