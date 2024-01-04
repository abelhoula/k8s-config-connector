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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkconnectivity/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkConnectivityHubs implements NetworkConnectivityHubInterface
type FakeNetworkConnectivityHubs struct {
	Fake *FakeNetworkconnectivityV1beta1
	ns   string
}

var networkconnectivityhubsResource = v1beta1.SchemeGroupVersion.WithResource("networkconnectivityhubs")

var networkconnectivityhubsKind = v1beta1.SchemeGroupVersion.WithKind("NetworkConnectivityHub")

// Get takes name of the networkConnectivityHub, and returns the corresponding networkConnectivityHub object, and an error if there is any.
func (c *FakeNetworkConnectivityHubs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkconnectivityhubsResource, c.ns, name), &v1beta1.NetworkConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkConnectivityHub), err
}

// List takes label and field selectors, and returns the list of NetworkConnectivityHubs that match those selectors.
func (c *FakeNetworkConnectivityHubs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkConnectivityHubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkconnectivityhubsResource, networkconnectivityhubsKind, c.ns, opts), &v1beta1.NetworkConnectivityHubList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkConnectivityHubList{ListMeta: obj.(*v1beta1.NetworkConnectivityHubList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkConnectivityHubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkConnectivityHubs.
func (c *FakeNetworkConnectivityHubs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkconnectivityhubsResource, c.ns, opts))

}

// Create takes the representation of a networkConnectivityHub and creates it.  Returns the server's representation of the networkConnectivityHub, and an error, if there is any.
func (c *FakeNetworkConnectivityHubs) Create(ctx context.Context, networkConnectivityHub *v1beta1.NetworkConnectivityHub, opts v1.CreateOptions) (result *v1beta1.NetworkConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkconnectivityhubsResource, c.ns, networkConnectivityHub), &v1beta1.NetworkConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkConnectivityHub), err
}

// Update takes the representation of a networkConnectivityHub and updates it. Returns the server's representation of the networkConnectivityHub, and an error, if there is any.
func (c *FakeNetworkConnectivityHubs) Update(ctx context.Context, networkConnectivityHub *v1beta1.NetworkConnectivityHub, opts v1.UpdateOptions) (result *v1beta1.NetworkConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkconnectivityhubsResource, c.ns, networkConnectivityHub), &v1beta1.NetworkConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkConnectivityHub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkConnectivityHubs) UpdateStatus(ctx context.Context, networkConnectivityHub *v1beta1.NetworkConnectivityHub, opts v1.UpdateOptions) (*v1beta1.NetworkConnectivityHub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkconnectivityhubsResource, "status", c.ns, networkConnectivityHub), &v1beta1.NetworkConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkConnectivityHub), err
}

// Delete takes name of the networkConnectivityHub and deletes it. Returns an error if one occurs.
func (c *FakeNetworkConnectivityHubs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkconnectivityhubsResource, c.ns, name, opts), &v1beta1.NetworkConnectivityHub{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkConnectivityHubs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkconnectivityhubsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkConnectivityHubList{})
	return err
}

// Patch applies the patch and returns the patched networkConnectivityHub.
func (c *FakeNetworkConnectivityHubs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkconnectivityhubsResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkConnectivityHub), err
}
