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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/binaryauthorization/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BinaryAuthorizationPoliciesGetter has a method to return a BinaryAuthorizationPolicyInterface.
// A group's client should implement this interface.
type BinaryAuthorizationPoliciesGetter interface {
	BinaryAuthorizationPolicies(namespace string) BinaryAuthorizationPolicyInterface
}

// BinaryAuthorizationPolicyInterface has methods to work with BinaryAuthorizationPolicy resources.
type BinaryAuthorizationPolicyInterface interface {
	Create(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.CreateOptions) (*v1beta1.BinaryAuthorizationPolicy, error)
	Update(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.UpdateOptions) (*v1beta1.BinaryAuthorizationPolicy, error)
	UpdateStatus(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.UpdateOptions) (*v1beta1.BinaryAuthorizationPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BinaryAuthorizationPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BinaryAuthorizationPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BinaryAuthorizationPolicy, err error)
	BinaryAuthorizationPolicyExpansion
}

// binaryAuthorizationPolicies implements BinaryAuthorizationPolicyInterface
type binaryAuthorizationPolicies struct {
	client rest.Interface
	ns     string
}

// newBinaryAuthorizationPolicies returns a BinaryAuthorizationPolicies
func newBinaryAuthorizationPolicies(c *BinaryauthorizationV1beta1Client, namespace string) *binaryAuthorizationPolicies {
	return &binaryAuthorizationPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the binaryAuthorizationPolicy, and returns the corresponding binaryAuthorizationPolicy object, and an error if there is any.
func (c *binaryAuthorizationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	result = &v1beta1.BinaryAuthorizationPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BinaryAuthorizationPolicies that match those selectors.
func (c *binaryAuthorizationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BinaryAuthorizationPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BinaryAuthorizationPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested binaryAuthorizationPolicies.
func (c *binaryAuthorizationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a binaryAuthorizationPolicy and creates it.  Returns the server's representation of the binaryAuthorizationPolicy, and an error, if there is any.
func (c *binaryAuthorizationPolicies) Create(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.CreateOptions) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	result = &v1beta1.BinaryAuthorizationPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(binaryAuthorizationPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a binaryAuthorizationPolicy and updates it. Returns the server's representation of the binaryAuthorizationPolicy, and an error, if there is any.
func (c *binaryAuthorizationPolicies) Update(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.UpdateOptions) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	result = &v1beta1.BinaryAuthorizationPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		Name(binaryAuthorizationPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(binaryAuthorizationPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *binaryAuthorizationPolicies) UpdateStatus(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.UpdateOptions) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	result = &v1beta1.BinaryAuthorizationPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		Name(binaryAuthorizationPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(binaryAuthorizationPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the binaryAuthorizationPolicy and deletes it. Returns an error if one occurs.
func (c *binaryAuthorizationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *binaryAuthorizationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched binaryAuthorizationPolicy.
func (c *binaryAuthorizationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	result = &v1beta1.BinaryAuthorizationPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("binaryauthorizationpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
