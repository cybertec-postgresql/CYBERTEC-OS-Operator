/*
Copyright 2024 Compose, Zalando SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/cybertec-postgresql/cybertec-pg-operator/pkg/apis/zalando.org/v1"
	scheme "github.com/cybertec-postgresql/cybertec-pg-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FabricEventStreamsGetter has a method to return a FabricEventStreamInterface.
// A group's client should implement this interface.
type FabricEventStreamsGetter interface {
	FabricEventStreams(namespace string) FabricEventStreamInterface
}

// FabricEventStreamInterface has methods to work with FabricEventStream resources.
type FabricEventStreamInterface interface {
	Create(ctx context.Context, fabricEventStream *v1.FabricEventStream, opts metav1.CreateOptions) (*v1.FabricEventStream, error)
	Update(ctx context.Context, fabricEventStream *v1.FabricEventStream, opts metav1.UpdateOptions) (*v1.FabricEventStream, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.FabricEventStream, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.FabricEventStreamList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FabricEventStream, err error)
	FabricEventStreamExpansion
}

// fabricEventStreams implements FabricEventStreamInterface
type fabricEventStreams struct {
	client rest.Interface
	ns     string
}

// newFabricEventStreams returns a FabricEventStreams
func newFabricEventStreams(c *ZalandoV1Client, namespace string) *fabricEventStreams {
	return &fabricEventStreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fabricEventStream, and returns the corresponding fabricEventStream object, and an error if there is any.
func (c *fabricEventStreams) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.FabricEventStream, err error) {
	result = &v1.FabricEventStream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fabriceventstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FabricEventStreams that match those selectors.
func (c *fabricEventStreams) List(ctx context.Context, opts metav1.ListOptions) (result *v1.FabricEventStreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.FabricEventStreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fabriceventstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fabricEventStreams.
func (c *fabricEventStreams) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fabriceventstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fabricEventStream and creates it.  Returns the server's representation of the fabricEventStream, and an error, if there is any.
func (c *fabricEventStreams) Create(ctx context.Context, fabricEventStream *v1.FabricEventStream, opts metav1.CreateOptions) (result *v1.FabricEventStream, err error) {
	result = &v1.FabricEventStream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fabriceventstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricEventStream).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fabricEventStream and updates it. Returns the server's representation of the fabricEventStream, and an error, if there is any.
func (c *fabricEventStreams) Update(ctx context.Context, fabricEventStream *v1.FabricEventStream, opts metav1.UpdateOptions) (result *v1.FabricEventStream, err error) {
	result = &v1.FabricEventStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fabriceventstreams").
		Name(fabricEventStream.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fabricEventStream).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fabricEventStream and deletes it. Returns an error if one occurs.
func (c *fabricEventStreams) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fabriceventstreams").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fabricEventStreams) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fabriceventstreams").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fabricEventStream.
func (c *fabricEventStreams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FabricEventStream, err error) {
	result = &v1.FabricEventStream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fabriceventstreams").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
