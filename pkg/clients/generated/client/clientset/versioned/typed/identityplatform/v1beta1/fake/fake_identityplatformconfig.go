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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityPlatformConfigs implements IdentityPlatformConfigInterface
type FakeIdentityPlatformConfigs struct {
	Fake *FakeIdentityplatformV1beta1
	ns   string
}

var identityplatformconfigsResource = v1beta1.SchemeGroupVersion.WithResource("identityplatformconfigs")

var identityplatformconfigsKind = v1beta1.SchemeGroupVersion.WithKind("IdentityPlatformConfig")

// Get takes name of the identityPlatformConfig, and returns the corresponding identityPlatformConfig object, and an error if there is any.
func (c *FakeIdentityPlatformConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IdentityPlatformConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityplatformconfigsResource, c.ns, name), &v1beta1.IdentityPlatformConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformConfig), err
}

// List takes label and field selectors, and returns the list of IdentityPlatformConfigs that match those selectors.
func (c *FakeIdentityPlatformConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IdentityPlatformConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityplatformconfigsResource, identityplatformconfigsKind, c.ns, opts), &v1beta1.IdentityPlatformConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IdentityPlatformConfigList{ListMeta: obj.(*v1beta1.IdentityPlatformConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.IdentityPlatformConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityPlatformConfigs.
func (c *FakeIdentityPlatformConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityplatformconfigsResource, c.ns, opts))

}

// Create takes the representation of a identityPlatformConfig and creates it.  Returns the server's representation of the identityPlatformConfig, and an error, if there is any.
func (c *FakeIdentityPlatformConfigs) Create(ctx context.Context, identityPlatformConfig *v1beta1.IdentityPlatformConfig, opts v1.CreateOptions) (result *v1beta1.IdentityPlatformConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityplatformconfigsResource, c.ns, identityPlatformConfig), &v1beta1.IdentityPlatformConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformConfig), err
}

// Update takes the representation of a identityPlatformConfig and updates it. Returns the server's representation of the identityPlatformConfig, and an error, if there is any.
func (c *FakeIdentityPlatformConfigs) Update(ctx context.Context, identityPlatformConfig *v1beta1.IdentityPlatformConfig, opts v1.UpdateOptions) (result *v1beta1.IdentityPlatformConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityplatformconfigsResource, c.ns, identityPlatformConfig), &v1beta1.IdentityPlatformConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityPlatformConfigs) UpdateStatus(ctx context.Context, identityPlatformConfig *v1beta1.IdentityPlatformConfig, opts v1.UpdateOptions) (*v1beta1.IdentityPlatformConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityplatformconfigsResource, "status", c.ns, identityPlatformConfig), &v1beta1.IdentityPlatformConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformConfig), err
}

// Delete takes name of the identityPlatformConfig and deletes it. Returns an error if one occurs.
func (c *FakeIdentityPlatformConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(identityplatformconfigsResource, c.ns, name, opts), &v1beta1.IdentityPlatformConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityPlatformConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityplatformconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IdentityPlatformConfigList{})
	return err
}

// Patch applies the patch and returns the patched identityPlatformConfig.
func (c *FakeIdentityPlatformConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IdentityPlatformConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityplatformconfigsResource, c.ns, name, pt, data, subresources...), &v1beta1.IdentityPlatformConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformConfig), err
}
