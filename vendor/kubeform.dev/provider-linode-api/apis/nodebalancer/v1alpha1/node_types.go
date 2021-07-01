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

type Node struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeSpec   `json:"spec,omitempty"`
	Status            NodeStatus `json:"status,omitempty"`
}

type NodeSpec struct {
	KubeformOutput *NodeSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource NodeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NodeSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The private IP Address and port (IP:PORT) where this backend can be reached. This must be a private IP address.
	Address *string `json:"address" tf:"address"`
	// The ID of the NodeBalancerConfig to access.
	ConfigID *int64 `json:"configID" tf:"config_id"`
	// The label for this node. This is for display purposes only.
	Label *string `json:"label" tf:"label"`
	// The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. If set to `backup` this backend will only accept traffic if all other nodes are down.
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// The ID of the NodeBalancer to access.
	NodebalancerID *int64 `json:"nodebalancerID" tf:"nodebalancer_id"`
	// The current status of this node, based on the configured checks of its NodeBalancer Config. (unknown, UP, DOWN)
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255)
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type NodeStatus struct {
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

// NodeList is a list of Nodes
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Node CRD objects
	Items []Node `json:"items,omitempty"`
}
