/*
Copyright 2018 The Voyager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "k8s-admission-webhook-with-extension-apiserver/apis/kubecar/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubecarLister helps list Kubecars.
type KubecarLister interface {
	// List lists all Kubecars in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Kubecar, err error)
	// Kubecars returns an object that can list and get Kubecars.
	Kubecars(namespace string) KubecarNamespaceLister
	KubecarListerExpansion
}

// kubecarLister implements the KubecarLister interface.
type kubecarLister struct {
	indexer cache.Indexer
}

// NewKubecarLister returns a new KubecarLister.
func NewKubecarLister(indexer cache.Indexer) KubecarLister {
	return &kubecarLister{indexer: indexer}
}

// List lists all Kubecars in the indexer.
func (s *kubecarLister) List(selector labels.Selector) (ret []*v1alpha1.Kubecar, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Kubecar))
	})
	return ret, err
}

// Kubecars returns an object that can list and get Kubecars.
func (s *kubecarLister) Kubecars(namespace string) KubecarNamespaceLister {
	return kubecarNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubecarNamespaceLister helps list and get Kubecars.
type KubecarNamespaceLister interface {
	// List lists all Kubecars in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Kubecar, err error)
	// Get retrieves the Kubecar from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Kubecar, error)
	KubecarNamespaceListerExpansion
}

// kubecarNamespaceLister implements the KubecarNamespaceLister
// interface.
type kubecarNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Kubecars in the indexer for a given namespace.
func (s kubecarNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Kubecar, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Kubecar))
	})
	return ret, err
}

// Get retrieves the Kubecar from the indexer for a given namespace and name.
func (s kubecarNamespaceLister) Get(name string) (*v1alpha1.Kubecar, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubecar"), name)
	}
	return obj.(*v1alpha1.Kubecar), nil
}
