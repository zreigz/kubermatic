// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	kubermaticv1 "github.com/kubermatic/kubermatic/pkg/crd/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImportedClusters implements ImportedClusterInterface
type FakeImportedClusters struct {
	Fake *FakeKubermaticV1
}

var importedclustersResource = schema.GroupVersionResource{Group: "kubermatic.k8s.io", Version: "v1", Resource: "importedclusters"}

var importedclustersKind = schema.GroupVersionKind{Group: "kubermatic.k8s.io", Version: "v1", Kind: "ImportedCluster"}

// Get takes name of the importedCluster, and returns the corresponding importedCluster object, and an error if there is any.
func (c *FakeImportedClusters) Get(name string, options v1.GetOptions) (result *kubermaticv1.ImportedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(importedclustersResource, name), &kubermaticv1.ImportedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ImportedCluster), err
}

// List takes label and field selectors, and returns the list of ImportedClusters that match those selectors.
func (c *FakeImportedClusters) List(opts v1.ListOptions) (result *kubermaticv1.ImportedClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(importedclustersResource, importedclustersKind, opts), &kubermaticv1.ImportedClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.ImportedClusterList{ListMeta: obj.(*kubermaticv1.ImportedClusterList).ListMeta}
	for _, item := range obj.(*kubermaticv1.ImportedClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested importedClusters.
func (c *FakeImportedClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(importedclustersResource, opts))
}

// Create takes the representation of a importedCluster and creates it.  Returns the server's representation of the importedCluster, and an error, if there is any.
func (c *FakeImportedClusters) Create(importedCluster *kubermaticv1.ImportedCluster) (result *kubermaticv1.ImportedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(importedclustersResource, importedCluster), &kubermaticv1.ImportedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ImportedCluster), err
}

// Update takes the representation of a importedCluster and updates it. Returns the server's representation of the importedCluster, and an error, if there is any.
func (c *FakeImportedClusters) Update(importedCluster *kubermaticv1.ImportedCluster) (result *kubermaticv1.ImportedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(importedclustersResource, importedCluster), &kubermaticv1.ImportedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ImportedCluster), err
}

// Delete takes name of the importedCluster and deletes it. Returns an error if one occurs.
func (c *FakeImportedClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(importedclustersResource, name), &kubermaticv1.ImportedCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImportedClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(importedclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &kubermaticv1.ImportedClusterList{})
	return err
}

// Patch applies the patch and returns the patched importedCluster.
func (c *FakeImportedClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kubermaticv1.ImportedCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(importedclustersResource, name, pt, data, subresources...), &kubermaticv1.ImportedCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ImportedCluster), err
}
