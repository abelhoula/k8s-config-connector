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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firebasehosting/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FirebaseHostingChannelsGetter has a method to return a FirebaseHostingChannelInterface.
// A group's client should implement this interface.
type FirebaseHostingChannelsGetter interface {
	FirebaseHostingChannels(namespace string) FirebaseHostingChannelInterface
}

// FirebaseHostingChannelInterface has methods to work with FirebaseHostingChannel resources.
type FirebaseHostingChannelInterface interface {
	Create(ctx context.Context, firebaseHostingChannel *v1alpha1.FirebaseHostingChannel, opts v1.CreateOptions) (*v1alpha1.FirebaseHostingChannel, error)
	Update(ctx context.Context, firebaseHostingChannel *v1alpha1.FirebaseHostingChannel, opts v1.UpdateOptions) (*v1alpha1.FirebaseHostingChannel, error)
	UpdateStatus(ctx context.Context, firebaseHostingChannel *v1alpha1.FirebaseHostingChannel, opts v1.UpdateOptions) (*v1alpha1.FirebaseHostingChannel, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FirebaseHostingChannel, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FirebaseHostingChannelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirebaseHostingChannel, err error)
	FirebaseHostingChannelExpansion
}

// firebaseHostingChannels implements FirebaseHostingChannelInterface
type firebaseHostingChannels struct {
	client rest.Interface
	ns     string
}

// newFirebaseHostingChannels returns a FirebaseHostingChannels
func newFirebaseHostingChannels(c *FirebasehostingV1alpha1Client, namespace string) *firebaseHostingChannels {
	return &firebaseHostingChannels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the firebaseHostingChannel, and returns the corresponding firebaseHostingChannel object, and an error if there is any.
func (c *firebaseHostingChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FirebaseHostingChannel, err error) {
	result = &v1alpha1.FirebaseHostingChannel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FirebaseHostingChannels that match those selectors.
func (c *firebaseHostingChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FirebaseHostingChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FirebaseHostingChannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested firebaseHostingChannels.
func (c *firebaseHostingChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a firebaseHostingChannel and creates it.  Returns the server's representation of the firebaseHostingChannel, and an error, if there is any.
func (c *firebaseHostingChannels) Create(ctx context.Context, firebaseHostingChannel *v1alpha1.FirebaseHostingChannel, opts v1.CreateOptions) (result *v1alpha1.FirebaseHostingChannel, err error) {
	result = &v1alpha1.FirebaseHostingChannel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(firebaseHostingChannel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a firebaseHostingChannel and updates it. Returns the server's representation of the firebaseHostingChannel, and an error, if there is any.
func (c *firebaseHostingChannels) Update(ctx context.Context, firebaseHostingChannel *v1alpha1.FirebaseHostingChannel, opts v1.UpdateOptions) (result *v1alpha1.FirebaseHostingChannel, err error) {
	result = &v1alpha1.FirebaseHostingChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		Name(firebaseHostingChannel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(firebaseHostingChannel).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *firebaseHostingChannels) UpdateStatus(ctx context.Context, firebaseHostingChannel *v1alpha1.FirebaseHostingChannel, opts v1.UpdateOptions) (result *v1alpha1.FirebaseHostingChannel, err error) {
	result = &v1alpha1.FirebaseHostingChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		Name(firebaseHostingChannel.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(firebaseHostingChannel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the firebaseHostingChannel and deletes it. Returns an error if one occurs.
func (c *firebaseHostingChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *firebaseHostingChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched firebaseHostingChannel.
func (c *firebaseHostingChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirebaseHostingChannel, err error) {
	result = &v1alpha1.FirebaseHostingChannel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("firebasehostingchannels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
