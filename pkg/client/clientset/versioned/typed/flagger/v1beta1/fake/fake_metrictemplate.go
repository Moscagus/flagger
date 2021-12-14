/*
Copyright 2020 The Flux authors

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

package fake

import (
	"context"

	v1beta1 "github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMetricTemplates implements MetricTemplateInterface
type FakeMetricTemplates struct {
	Fake *FakeFlaggerV1beta1
	ns   string
}

var metrictemplatesResource = schema.GroupVersionResource{Group: "flagger.app", Version: "v1beta1", Resource: "metrictemplates"}

var metrictemplatesKind = schema.GroupVersionKind{Group: "flagger.app", Version: "v1beta1", Kind: "MetricTemplate"}

// Get takes name of the metricTemplate, and returns the corresponding metricTemplate object, and an error if there is any.
func (c *FakeMetricTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MetricTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metrictemplatesResource, c.ns, name), &v1beta1.MetricTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MetricTemplate), err
}

// List takes label and field selectors, and returns the list of MetricTemplates that match those selectors.
func (c *FakeMetricTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MetricTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metrictemplatesResource, metrictemplatesKind, c.ns, opts), &v1beta1.MetricTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MetricTemplateList{ListMeta: obj.(*v1beta1.MetricTemplateList).ListMeta}
	for _, item := range obj.(*v1beta1.MetricTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metricTemplates.
func (c *FakeMetricTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metrictemplatesResource, c.ns, opts))

}

// Create takes the representation of a metricTemplate and creates it.  Returns the server's representation of the metricTemplate, and an error, if there is any.
func (c *FakeMetricTemplates) Create(ctx context.Context, metricTemplate *v1beta1.MetricTemplate, opts v1.CreateOptions) (result *v1beta1.MetricTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metrictemplatesResource, c.ns, metricTemplate), &v1beta1.MetricTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MetricTemplate), err
}

// Update takes the representation of a metricTemplate and updates it. Returns the server's representation of the metricTemplate, and an error, if there is any.
func (c *FakeMetricTemplates) Update(ctx context.Context, metricTemplate *v1beta1.MetricTemplate, opts v1.UpdateOptions) (result *v1beta1.MetricTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metrictemplatesResource, c.ns, metricTemplate), &v1beta1.MetricTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MetricTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetricTemplates) UpdateStatus(ctx context.Context, metricTemplate *v1beta1.MetricTemplate, opts v1.UpdateOptions) (*v1beta1.MetricTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metrictemplatesResource, "status", c.ns, metricTemplate), &v1beta1.MetricTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MetricTemplate), err
}

// Delete takes name of the metricTemplate and deletes it. Returns an error if one occurs.
func (c *FakeMetricTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(metrictemplatesResource, c.ns, name, opts), &v1beta1.MetricTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetricTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metrictemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MetricTemplateList{})
	return err
}

// Patch applies the patch and returns the patched metricTemplate.
func (c *FakeMetricTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MetricTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metrictemplatesResource, c.ns, name, pt, data, subresources...), &v1beta1.MetricTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MetricTemplate), err
}
