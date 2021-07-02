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

type Nodebalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerSpec   `json:"spec,omitempty"`
	Status            NodebalancerStatus `json:"status,omitempty"`
}

type NodebalancerSpecTransfer struct {
	// The total transfer, in MB, used by this NodeBalancer this month
	// +optional
	In *float64 `json:"in,omitempty" tf:"in"`
	// The total inbound transfer, in MB, used for this NodeBalancer this month
	// +optional
	Out *float64 `json:"out,omitempty" tf:"out"`
	// The total outbound transfer, in MB, used for this NodeBalancer this month
	// +optional
	Total *float64 `json:"total,omitempty" tf:"total"`
}

type NodebalancerSpec struct {
	KubeformOutput *NodebalancerSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource NodebalancerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NodebalancerSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Throttle connections per second (0-20). Set to 0 (zero) to disable throttling.
	// +optional
	ClientConnThrottle *int64 `json:"clientConnThrottle,omitempty" tf:"client_conn_throttle"`
	// When this NodeBalancer was created.
	// +optional
	Created *string `json:"created,omitempty" tf:"created"`
	// This NodeBalancer's hostname, ending with .nodebalancer.linode.com
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// The Public IPv4 Address of this NodeBalancer
	// +optional
	Ipv4 *string `json:"ipv4,omitempty" tf:"ipv4"`
	// The Public IPv6 Address of this NodeBalancer
	// +optional
	Ipv6 *string `json:"ipv6,omitempty" tf:"ipv6"`
	// The label of the Linode NodeBalancer.
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// The region where this NodeBalancer will be deployed.
	Region *string `json:"region" tf:"region"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Information about the amount of transfer this NodeBalancer has had so far this month.
	// +optional
	Transfer []NodebalancerSpecTransfer `json:"transfer,omitempty" tf:"transfer"`
	// When this NodeBalancer was last updated.
	// +optional
	Updated *string `json:"updated,omitempty" tf:"updated"`
}

type NodebalancerStatus struct {
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

// NodebalancerList is a list of Nodebalancers
type NodebalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Nodebalancer CRD objects
	Items []Nodebalancer `json:"items,omitempty"`
}
