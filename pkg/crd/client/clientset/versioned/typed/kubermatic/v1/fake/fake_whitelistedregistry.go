// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWhitelistedRegistries implements WhitelistedRegistryInterface
type FakeWhitelistedRegistries struct {
	Fake *FakeKubermaticV1
}

var whitelistedregistriesResource = schema.GroupVersionResource{Group: "kubermatic.k8s.io", Version: "v1", Resource: "whitelistedregistries"}

var whitelistedregistriesKind = schema.GroupVersionKind{Group: "kubermatic.k8s.io", Version: "v1", Kind: "WhitelistedRegistry"}

// Get takes name of the whitelistedRegistry, and returns the corresponding whitelistedRegistry object, and an error if there is any.
func (c *FakeWhitelistedRegistries) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.WhitelistedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(whitelistedregistriesResource, name), &kubermaticv1.WhitelistedRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.WhitelistedRegistry), err
}

// List takes label and field selectors, and returns the list of WhitelistedRegistries that match those selectors.
func (c *FakeWhitelistedRegistries) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.WhitelistedRegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(whitelistedregistriesResource, whitelistedregistriesKind, opts), &kubermaticv1.WhitelistedRegistryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.WhitelistedRegistryList{ListMeta: obj.(*kubermaticv1.WhitelistedRegistryList).ListMeta}
	for _, item := range obj.(*kubermaticv1.WhitelistedRegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested whitelistedRegistries.
func (c *FakeWhitelistedRegistries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(whitelistedregistriesResource, opts))
}

// Create takes the representation of a whitelistedRegistry and creates it.  Returns the server's representation of the whitelistedRegistry, and an error, if there is any.
func (c *FakeWhitelistedRegistries) Create(ctx context.Context, whitelistedRegistry *kubermaticv1.WhitelistedRegistry, opts v1.CreateOptions) (result *kubermaticv1.WhitelistedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(whitelistedregistriesResource, whitelistedRegistry), &kubermaticv1.WhitelistedRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.WhitelistedRegistry), err
}

// Update takes the representation of a whitelistedRegistry and updates it. Returns the server's representation of the whitelistedRegistry, and an error, if there is any.
func (c *FakeWhitelistedRegistries) Update(ctx context.Context, whitelistedRegistry *kubermaticv1.WhitelistedRegistry, opts v1.UpdateOptions) (result *kubermaticv1.WhitelistedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(whitelistedregistriesResource, whitelistedRegistry), &kubermaticv1.WhitelistedRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.WhitelistedRegistry), err
}

// Delete takes name of the whitelistedRegistry and deletes it. Returns an error if one occurs.
func (c *FakeWhitelistedRegistries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(whitelistedregistriesResource, name), &kubermaticv1.WhitelistedRegistry{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWhitelistedRegistries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(whitelistedregistriesResource, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.WhitelistedRegistryList{})
	return err
}

// Patch applies the patch and returns the patched whitelistedRegistry.
func (c *FakeWhitelistedRegistries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.WhitelistedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(whitelistedregistriesResource, name, pt, data, subresources...), &kubermaticv1.WhitelistedRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.WhitelistedRegistry), err
}
