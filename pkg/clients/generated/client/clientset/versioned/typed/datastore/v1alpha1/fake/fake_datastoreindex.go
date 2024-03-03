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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datastore/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatastoreIndexes implements DatastoreIndexInterface
type FakeDatastoreIndexes struct {
	Fake *FakeDatastoreV1alpha1
	ns   string
}

var datastoreindexesResource = v1alpha1.SchemeGroupVersion.WithResource("datastoreindexes")

var datastoreindexesKind = v1alpha1.SchemeGroupVersion.WithKind("DatastoreIndex")

// Get takes name of the datastoreIndex, and returns the corresponding datastoreIndex object, and an error if there is any.
func (c *FakeDatastoreIndexes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatastoreIndex, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datastoreindexesResource, c.ns, name), &v1alpha1.DatastoreIndex{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatastoreIndex), err
}

// List takes label and field selectors, and returns the list of DatastoreIndexes that match those selectors.
func (c *FakeDatastoreIndexes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatastoreIndexList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datastoreindexesResource, datastoreindexesKind, c.ns, opts), &v1alpha1.DatastoreIndexList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatastoreIndexList{ListMeta: obj.(*v1alpha1.DatastoreIndexList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatastoreIndexList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datastoreIndexes.
func (c *FakeDatastoreIndexes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datastoreindexesResource, c.ns, opts))

}

// Create takes the representation of a datastoreIndex and creates it.  Returns the server's representation of the datastoreIndex, and an error, if there is any.
func (c *FakeDatastoreIndexes) Create(ctx context.Context, datastoreIndex *v1alpha1.DatastoreIndex, opts v1.CreateOptions) (result *v1alpha1.DatastoreIndex, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datastoreindexesResource, c.ns, datastoreIndex), &v1alpha1.DatastoreIndex{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatastoreIndex), err
}

// Update takes the representation of a datastoreIndex and updates it. Returns the server's representation of the datastoreIndex, and an error, if there is any.
func (c *FakeDatastoreIndexes) Update(ctx context.Context, datastoreIndex *v1alpha1.DatastoreIndex, opts v1.UpdateOptions) (result *v1alpha1.DatastoreIndex, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datastoreindexesResource, c.ns, datastoreIndex), &v1alpha1.DatastoreIndex{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatastoreIndex), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatastoreIndexes) UpdateStatus(ctx context.Context, datastoreIndex *v1alpha1.DatastoreIndex, opts v1.UpdateOptions) (*v1alpha1.DatastoreIndex, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datastoreindexesResource, "status", c.ns, datastoreIndex), &v1alpha1.DatastoreIndex{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatastoreIndex), err
}

// Delete takes name of the datastoreIndex and deletes it. Returns an error if one occurs.
func (c *FakeDatastoreIndexes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(datastoreindexesResource, c.ns, name, opts), &v1alpha1.DatastoreIndex{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatastoreIndexes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datastoreindexesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatastoreIndexList{})
	return err
}

// Patch applies the patch and returns the patched datastoreIndex.
func (c *FakeDatastoreIndexes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatastoreIndex, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datastoreindexesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatastoreIndex{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatastoreIndex), err
}
