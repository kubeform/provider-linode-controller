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

type StorageObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageObjectSpec   `json:"spec,omitempty"`
	Status            StorageObjectStatus `json:"status,omitempty"`
}

type StorageObjectSpec struct {
	KubeformOutput *StorageObjectSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource StorageObjectSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type StorageObjectSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The S3 access key with access to the target bucket.
	AccessKey *string `json:"accessKey" tf:"access_key"`
	// The ACL config given to this object.
	// +optional
	Acl *string `json:"acl,omitempty" tf:"acl"`
	// The target bucket to put this object in.
	Bucket *string `json:"bucket" tf:"bucket"`
	// This cache_control configuration of this object.
	// +optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control"`
	// The target cluster that the bucket is in.
	Cluster *string `json:"cluster" tf:"cluster"`
	// The contents of the Object to upload.
	// +optional
	Content *string `json:"content,omitempty" tf:"content"`
	// The base64 contents of the Object to upload.
	// +optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64"`
	// The content disposition configuration of this object.
	// +optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition"`
	// The encoding of the content of this object.
	// +optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding"`
	// The language metadata of this object.
	// +optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language"`
	// The MIME type of the content.
	// +optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type"`
	// The specific version of this object.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// Whether the object should bypass deletion restrictions.
	// +optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`
	// The name of the uploaded object.
	Key *string `json:"key" tf:"key"`
	// The metadata of this object
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
	// The S3 secret key with access to the target bucket.
	SecretKey *string `json:"secretKey" tf:"secret_key"`
	// The source file to upload.
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// The version ID of this object.
	// +optional
	VersionID *string `json:"versionID,omitempty" tf:"version_id"`
	// The website redirect location of this object.
	// +optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect"`
}

type StorageObjectStatus struct {
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

// StorageObjectList is a list of StorageObjects
type StorageObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageObject CRD objects
	Items []StorageObject `json:"items,omitempty"`
}
