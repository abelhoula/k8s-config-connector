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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/appengine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppEngineFlexibleAppVersions implements AppEngineFlexibleAppVersionInterface
type FakeAppEngineFlexibleAppVersions struct {
	Fake *FakeAppengineV1alpha1
	ns   string
}

var appengineflexibleappversionsResource = v1alpha1.SchemeGroupVersion.WithResource("appengineflexibleappversions")

var appengineflexibleappversionsKind = v1alpha1.SchemeGroupVersion.WithKind("AppEngineFlexibleAppVersion")

// Get takes name of the appEngineFlexibleAppVersion, and returns the corresponding appEngineFlexibleAppVersion object, and an error if there is any.
func (c *FakeAppEngineFlexibleAppVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppEngineFlexibleAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appengineflexibleappversionsResource, c.ns, name), &v1alpha1.AppEngineFlexibleAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFlexibleAppVersion), err
}

// List takes label and field selectors, and returns the list of AppEngineFlexibleAppVersions that match those selectors.
func (c *FakeAppEngineFlexibleAppVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppEngineFlexibleAppVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appengineflexibleappversionsResource, appengineflexibleappversionsKind, c.ns, opts), &v1alpha1.AppEngineFlexibleAppVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppEngineFlexibleAppVersionList{ListMeta: obj.(*v1alpha1.AppEngineFlexibleAppVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppEngineFlexibleAppVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appEngineFlexibleAppVersions.
func (c *FakeAppEngineFlexibleAppVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appengineflexibleappversionsResource, c.ns, opts))

}

// Create takes the representation of a appEngineFlexibleAppVersion and creates it.  Returns the server's representation of the appEngineFlexibleAppVersion, and an error, if there is any.
func (c *FakeAppEngineFlexibleAppVersions) Create(ctx context.Context, appEngineFlexibleAppVersion *v1alpha1.AppEngineFlexibleAppVersion, opts v1.CreateOptions) (result *v1alpha1.AppEngineFlexibleAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appengineflexibleappversionsResource, c.ns, appEngineFlexibleAppVersion), &v1alpha1.AppEngineFlexibleAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFlexibleAppVersion), err
}

// Update takes the representation of a appEngineFlexibleAppVersion and updates it. Returns the server's representation of the appEngineFlexibleAppVersion, and an error, if there is any.
func (c *FakeAppEngineFlexibleAppVersions) Update(ctx context.Context, appEngineFlexibleAppVersion *v1alpha1.AppEngineFlexibleAppVersion, opts v1.UpdateOptions) (result *v1alpha1.AppEngineFlexibleAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appengineflexibleappversionsResource, c.ns, appEngineFlexibleAppVersion), &v1alpha1.AppEngineFlexibleAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFlexibleAppVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppEngineFlexibleAppVersions) UpdateStatus(ctx context.Context, appEngineFlexibleAppVersion *v1alpha1.AppEngineFlexibleAppVersion, opts v1.UpdateOptions) (*v1alpha1.AppEngineFlexibleAppVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appengineflexibleappversionsResource, "status", c.ns, appEngineFlexibleAppVersion), &v1alpha1.AppEngineFlexibleAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFlexibleAppVersion), err
}

// Delete takes name of the appEngineFlexibleAppVersion and deletes it. Returns an error if one occurs.
func (c *FakeAppEngineFlexibleAppVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(appengineflexibleappversionsResource, c.ns, name, opts), &v1alpha1.AppEngineFlexibleAppVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppEngineFlexibleAppVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appengineflexibleappversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppEngineFlexibleAppVersionList{})
	return err
}

// Patch applies the patch and returns the patched appEngineFlexibleAppVersion.
func (c *FakeAppEngineFlexibleAppVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppEngineFlexibleAppVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appengineflexibleappversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppEngineFlexibleAppVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppEngineFlexibleAppVersion), err
}
