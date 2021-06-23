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

type Vlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VlanSpec   `json:"spec,omitempty"`
	Status            VlanStatus `json:"status,omitempty"`
}

type VlanSpecAttachedLinodes struct {
	// +optional
	ID *int64 `json:"ID,omitempty" tf:"id"`
	// +optional
	Ipv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address"`
	// +optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address"`
}

type VlanSpec struct {
	KubeformOutput *VlanSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource VlanSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VlanSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The Linodes attached to this vlan.
	// +optional
	AttachedLinodes []VlanSpecAttachedLinodes `json:"attachedLinodes,omitempty" tf:"attached_linodes"`
	// +optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`
	// Description of the vlan for display purposes only.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The IDs of the Linodes to attach to this vlan.
	// +optional
	Linodes []int64 `json:"linodes,omitempty" tf:"linodes"`
	// The region where the vlan is deployed.
	Region *string `json:"region" tf:"region"`
}

type VlanStatus struct {
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

// VlanList is a list of Vlans
type VlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Vlan CRD objects
	Items []Vlan `json:"items,omitempty"`
}
