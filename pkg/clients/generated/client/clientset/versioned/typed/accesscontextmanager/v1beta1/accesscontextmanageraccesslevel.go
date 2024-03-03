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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/accesscontextmanager/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AccessContextManagerAccessLevelsGetter has a method to return a AccessContextManagerAccessLevelInterface.
// A group's client should implement this interface.
type AccessContextManagerAccessLevelsGetter interface {
	AccessContextManagerAccessLevels(namespace string) AccessContextManagerAccessLevelInterface
}

// AccessContextManagerAccessLevelInterface has methods to work with AccessContextManagerAccessLevel resources.
type AccessContextManagerAccessLevelInterface interface {
	Create(ctx context.Context, accessContextManagerAccessLevel *v1beta1.AccessContextManagerAccessLevel, opts v1.CreateOptions) (*v1beta1.AccessContextManagerAccessLevel, error)
	Update(ctx context.Context, accessContextManagerAccessLevel *v1beta1.AccessContextManagerAccessLevel, opts v1.UpdateOptions) (*v1beta1.AccessContextManagerAccessLevel, error)
	UpdateStatus(ctx context.Context, accessContextManagerAccessLevel *v1beta1.AccessContextManagerAccessLevel, opts v1.UpdateOptions) (*v1beta1.AccessContextManagerAccessLevel, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.AccessContextManagerAccessLevel, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.AccessContextManagerAccessLevelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AccessContextManagerAccessLevel, err error)
	AccessContextManagerAccessLevelExpansion
}

// accessContextManagerAccessLevels implements AccessContextManagerAccessLevelInterface
type accessContextManagerAccessLevels struct {
	client rest.Interface
	ns     string
}

// newAccessContextManagerAccessLevels returns a AccessContextManagerAccessLevels
func newAccessContextManagerAccessLevels(c *AccesscontextmanagerV1beta1Client, namespace string) *accessContextManagerAccessLevels {
	return &accessContextManagerAccessLevels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the accessContextManagerAccessLevel, and returns the corresponding accessContextManagerAccessLevel object, and an error if there is any.
func (c *accessContextManagerAccessLevels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AccessContextManagerAccessLevel, err error) {
	result = &v1beta1.AccessContextManagerAccessLevel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AccessContextManagerAccessLevels that match those selectors.
func (c *accessContextManagerAccessLevels) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AccessContextManagerAccessLevelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.AccessContextManagerAccessLevelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested accessContextManagerAccessLevels.
func (c *accessContextManagerAccessLevels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a accessContextManagerAccessLevel and creates it.  Returns the server's representation of the accessContextManagerAccessLevel, and an error, if there is any.
func (c *accessContextManagerAccessLevels) Create(ctx context.Context, accessContextManagerAccessLevel *v1beta1.AccessContextManagerAccessLevel, opts v1.CreateOptions) (result *v1beta1.AccessContextManagerAccessLevel, err error) {
	result = &v1beta1.AccessContextManagerAccessLevel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(accessContextManagerAccessLevel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a accessContextManagerAccessLevel and updates it. Returns the server's representation of the accessContextManagerAccessLevel, and an error, if there is any.
func (c *accessContextManagerAccessLevels) Update(ctx context.Context, accessContextManagerAccessLevel *v1beta1.AccessContextManagerAccessLevel, opts v1.UpdateOptions) (result *v1beta1.AccessContextManagerAccessLevel, err error) {
	result = &v1beta1.AccessContextManagerAccessLevel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		Name(accessContextManagerAccessLevel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(accessContextManagerAccessLevel).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *accessContextManagerAccessLevels) UpdateStatus(ctx context.Context, accessContextManagerAccessLevel *v1beta1.AccessContextManagerAccessLevel, opts v1.UpdateOptions) (result *v1beta1.AccessContextManagerAccessLevel, err error) {
	result = &v1beta1.AccessContextManagerAccessLevel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		Name(accessContextManagerAccessLevel.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(accessContextManagerAccessLevel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the accessContextManagerAccessLevel and deletes it. Returns an error if one occurs.
func (c *accessContextManagerAccessLevels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *accessContextManagerAccessLevels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched accessContextManagerAccessLevel.
func (c *accessContextManagerAccessLevels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AccessContextManagerAccessLevel, err error) {
	result = &v1beta1.AccessContextManagerAccessLevel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("accesscontextmanageraccesslevels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
