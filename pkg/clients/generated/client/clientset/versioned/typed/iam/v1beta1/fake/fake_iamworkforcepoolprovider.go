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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIAMWorkforcePoolProviders implements IAMWorkforcePoolProviderInterface
type FakeIAMWorkforcePoolProviders struct {
	Fake *FakeIamV1beta1
	ns   string
}

var iamworkforcepoolprovidersResource = v1beta1.SchemeGroupVersion.WithResource("iamworkforcepoolproviders")

var iamworkforcepoolprovidersKind = v1beta1.SchemeGroupVersion.WithKind("IAMWorkforcePoolProvider")

// Get takes name of the iAMWorkforcePoolProvider, and returns the corresponding iAMWorkforcePoolProvider object, and an error if there is any.
func (c *FakeIAMWorkforcePoolProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IAMWorkforcePoolProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamworkforcepoolprovidersResource, c.ns, name), &v1beta1.IAMWorkforcePoolProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMWorkforcePoolProvider), err
}

// List takes label and field selectors, and returns the list of IAMWorkforcePoolProviders that match those selectors.
func (c *FakeIAMWorkforcePoolProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IAMWorkforcePoolProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamworkforcepoolprovidersResource, iamworkforcepoolprovidersKind, c.ns, opts), &v1beta1.IAMWorkforcePoolProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IAMWorkforcePoolProviderList{ListMeta: obj.(*v1beta1.IAMWorkforcePoolProviderList).ListMeta}
	for _, item := range obj.(*v1beta1.IAMWorkforcePoolProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iAMWorkforcePoolProviders.
func (c *FakeIAMWorkforcePoolProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamworkforcepoolprovidersResource, c.ns, opts))

}

// Create takes the representation of a iAMWorkforcePoolProvider and creates it.  Returns the server's representation of the iAMWorkforcePoolProvider, and an error, if there is any.
func (c *FakeIAMWorkforcePoolProviders) Create(ctx context.Context, iAMWorkforcePoolProvider *v1beta1.IAMWorkforcePoolProvider, opts v1.CreateOptions) (result *v1beta1.IAMWorkforcePoolProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamworkforcepoolprovidersResource, c.ns, iAMWorkforcePoolProvider), &v1beta1.IAMWorkforcePoolProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMWorkforcePoolProvider), err
}

// Update takes the representation of a iAMWorkforcePoolProvider and updates it. Returns the server's representation of the iAMWorkforcePoolProvider, and an error, if there is any.
func (c *FakeIAMWorkforcePoolProviders) Update(ctx context.Context, iAMWorkforcePoolProvider *v1beta1.IAMWorkforcePoolProvider, opts v1.UpdateOptions) (result *v1beta1.IAMWorkforcePoolProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamworkforcepoolprovidersResource, c.ns, iAMWorkforcePoolProvider), &v1beta1.IAMWorkforcePoolProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMWorkforcePoolProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIAMWorkforcePoolProviders) UpdateStatus(ctx context.Context, iAMWorkforcePoolProvider *v1beta1.IAMWorkforcePoolProvider, opts v1.UpdateOptions) (*v1beta1.IAMWorkforcePoolProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamworkforcepoolprovidersResource, "status", c.ns, iAMWorkforcePoolProvider), &v1beta1.IAMWorkforcePoolProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMWorkforcePoolProvider), err
}

// Delete takes name of the iAMWorkforcePoolProvider and deletes it. Returns an error if one occurs.
func (c *FakeIAMWorkforcePoolProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(iamworkforcepoolprovidersResource, c.ns, name, opts), &v1beta1.IAMWorkforcePoolProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIAMWorkforcePoolProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamworkforcepoolprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IAMWorkforcePoolProviderList{})
	return err
}

// Patch applies the patch and returns the patched iAMWorkforcePoolProvider.
func (c *FakeIAMWorkforcePoolProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IAMWorkforcePoolProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamworkforcepoolprovidersResource, c.ns, name, pt, data, subresources...), &v1beta1.IAMWorkforcePoolProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMWorkforcePoolProvider), err
}
