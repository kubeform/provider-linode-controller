// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"fmt"

	base "kubeform.dev/apimachinery/api/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *Stackscript) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-stackscript-linode-kubeform-com-v1alpha1-stackscript,mutating=false,failurePolicy=fail,groups=stackscript.linode.kubeform.com,resources=stackscripts,versions=v1alpha1,name=stackscript.stackscript.linode.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &Stackscript{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Stackscript) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Stackscript) ValidateUpdate(old runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Stackscript) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`stackscript "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
