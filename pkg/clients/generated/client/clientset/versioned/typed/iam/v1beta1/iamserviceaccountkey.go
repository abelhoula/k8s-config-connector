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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IAMServiceAccountKeysGetter has a method to return a IAMServiceAccountKeyInterface.
// A group's client should implement this interface.
type IAMServiceAccountKeysGetter interface {
	IAMServiceAccountKeys(namespace string) IAMServiceAccountKeyInterface
}

// IAMServiceAccountKeyInterface has methods to work with IAMServiceAccountKey resources.
type IAMServiceAccountKeyInterface interface {
	Create(ctx context.Context, iAMServiceAccountKey *v1beta1.IAMServiceAccountKey, opts v1.CreateOptions) (*v1beta1.IAMServiceAccountKey, error)
	Update(ctx context.Context, iAMServiceAccountKey *v1beta1.IAMServiceAccountKey, opts v1.UpdateOptions) (*v1beta1.IAMServiceAccountKey, error)
	UpdateStatus(ctx context.Context, iAMServiceAccountKey *v1beta1.IAMServiceAccountKey, opts v1.UpdateOptions) (*v1beta1.IAMServiceAccountKey, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.IAMServiceAccountKey, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.IAMServiceAccountKeyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IAMServiceAccountKey, err error)
	IAMServiceAccountKeyExpansion
}

// iAMServiceAccountKeys implements IAMServiceAccountKeyInterface
type iAMServiceAccountKeys struct {
	client rest.Interface
	ns     string
}

// newIAMServiceAccountKeys returns a IAMServiceAccountKeys
func newIAMServiceAccountKeys(c *IamV1beta1Client, namespace string) *iAMServiceAccountKeys {
	return &iAMServiceAccountKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iAMServiceAccountKey, and returns the corresponding iAMServiceAccountKey object, and an error if there is any.
func (c *iAMServiceAccountKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IAMServiceAccountKey, err error) {
	result = &v1beta1.IAMServiceAccountKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IAMServiceAccountKeys that match those selectors.
func (c *iAMServiceAccountKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IAMServiceAccountKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.IAMServiceAccountKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iAMServiceAccountKeys.
func (c *iAMServiceAccountKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iAMServiceAccountKey and creates it.  Returns the server's representation of the iAMServiceAccountKey, and an error, if there is any.
func (c *iAMServiceAccountKeys) Create(ctx context.Context, iAMServiceAccountKey *v1beta1.IAMServiceAccountKey, opts v1.CreateOptions) (result *v1beta1.IAMServiceAccountKey, err error) {
	result = &v1beta1.IAMServiceAccountKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iAMServiceAccountKey).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iAMServiceAccountKey and updates it. Returns the server's representation of the iAMServiceAccountKey, and an error, if there is any.
func (c *iAMServiceAccountKeys) Update(ctx context.Context, iAMServiceAccountKey *v1beta1.IAMServiceAccountKey, opts v1.UpdateOptions) (result *v1beta1.IAMServiceAccountKey, err error) {
	result = &v1beta1.IAMServiceAccountKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		Name(iAMServiceAccountKey.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iAMServiceAccountKey).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *iAMServiceAccountKeys) UpdateStatus(ctx context.Context, iAMServiceAccountKey *v1beta1.IAMServiceAccountKey, opts v1.UpdateOptions) (result *v1beta1.IAMServiceAccountKey, err error) {
	result = &v1beta1.IAMServiceAccountKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		Name(iAMServiceAccountKey.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iAMServiceAccountKey).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iAMServiceAccountKey and deletes it. Returns an error if one occurs.
func (c *iAMServiceAccountKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iAMServiceAccountKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iAMServiceAccountKey.
func (c *iAMServiceAccountKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IAMServiceAccountKey, err error) {
	result = &v1beta1.IAMServiceAccountKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamserviceaccountkeys").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
