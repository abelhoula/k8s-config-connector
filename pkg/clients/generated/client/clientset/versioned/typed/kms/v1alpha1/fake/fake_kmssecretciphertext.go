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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/kms/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKMSSecretCiphertexts implements KMSSecretCiphertextInterface
type FakeKMSSecretCiphertexts struct {
	Fake *FakeKmsV1alpha1
	ns   string
}

var kmssecretciphertextsResource = v1alpha1.SchemeGroupVersion.WithResource("kmssecretciphertexts")

var kmssecretciphertextsKind = v1alpha1.SchemeGroupVersion.WithKind("KMSSecretCiphertext")

// Get takes name of the kMSSecretCiphertext, and returns the corresponding kMSSecretCiphertext object, and an error if there is any.
func (c *FakeKMSSecretCiphertexts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KMSSecretCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kmssecretciphertextsResource, c.ns, name), &v1alpha1.KMSSecretCiphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSSecretCiphertext), err
}

// List takes label and field selectors, and returns the list of KMSSecretCiphertexts that match those selectors.
func (c *FakeKMSSecretCiphertexts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KMSSecretCiphertextList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kmssecretciphertextsResource, kmssecretciphertextsKind, c.ns, opts), &v1alpha1.KMSSecretCiphertextList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KMSSecretCiphertextList{ListMeta: obj.(*v1alpha1.KMSSecretCiphertextList).ListMeta}
	for _, item := range obj.(*v1alpha1.KMSSecretCiphertextList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kMSSecretCiphertexts.
func (c *FakeKMSSecretCiphertexts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kmssecretciphertextsResource, c.ns, opts))

}

// Create takes the representation of a kMSSecretCiphertext and creates it.  Returns the server's representation of the kMSSecretCiphertext, and an error, if there is any.
func (c *FakeKMSSecretCiphertexts) Create(ctx context.Context, kMSSecretCiphertext *v1alpha1.KMSSecretCiphertext, opts v1.CreateOptions) (result *v1alpha1.KMSSecretCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kmssecretciphertextsResource, c.ns, kMSSecretCiphertext), &v1alpha1.KMSSecretCiphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSSecretCiphertext), err
}

// Update takes the representation of a kMSSecretCiphertext and updates it. Returns the server's representation of the kMSSecretCiphertext, and an error, if there is any.
func (c *FakeKMSSecretCiphertexts) Update(ctx context.Context, kMSSecretCiphertext *v1alpha1.KMSSecretCiphertext, opts v1.UpdateOptions) (result *v1alpha1.KMSSecretCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kmssecretciphertextsResource, c.ns, kMSSecretCiphertext), &v1alpha1.KMSSecretCiphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSSecretCiphertext), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKMSSecretCiphertexts) UpdateStatus(ctx context.Context, kMSSecretCiphertext *v1alpha1.KMSSecretCiphertext, opts v1.UpdateOptions) (*v1alpha1.KMSSecretCiphertext, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kmssecretciphertextsResource, "status", c.ns, kMSSecretCiphertext), &v1alpha1.KMSSecretCiphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSSecretCiphertext), err
}

// Delete takes name of the kMSSecretCiphertext and deletes it. Returns an error if one occurs.
func (c *FakeKMSSecretCiphertexts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kmssecretciphertextsResource, c.ns, name, opts), &v1alpha1.KMSSecretCiphertext{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKMSSecretCiphertexts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kmssecretciphertextsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KMSSecretCiphertextList{})
	return err
}

// Patch applies the patch and returns the patched kMSSecretCiphertext.
func (c *FakeKMSSecretCiphertexts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KMSSecretCiphertext, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kmssecretciphertextsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KMSSecretCiphertext{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KMSSecretCiphertext), err
}
