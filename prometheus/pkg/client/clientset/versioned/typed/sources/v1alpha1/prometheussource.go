/*
Copyright 2020 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-contrib/prometheus/pkg/apis/sources/v1alpha1"
	scheme "knative.dev/eventing-contrib/prometheus/pkg/client/clientset/versioned/scheme"
)

// PrometheusSourcesGetter has a method to return a PrometheusSourceInterface.
// A group's client should implement this interface.
type PrometheusSourcesGetter interface {
	PrometheusSources(namespace string) PrometheusSourceInterface
}

// PrometheusSourceInterface has methods to work with PrometheusSource resources.
type PrometheusSourceInterface interface {
	Create(*v1alpha1.PrometheusSource) (*v1alpha1.PrometheusSource, error)
	Update(*v1alpha1.PrometheusSource) (*v1alpha1.PrometheusSource, error)
	UpdateStatus(*v1alpha1.PrometheusSource) (*v1alpha1.PrometheusSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PrometheusSource, error)
	List(opts v1.ListOptions) (*v1alpha1.PrometheusSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrometheusSource, err error)
	PrometheusSourceExpansion
}

// prometheusSources implements PrometheusSourceInterface
type prometheusSources struct {
	client rest.Interface
	ns     string
}

// newPrometheusSources returns a PrometheusSources
func newPrometheusSources(c *SourcesV1alpha1Client, namespace string) *prometheusSources {
	return &prometheusSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the prometheusSource, and returns the corresponding prometheusSource object, and an error if there is any.
func (c *prometheusSources) Get(name string, options v1.GetOptions) (result *v1alpha1.PrometheusSource, err error) {
	result = &v1alpha1.PrometheusSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheussources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PrometheusSources that match those selectors.
func (c *prometheusSources) List(opts v1.ListOptions) (result *v1alpha1.PrometheusSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PrometheusSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("prometheussources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested prometheusSources.
func (c *prometheusSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("prometheussources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a prometheusSource and creates it.  Returns the server's representation of the prometheusSource, and an error, if there is any.
func (c *prometheusSources) Create(prometheusSource *v1alpha1.PrometheusSource) (result *v1alpha1.PrometheusSource, err error) {
	result = &v1alpha1.PrometheusSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("prometheussources").
		Body(prometheusSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a prometheusSource and updates it. Returns the server's representation of the prometheusSource, and an error, if there is any.
func (c *prometheusSources) Update(prometheusSource *v1alpha1.PrometheusSource) (result *v1alpha1.PrometheusSource, err error) {
	result = &v1alpha1.PrometheusSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prometheussources").
		Name(prometheusSource.Name).
		Body(prometheusSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *prometheusSources) UpdateStatus(prometheusSource *v1alpha1.PrometheusSource) (result *v1alpha1.PrometheusSource, err error) {
	result = &v1alpha1.PrometheusSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("prometheussources").
		Name(prometheusSource.Name).
		SubResource("status").
		Body(prometheusSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the prometheusSource and deletes it. Returns an error if one occurs.
func (c *prometheusSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheussources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *prometheusSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("prometheussources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched prometheusSource.
func (c *prometheusSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrometheusSource, err error) {
	result = &v1alpha1.PrometheusSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("prometheussources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
