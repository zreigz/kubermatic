// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConstraintTemplateLister helps list ConstraintTemplates.
// All objects returned here must be treated as read-only.
type ConstraintTemplateLister interface {
	// List lists all ConstraintTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ConstraintTemplate, err error)
	// Get retrieves the ConstraintTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ConstraintTemplate, error)
	ConstraintTemplateListerExpansion
}

// constraintTemplateLister implements the ConstraintTemplateLister interface.
type constraintTemplateLister struct {
	indexer cache.Indexer
}

// NewConstraintTemplateLister returns a new ConstraintTemplateLister.
func NewConstraintTemplateLister(indexer cache.Indexer) ConstraintTemplateLister {
	return &constraintTemplateLister{indexer: indexer}
}

// List lists all ConstraintTemplates in the indexer.
func (s *constraintTemplateLister) List(selector labels.Selector) (ret []*v1.ConstraintTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ConstraintTemplate))
	})
	return ret, err
}

// Get retrieves the ConstraintTemplate from the index for a given name.
func (s *constraintTemplateLister) Get(name string) (*v1.ConstraintTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("constrainttemplate"), name)
	}
	return obj.(*v1.ConstraintTemplate), nil
}