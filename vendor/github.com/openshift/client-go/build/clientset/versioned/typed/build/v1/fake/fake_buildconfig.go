// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/build/v1"
	buildv1 "github.com/openshift/client-go/build/applyconfigurations/build/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuildConfigs implements BuildConfigInterface
type FakeBuildConfigs struct {
	Fake *FakeBuildV1
	ns   string
}

var buildconfigsResource = v1.SchemeGroupVersion.WithResource("buildconfigs")

var buildconfigsKind = v1.SchemeGroupVersion.WithKind("BuildConfig")

// Get takes name of the buildConfig, and returns the corresponding buildConfig object, and an error if there is any.
func (c *FakeBuildConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildconfigsResource, c.ns, name), &v1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BuildConfig), err
}

// List takes label and field selectors, and returns the list of BuildConfigs that match those selectors.
func (c *FakeBuildConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.BuildConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildconfigsResource, buildconfigsKind, c.ns, opts), &v1.BuildConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.BuildConfigList{ListMeta: obj.(*v1.BuildConfigList).ListMeta}
	for _, item := range obj.(*v1.BuildConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buildConfigs.
func (c *FakeBuildConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildconfigsResource, c.ns, opts))

}

// Create takes the representation of a buildConfig and creates it.  Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *FakeBuildConfigs) Create(ctx context.Context, buildConfig *v1.BuildConfig, opts metav1.CreateOptions) (result *v1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildconfigsResource, c.ns, buildConfig), &v1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BuildConfig), err
}

// Update takes the representation of a buildConfig and updates it. Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *FakeBuildConfigs) Update(ctx context.Context, buildConfig *v1.BuildConfig, opts metav1.UpdateOptions) (result *v1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildconfigsResource, c.ns, buildConfig), &v1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BuildConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuildConfigs) UpdateStatus(ctx context.Context, buildConfig *v1.BuildConfig, opts metav1.UpdateOptions) (*v1.BuildConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildconfigsResource, "status", c.ns, buildConfig), &v1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BuildConfig), err
}

// Delete takes name of the buildConfig and deletes it. Returns an error if one occurs.
func (c *FakeBuildConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(buildconfigsResource, c.ns, name, opts), &v1.BuildConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuildConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.BuildConfigList{})
	return err
}

// Patch applies the patch and returns the patched buildConfig.
func (c *FakeBuildConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildconfigsResource, c.ns, name, pt, data, subresources...), &v1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BuildConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied buildConfig.
func (c *FakeBuildConfigs) Apply(ctx context.Context, buildConfig *buildv1.BuildConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.BuildConfig, err error) {
	if buildConfig == nil {
		return nil, fmt.Errorf("buildConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(buildConfig)
	if err != nil {
		return nil, err
	}
	name := buildConfig.Name
	if name == nil {
		return nil, fmt.Errorf("buildConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildconfigsResource, c.ns, *name, types.ApplyPatchType, data), &v1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BuildConfig), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeBuildConfigs) ApplyStatus(ctx context.Context, buildConfig *buildv1.BuildConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.BuildConfig, err error) {
	if buildConfig == nil {
		return nil, fmt.Errorf("buildConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(buildConfig)
	if err != nil {
		return nil, err
	}
	name := buildConfig.Name
	if name == nil {
		return nil, fmt.Errorf("buildConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildconfigsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.BuildConfig), err
}

// Instantiate takes the representation of a buildRequest and creates it.  Returns the server's representation of the build, and an error, if there is any.
func (c *FakeBuildConfigs) Instantiate(ctx context.Context, buildConfigName string, buildRequest *v1.BuildRequest, opts metav1.CreateOptions) (result *v1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceAction(buildconfigsResource, buildConfigName, "instantiate", c.ns, buildRequest), &v1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Build), err
}