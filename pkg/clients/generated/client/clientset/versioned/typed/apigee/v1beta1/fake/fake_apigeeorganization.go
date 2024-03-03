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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigee/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApigeeOrganizations implements ApigeeOrganizationInterface
type FakeApigeeOrganizations struct {
	Fake *FakeApigeeV1beta1
	ns   string
}

var apigeeorganizationsResource = v1beta1.SchemeGroupVersion.WithResource("apigeeorganizations")

var apigeeorganizationsKind = v1beta1.SchemeGroupVersion.WithKind("ApigeeOrganization")

// Get takes name of the apigeeOrganization, and returns the corresponding apigeeOrganization object, and an error if there is any.
func (c *FakeApigeeOrganizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ApigeeOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigeeorganizationsResource, c.ns, name), &v1beta1.ApigeeOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ApigeeOrganization), err
}

// List takes label and field selectors, and returns the list of ApigeeOrganizations that match those selectors.
func (c *FakeApigeeOrganizations) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ApigeeOrganizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigeeorganizationsResource, apigeeorganizationsKind, c.ns, opts), &v1beta1.ApigeeOrganizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ApigeeOrganizationList{ListMeta: obj.(*v1beta1.ApigeeOrganizationList).ListMeta}
	for _, item := range obj.(*v1beta1.ApigeeOrganizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apigeeOrganizations.
func (c *FakeApigeeOrganizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigeeorganizationsResource, c.ns, opts))

}

// Create takes the representation of a apigeeOrganization and creates it.  Returns the server's representation of the apigeeOrganization, and an error, if there is any.
func (c *FakeApigeeOrganizations) Create(ctx context.Context, apigeeOrganization *v1beta1.ApigeeOrganization, opts v1.CreateOptions) (result *v1beta1.ApigeeOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigeeorganizationsResource, c.ns, apigeeOrganization), &v1beta1.ApigeeOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ApigeeOrganization), err
}

// Update takes the representation of a apigeeOrganization and updates it. Returns the server's representation of the apigeeOrganization, and an error, if there is any.
func (c *FakeApigeeOrganizations) Update(ctx context.Context, apigeeOrganization *v1beta1.ApigeeOrganization, opts v1.UpdateOptions) (result *v1beta1.ApigeeOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigeeorganizationsResource, c.ns, apigeeOrganization), &v1beta1.ApigeeOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ApigeeOrganization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApigeeOrganizations) UpdateStatus(ctx context.Context, apigeeOrganization *v1beta1.ApigeeOrganization, opts v1.UpdateOptions) (*v1beta1.ApigeeOrganization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigeeorganizationsResource, "status", c.ns, apigeeOrganization), &v1beta1.ApigeeOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ApigeeOrganization), err
}

// Delete takes name of the apigeeOrganization and deletes it. Returns an error if one occurs.
func (c *FakeApigeeOrganizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apigeeorganizationsResource, c.ns, name, opts), &v1beta1.ApigeeOrganization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApigeeOrganizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigeeorganizationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ApigeeOrganizationList{})
	return err
}

// Patch applies the patch and returns the patched apigeeOrganization.
func (c *FakeApigeeOrganizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ApigeeOrganization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigeeorganizationsResource, c.ns, name, pt, data, subresources...), &v1beta1.ApigeeOrganization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ApigeeOrganization), err
}
