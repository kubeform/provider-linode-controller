/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecPoolNodes struct {
	// The ID of the node.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The ID of the underlying Linode instance.
	// +optional
	InstanceID *int64 `json:"instanceID,omitempty" tf:"instance_id"`
	// The status of the node.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type ClusterSpecPool struct {
	// The number of nodes in the Node Pool.
	Count *int64 `json:"count" tf:"count"`
	// The ID of the Node Pool.
	// +optional
	ID *int64 `json:"ID,omitempty" tf:"id"`
	// The nodes in the node pool.
	// +optional
	Nodes []ClusterSpecPoolNodes `json:"nodes,omitempty" tf:"nodes"`
	// A Linode Type for all of the nodes in the Node Pool.
	Type *string `json:"type" tf:"type"`
}

type ClusterSpec struct {
	KubeformOutput *ClusterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The API endpoints for the cluster.
	// +optional
	ApiEndpoints []string `json:"apiEndpoints,omitempty" tf:"api_endpoints"`
	// The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.
	K8sVersion *string `json:"k8sVersion" tf:"k8s_version"`
	// The Base64-encoded Kubeconfig for the cluster.
	// +optional
	Kubeconfig *string `json:"-" sensitive:"true" tf:"kubeconfig"`
	// The unique label for the cluster.
	Label *string `json:"label" tf:"label"`
	// A node pool in the cluster.
	// +kubebuilder:validation:MinItems=1
	Pool []ClusterSpecPool `json:"pool" tf:"pool"`
	// This cluster's location.
	Region *string `json:"region" tf:"region"`
	// The status of the cluster.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
}

type ClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
