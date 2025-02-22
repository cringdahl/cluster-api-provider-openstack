/*
Copyright 2024 The Kubernetes Authors.

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

package v1alpha7

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1alpha7 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha7"
	apiv1alpha7 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/applyconfiguration/api/v1alpha7"
	scheme "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/clientset/clientset/scheme"
)

// OpenStackClustersGetter has a method to return a OpenStackClusterInterface.
// A group's client should implement this interface.
type OpenStackClustersGetter interface {
	OpenStackClusters(namespace string) OpenStackClusterInterface
}

// OpenStackClusterInterface has methods to work with OpenStackCluster resources.
type OpenStackClusterInterface interface {
	Create(ctx context.Context, openStackCluster *v1alpha7.OpenStackCluster, opts v1.CreateOptions) (*v1alpha7.OpenStackCluster, error)
	Update(ctx context.Context, openStackCluster *v1alpha7.OpenStackCluster, opts v1.UpdateOptions) (*v1alpha7.OpenStackCluster, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, openStackCluster *v1alpha7.OpenStackCluster, opts v1.UpdateOptions) (*v1alpha7.OpenStackCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha7.OpenStackCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha7.OpenStackClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha7.OpenStackCluster, err error)
	Apply(ctx context.Context, openStackCluster *apiv1alpha7.OpenStackClusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha7.OpenStackCluster, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, openStackCluster *apiv1alpha7.OpenStackClusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha7.OpenStackCluster, err error)
	OpenStackClusterExpansion
}

// openStackClusters implements OpenStackClusterInterface
type openStackClusters struct {
	*gentype.ClientWithListAndApply[*v1alpha7.OpenStackCluster, *v1alpha7.OpenStackClusterList, *apiv1alpha7.OpenStackClusterApplyConfiguration]
}

// newOpenStackClusters returns a OpenStackClusters
func newOpenStackClusters(c *InfrastructureV1alpha7Client, namespace string) *openStackClusters {
	return &openStackClusters{
		gentype.NewClientWithListAndApply[*v1alpha7.OpenStackCluster, *v1alpha7.OpenStackClusterList, *apiv1alpha7.OpenStackClusterApplyConfiguration](
			"openstackclusters",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha7.OpenStackCluster { return &v1alpha7.OpenStackCluster{} },
			func() *v1alpha7.OpenStackClusterList { return &v1alpha7.OpenStackClusterList{} }),
	}
}
