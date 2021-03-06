/*
Copyright 2020 The SuperEdge Authors.

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

	v1beta1 "github.com/superedge/superedge/pkg/penetrator/apis/nodetask.apps.superedge.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeTasks implements NodeTaskInterface
type FakeNodeTasks struct {
	Fake *FakeNodestaskV1beta1
}

var nodetasksResource = schema.GroupVersionResource{Group: "nodestask.tkeedge.io", Version: "v1beta1", Resource: "nodetasks"}

var nodetasksKind = schema.GroupVersionKind{Group: "nodestask.tkeedge.io", Version: "v1beta1", Kind: "NodeTask"}

// Get takes name of the nodeTask, and returns the corresponding nodeTask object, and an error if there is any.
func (c *FakeNodeTasks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NodeTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodetasksResource, name), &v1beta1.NodeTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NodeTask), err
}

// List takes label and field selectors, and returns the list of NodeTasks that match those selectors.
func (c *FakeNodeTasks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NodeTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodetasksResource, nodetasksKind, opts), &v1beta1.NodeTaskList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NodeTaskList{ListMeta: obj.(*v1beta1.NodeTaskList).ListMeta}
	for _, item := range obj.(*v1beta1.NodeTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeTasks.
func (c *FakeNodeTasks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodetasksResource, opts))
}

// Create takes the representation of a nodeTask and creates it.  Returns the server's representation of the nodeTask, and an error, if there is any.
func (c *FakeNodeTasks) Create(ctx context.Context, nodeTask *v1beta1.NodeTask, opts v1.CreateOptions) (result *v1beta1.NodeTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodetasksResource, nodeTask), &v1beta1.NodeTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NodeTask), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeTasks) UpdateStatus(ctx context.Context, nodeTask *v1beta1.NodeTask, opts v1.UpdateOptions) (*v1beta1.NodeTask, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodetasksResource, "status", nodeTask), &v1beta1.NodeTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NodeTask), err
}

// Delete takes name of the nodeTask and deletes it. Returns an error if one occurs.
func (c *FakeNodeTasks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodetasksResource, name), &v1beta1.NodeTask{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeTasks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodetasksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NodeTaskList{})
	return err
}

// Patch applies the patch and returns the patched nodeTask.
func (c *FakeNodeTasks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NodeTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodetasksResource, name, pt, data, subresources...), &v1beta1.NodeTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NodeTask), err
}
