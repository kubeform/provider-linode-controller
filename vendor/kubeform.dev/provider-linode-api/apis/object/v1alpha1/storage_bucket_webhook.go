// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"fmt"

	base "kubeform.dev/apimachinery/api/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *StorageBucket) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-object-linode-kubeform-com-v1alpha1-storagebucket,mutating=false,failurePolicy=fail,groups=object.linode.kubeform.com,resources=storagebuckets,versions=v1alpha1,name=storagebucket.object.linode.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &StorageBucket{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *StorageBucket) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *StorageBucket) ValidateUpdate(old runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *StorageBucket) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`storagebucket "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
