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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsublite/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePubSubLiteSubscriptions implements PubSubLiteSubscriptionInterface
type FakePubSubLiteSubscriptions struct {
	Fake *FakePubsubliteV1alpha1
	ns   string
}

var pubsublitesubscriptionsResource = v1alpha1.SchemeGroupVersion.WithResource("pubsublitesubscriptions")

var pubsublitesubscriptionsKind = v1alpha1.SchemeGroupVersion.WithKind("PubSubLiteSubscription")

// Get takes name of the pubSubLiteSubscription, and returns the corresponding pubSubLiteSubscription object, and an error if there is any.
func (c *FakePubSubLiteSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PubSubLiteSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pubsublitesubscriptionsResource, c.ns, name), &v1alpha1.PubSubLiteSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubSubLiteSubscription), err
}

// List takes label and field selectors, and returns the list of PubSubLiteSubscriptions that match those selectors.
func (c *FakePubSubLiteSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PubSubLiteSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pubsublitesubscriptionsResource, pubsublitesubscriptionsKind, c.ns, opts), &v1alpha1.PubSubLiteSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PubSubLiteSubscriptionList{ListMeta: obj.(*v1alpha1.PubSubLiteSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PubSubLiteSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubSubLiteSubscriptions.
func (c *FakePubSubLiteSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pubsublitesubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a pubSubLiteSubscription and creates it.  Returns the server's representation of the pubSubLiteSubscription, and an error, if there is any.
func (c *FakePubSubLiteSubscriptions) Create(ctx context.Context, pubSubLiteSubscription *v1alpha1.PubSubLiteSubscription, opts v1.CreateOptions) (result *v1alpha1.PubSubLiteSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pubsublitesubscriptionsResource, c.ns, pubSubLiteSubscription), &v1alpha1.PubSubLiteSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubSubLiteSubscription), err
}

// Update takes the representation of a pubSubLiteSubscription and updates it. Returns the server's representation of the pubSubLiteSubscription, and an error, if there is any.
func (c *FakePubSubLiteSubscriptions) Update(ctx context.Context, pubSubLiteSubscription *v1alpha1.PubSubLiteSubscription, opts v1.UpdateOptions) (result *v1alpha1.PubSubLiteSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pubsublitesubscriptionsResource, c.ns, pubSubLiteSubscription), &v1alpha1.PubSubLiteSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubSubLiteSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubSubLiteSubscriptions) UpdateStatus(ctx context.Context, pubSubLiteSubscription *v1alpha1.PubSubLiteSubscription, opts v1.UpdateOptions) (*v1alpha1.PubSubLiteSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pubsublitesubscriptionsResource, "status", c.ns, pubSubLiteSubscription), &v1alpha1.PubSubLiteSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubSubLiteSubscription), err
}

// Delete takes name of the pubSubLiteSubscription and deletes it. Returns an error if one occurs.
func (c *FakePubSubLiteSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(pubsublitesubscriptionsResource, c.ns, name, opts), &v1alpha1.PubSubLiteSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubSubLiteSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pubsublitesubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PubSubLiteSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched pubSubLiteSubscription.
func (c *FakePubSubLiteSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PubSubLiteSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pubsublitesubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PubSubLiteSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubSubLiteSubscription), err
}
