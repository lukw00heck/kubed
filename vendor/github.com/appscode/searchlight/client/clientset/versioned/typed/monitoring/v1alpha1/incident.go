/*
Copyright 2018 The Searchlight Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	scheme "github.com/appscode/searchlight/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IncidentsGetter has a method to return a IncidentInterface.
// A group's client should implement this interface.
type IncidentsGetter interface {
	Incidents(namespace string) IncidentInterface
}

// IncidentInterface has methods to work with Incident resources.
type IncidentInterface interface {
	Create(*v1alpha1.Incident) (*v1alpha1.Incident, error)
	Update(*v1alpha1.Incident) (*v1alpha1.Incident, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Incident, error)
	List(opts v1.ListOptions) (*v1alpha1.IncidentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Incident, err error)
	IncidentExpansion
}

// incidents implements IncidentInterface
type incidents struct {
	client rest.Interface
	ns     string
}

// newIncidents returns a Incidents
func newIncidents(c *MonitoringV1alpha1Client, namespace string) *incidents {
	return &incidents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the incident, and returns the corresponding incident object, and an error if there is any.
func (c *incidents) Get(name string, options v1.GetOptions) (result *v1alpha1.Incident, err error) {
	result = &v1alpha1.Incident{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("incidents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Incidents that match those selectors.
func (c *incidents) List(opts v1.ListOptions) (result *v1alpha1.IncidentList, err error) {
	result = &v1alpha1.IncidentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("incidents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested incidents.
func (c *incidents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("incidents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a incident and creates it.  Returns the server's representation of the incident, and an error, if there is any.
func (c *incidents) Create(incident *v1alpha1.Incident) (result *v1alpha1.Incident, err error) {
	result = &v1alpha1.Incident{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("incidents").
		Body(incident).
		Do().
		Into(result)
	return
}

// Update takes the representation of a incident and updates it. Returns the server's representation of the incident, and an error, if there is any.
func (c *incidents) Update(incident *v1alpha1.Incident) (result *v1alpha1.Incident, err error) {
	result = &v1alpha1.Incident{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("incidents").
		Name(incident.Name).
		Body(incident).
		Do().
		Into(result)
	return
}

// Delete takes name of the incident and deletes it. Returns an error if one occurs.
func (c *incidents) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("incidents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *incidents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("incidents").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched incident.
func (c *incidents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Incident, err error) {
	result = &v1alpha1.Incident{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("incidents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
