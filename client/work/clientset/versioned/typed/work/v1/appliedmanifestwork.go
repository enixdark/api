// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "github.com/open-cluster-management/api/client/work/clientset/versioned/scheme"
	v1 "github.com/open-cluster-management/api/work/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppliedManifestWorksGetter has a method to return a AppliedManifestWorkInterface.
// A group's client should implement this interface.
type AppliedManifestWorksGetter interface {
	AppliedManifestWorks() AppliedManifestWorkInterface
}

// AppliedManifestWorkInterface has methods to work with AppliedManifestWork resources.
type AppliedManifestWorkInterface interface {
	Create(ctx context.Context, appliedManifestWork *v1.AppliedManifestWork, opts metav1.CreateOptions) (*v1.AppliedManifestWork, error)
	Update(ctx context.Context, appliedManifestWork *v1.AppliedManifestWork, opts metav1.UpdateOptions) (*v1.AppliedManifestWork, error)
	UpdateStatus(ctx context.Context, appliedManifestWork *v1.AppliedManifestWork, opts metav1.UpdateOptions) (*v1.AppliedManifestWork, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AppliedManifestWork, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AppliedManifestWorkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AppliedManifestWork, err error)
	AppliedManifestWorkExpansion
}

// appliedManifestWorks implements AppliedManifestWorkInterface
type appliedManifestWorks struct {
	client rest.Interface
}

// newAppliedManifestWorks returns a AppliedManifestWorks
func newAppliedManifestWorks(c *WorkV1Client) *appliedManifestWorks {
	return &appliedManifestWorks{
		client: c.RESTClient(),
	}
}

// Get takes name of the appliedManifestWork, and returns the corresponding appliedManifestWork object, and an error if there is any.
func (c *appliedManifestWorks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AppliedManifestWork, err error) {
	result = &v1.AppliedManifestWork{}
	err = c.client.Get().
		Resource("appliedmanifestworks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppliedManifestWorks that match those selectors.
func (c *appliedManifestWorks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AppliedManifestWorkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AppliedManifestWorkList{}
	err = c.client.Get().
		Resource("appliedmanifestworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appliedManifestWorks.
func (c *appliedManifestWorks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("appliedmanifestworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a appliedManifestWork and creates it.  Returns the server's representation of the appliedManifestWork, and an error, if there is any.
func (c *appliedManifestWorks) Create(ctx context.Context, appliedManifestWork *v1.AppliedManifestWork, opts metav1.CreateOptions) (result *v1.AppliedManifestWork, err error) {
	result = &v1.AppliedManifestWork{}
	err = c.client.Post().
		Resource("appliedmanifestworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appliedManifestWork).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a appliedManifestWork and updates it. Returns the server's representation of the appliedManifestWork, and an error, if there is any.
func (c *appliedManifestWorks) Update(ctx context.Context, appliedManifestWork *v1.AppliedManifestWork, opts metav1.UpdateOptions) (result *v1.AppliedManifestWork, err error) {
	result = &v1.AppliedManifestWork{}
	err = c.client.Put().
		Resource("appliedmanifestworks").
		Name(appliedManifestWork.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appliedManifestWork).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *appliedManifestWorks) UpdateStatus(ctx context.Context, appliedManifestWork *v1.AppliedManifestWork, opts metav1.UpdateOptions) (result *v1.AppliedManifestWork, err error) {
	result = &v1.AppliedManifestWork{}
	err = c.client.Put().
		Resource("appliedmanifestworks").
		Name(appliedManifestWork.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appliedManifestWork).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the appliedManifestWork and deletes it. Returns an error if one occurs.
func (c *appliedManifestWorks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("appliedmanifestworks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appliedManifestWorks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("appliedmanifestworks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched appliedManifestWork.
func (c *appliedManifestWorks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AppliedManifestWork, err error) {
	result = &v1.AppliedManifestWork{}
	err = c.client.Patch(pt).
		Resource("appliedmanifestworks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
