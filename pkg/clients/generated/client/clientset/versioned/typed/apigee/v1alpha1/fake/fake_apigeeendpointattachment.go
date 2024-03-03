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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigee/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApigeeEndpointAttachments implements ApigeeEndpointAttachmentInterface
type FakeApigeeEndpointAttachments struct {
	Fake *FakeApigeeV1alpha1
	ns   string
}

var apigeeendpointattachmentsResource = v1alpha1.SchemeGroupVersion.WithResource("apigeeendpointattachments")

var apigeeendpointattachmentsKind = v1alpha1.SchemeGroupVersion.WithKind("ApigeeEndpointAttachment")

// Get takes name of the apigeeEndpointAttachment, and returns the corresponding apigeeEndpointAttachment object, and an error if there is any.
func (c *FakeApigeeEndpointAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApigeeEndpointAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigeeendpointattachmentsResource, c.ns, name), &v1alpha1.ApigeeEndpointAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEndpointAttachment), err
}

// List takes label and field selectors, and returns the list of ApigeeEndpointAttachments that match those selectors.
func (c *FakeApigeeEndpointAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApigeeEndpointAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigeeendpointattachmentsResource, apigeeendpointattachmentsKind, c.ns, opts), &v1alpha1.ApigeeEndpointAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApigeeEndpointAttachmentList{ListMeta: obj.(*v1alpha1.ApigeeEndpointAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApigeeEndpointAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apigeeEndpointAttachments.
func (c *FakeApigeeEndpointAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigeeendpointattachmentsResource, c.ns, opts))

}

// Create takes the representation of a apigeeEndpointAttachment and creates it.  Returns the server's representation of the apigeeEndpointAttachment, and an error, if there is any.
func (c *FakeApigeeEndpointAttachments) Create(ctx context.Context, apigeeEndpointAttachment *v1alpha1.ApigeeEndpointAttachment, opts v1.CreateOptions) (result *v1alpha1.ApigeeEndpointAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigeeendpointattachmentsResource, c.ns, apigeeEndpointAttachment), &v1alpha1.ApigeeEndpointAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEndpointAttachment), err
}

// Update takes the representation of a apigeeEndpointAttachment and updates it. Returns the server's representation of the apigeeEndpointAttachment, and an error, if there is any.
func (c *FakeApigeeEndpointAttachments) Update(ctx context.Context, apigeeEndpointAttachment *v1alpha1.ApigeeEndpointAttachment, opts v1.UpdateOptions) (result *v1alpha1.ApigeeEndpointAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigeeendpointattachmentsResource, c.ns, apigeeEndpointAttachment), &v1alpha1.ApigeeEndpointAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEndpointAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApigeeEndpointAttachments) UpdateStatus(ctx context.Context, apigeeEndpointAttachment *v1alpha1.ApigeeEndpointAttachment, opts v1.UpdateOptions) (*v1alpha1.ApigeeEndpointAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigeeendpointattachmentsResource, "status", c.ns, apigeeEndpointAttachment), &v1alpha1.ApigeeEndpointAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEndpointAttachment), err
}

// Delete takes name of the apigeeEndpointAttachment and deletes it. Returns an error if one occurs.
func (c *FakeApigeeEndpointAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apigeeendpointattachmentsResource, c.ns, name, opts), &v1alpha1.ApigeeEndpointAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApigeeEndpointAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigeeendpointattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApigeeEndpointAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched apigeeEndpointAttachment.
func (c *FakeApigeeEndpointAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApigeeEndpointAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigeeendpointattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApigeeEndpointAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEndpointAttachment), err
}
