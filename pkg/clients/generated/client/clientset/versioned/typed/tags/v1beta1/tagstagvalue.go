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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/tags/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TagsTagValuesGetter has a method to return a TagsTagValueInterface.
// A group's client should implement this interface.
type TagsTagValuesGetter interface {
	TagsTagValues(namespace string) TagsTagValueInterface
}

// TagsTagValueInterface has methods to work with TagsTagValue resources.
type TagsTagValueInterface interface {
	Create(ctx context.Context, tagsTagValue *v1beta1.TagsTagValue, opts v1.CreateOptions) (*v1beta1.TagsTagValue, error)
	Update(ctx context.Context, tagsTagValue *v1beta1.TagsTagValue, opts v1.UpdateOptions) (*v1beta1.TagsTagValue, error)
	UpdateStatus(ctx context.Context, tagsTagValue *v1beta1.TagsTagValue, opts v1.UpdateOptions) (*v1beta1.TagsTagValue, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.TagsTagValue, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.TagsTagValueList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TagsTagValue, err error)
	TagsTagValueExpansion
}

// tagsTagValues implements TagsTagValueInterface
type tagsTagValues struct {
	client rest.Interface
	ns     string
}

// newTagsTagValues returns a TagsTagValues
func newTagsTagValues(c *TagsV1beta1Client, namespace string) *tagsTagValues {
	return &tagsTagValues{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tagsTagValue, and returns the corresponding tagsTagValue object, and an error if there is any.
func (c *tagsTagValues) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.TagsTagValue, err error) {
	result = &v1beta1.TagsTagValue{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tagstagvalues").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TagsTagValues that match those selectors.
func (c *tagsTagValues) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.TagsTagValueList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.TagsTagValueList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tagstagvalues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tagsTagValues.
func (c *tagsTagValues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tagstagvalues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tagsTagValue and creates it.  Returns the server's representation of the tagsTagValue, and an error, if there is any.
func (c *tagsTagValues) Create(ctx context.Context, tagsTagValue *v1beta1.TagsTagValue, opts v1.CreateOptions) (result *v1beta1.TagsTagValue, err error) {
	result = &v1beta1.TagsTagValue{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tagstagvalues").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tagsTagValue).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tagsTagValue and updates it. Returns the server's representation of the tagsTagValue, and an error, if there is any.
func (c *tagsTagValues) Update(ctx context.Context, tagsTagValue *v1beta1.TagsTagValue, opts v1.UpdateOptions) (result *v1beta1.TagsTagValue, err error) {
	result = &v1beta1.TagsTagValue{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tagstagvalues").
		Name(tagsTagValue.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tagsTagValue).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tagsTagValues) UpdateStatus(ctx context.Context, tagsTagValue *v1beta1.TagsTagValue, opts v1.UpdateOptions) (result *v1beta1.TagsTagValue, err error) {
	result = &v1beta1.TagsTagValue{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tagstagvalues").
		Name(tagsTagValue.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tagsTagValue).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tagsTagValue and deletes it. Returns an error if one occurs.
func (c *tagsTagValues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tagstagvalues").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tagsTagValues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tagstagvalues").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tagsTagValue.
func (c *tagsTagValues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TagsTagValue, err error) {
	result = &v1beta1.TagsTagValue{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tagstagvalues").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
