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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeManagedSSLCertificates implements ComputeManagedSSLCertificateInterface
type FakeComputeManagedSSLCertificates struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var computemanagedsslcertificatesResource = v1alpha1.SchemeGroupVersion.WithResource("computemanagedsslcertificates")

var computemanagedsslcertificatesKind = v1alpha1.SchemeGroupVersion.WithKind("ComputeManagedSSLCertificate")

// Get takes name of the computeManagedSSLCertificate, and returns the corresponding computeManagedSSLCertificate object, and an error if there is any.
func (c *FakeComputeManagedSSLCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeManagedSSLCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computemanagedsslcertificatesResource, c.ns, name), &v1alpha1.ComputeManagedSSLCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeManagedSSLCertificate), err
}

// List takes label and field selectors, and returns the list of ComputeManagedSSLCertificates that match those selectors.
func (c *FakeComputeManagedSSLCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeManagedSSLCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computemanagedsslcertificatesResource, computemanagedsslcertificatesKind, c.ns, opts), &v1alpha1.ComputeManagedSSLCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeManagedSSLCertificateList{ListMeta: obj.(*v1alpha1.ComputeManagedSSLCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeManagedSSLCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeManagedSSLCertificates.
func (c *FakeComputeManagedSSLCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computemanagedsslcertificatesResource, c.ns, opts))

}

// Create takes the representation of a computeManagedSSLCertificate and creates it.  Returns the server's representation of the computeManagedSSLCertificate, and an error, if there is any.
func (c *FakeComputeManagedSSLCertificates) Create(ctx context.Context, computeManagedSSLCertificate *v1alpha1.ComputeManagedSSLCertificate, opts v1.CreateOptions) (result *v1alpha1.ComputeManagedSSLCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computemanagedsslcertificatesResource, c.ns, computeManagedSSLCertificate), &v1alpha1.ComputeManagedSSLCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeManagedSSLCertificate), err
}

// Update takes the representation of a computeManagedSSLCertificate and updates it. Returns the server's representation of the computeManagedSSLCertificate, and an error, if there is any.
func (c *FakeComputeManagedSSLCertificates) Update(ctx context.Context, computeManagedSSLCertificate *v1alpha1.ComputeManagedSSLCertificate, opts v1.UpdateOptions) (result *v1alpha1.ComputeManagedSSLCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computemanagedsslcertificatesResource, c.ns, computeManagedSSLCertificate), &v1alpha1.ComputeManagedSSLCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeManagedSSLCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeManagedSSLCertificates) UpdateStatus(ctx context.Context, computeManagedSSLCertificate *v1alpha1.ComputeManagedSSLCertificate, opts v1.UpdateOptions) (*v1alpha1.ComputeManagedSSLCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computemanagedsslcertificatesResource, "status", c.ns, computeManagedSSLCertificate), &v1alpha1.ComputeManagedSSLCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeManagedSSLCertificate), err
}

// Delete takes name of the computeManagedSSLCertificate and deletes it. Returns an error if one occurs.
func (c *FakeComputeManagedSSLCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computemanagedsslcertificatesResource, c.ns, name, opts), &v1alpha1.ComputeManagedSSLCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeManagedSSLCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computemanagedsslcertificatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeManagedSSLCertificateList{})
	return err
}

// Patch applies the patch and returns the patched computeManagedSSLCertificate.
func (c *FakeComputeManagedSSLCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeManagedSSLCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computemanagedsslcertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeManagedSSLCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeManagedSSLCertificate), err
}
