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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServicesGRPCRoutes implements NetworkServicesGRPCRouteInterface
type FakeNetworkServicesGRPCRoutes struct {
	Fake *FakeNetworkservicesV1beta1
	ns   string
}

var networkservicesgrpcroutesResource = v1beta1.SchemeGroupVersion.WithResource("networkservicesgrpcroutes")

var networkservicesgrpcroutesKind = v1beta1.SchemeGroupVersion.WithKind("NetworkServicesGRPCRoute")

// Get takes name of the networkServicesGRPCRoute, and returns the corresponding networkServicesGRPCRoute object, and an error if there is any.
func (c *FakeNetworkServicesGRPCRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkServicesGRPCRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicesgrpcroutesResource, c.ns, name), &v1beta1.NetworkServicesGRPCRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesGRPCRoute), err
}

// List takes label and field selectors, and returns the list of NetworkServicesGRPCRoutes that match those selectors.
func (c *FakeNetworkServicesGRPCRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkServicesGRPCRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicesgrpcroutesResource, networkservicesgrpcroutesKind, c.ns, opts), &v1beta1.NetworkServicesGRPCRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkServicesGRPCRouteList{ListMeta: obj.(*v1beta1.NetworkServicesGRPCRouteList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkServicesGRPCRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServicesGRPCRoutes.
func (c *FakeNetworkServicesGRPCRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicesgrpcroutesResource, c.ns, opts))

}

// Create takes the representation of a networkServicesGRPCRoute and creates it.  Returns the server's representation of the networkServicesGRPCRoute, and an error, if there is any.
func (c *FakeNetworkServicesGRPCRoutes) Create(ctx context.Context, networkServicesGRPCRoute *v1beta1.NetworkServicesGRPCRoute, opts v1.CreateOptions) (result *v1beta1.NetworkServicesGRPCRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicesgrpcroutesResource, c.ns, networkServicesGRPCRoute), &v1beta1.NetworkServicesGRPCRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesGRPCRoute), err
}

// Update takes the representation of a networkServicesGRPCRoute and updates it. Returns the server's representation of the networkServicesGRPCRoute, and an error, if there is any.
func (c *FakeNetworkServicesGRPCRoutes) Update(ctx context.Context, networkServicesGRPCRoute *v1beta1.NetworkServicesGRPCRoute, opts v1.UpdateOptions) (result *v1beta1.NetworkServicesGRPCRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicesgrpcroutesResource, c.ns, networkServicesGRPCRoute), &v1beta1.NetworkServicesGRPCRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesGRPCRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServicesGRPCRoutes) UpdateStatus(ctx context.Context, networkServicesGRPCRoute *v1beta1.NetworkServicesGRPCRoute, opts v1.UpdateOptions) (*v1beta1.NetworkServicesGRPCRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkservicesgrpcroutesResource, "status", c.ns, networkServicesGRPCRoute), &v1beta1.NetworkServicesGRPCRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesGRPCRoute), err
}

// Delete takes name of the networkServicesGRPCRoute and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServicesGRPCRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkservicesgrpcroutesResource, c.ns, name, opts), &v1beta1.NetworkServicesGRPCRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServicesGRPCRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicesgrpcroutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkServicesGRPCRouteList{})
	return err
}

// Patch applies the patch and returns the patched networkServicesGRPCRoute.
func (c *FakeNetworkServicesGRPCRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkServicesGRPCRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicesgrpcroutesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkServicesGRPCRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesGRPCRoute), err
}
