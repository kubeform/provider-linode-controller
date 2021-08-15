// Code generated by Kubeform. DO NOT EDIT.

package token

import (
	"context"

	"github.com/go-errors/errors"
	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	auditlib "go.bytebuilders.dev/audit/lib"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	meta_util "kmodules.xyz/client-go/meta"
	tokenv1alpha1 "kubeform.dev/provider-linode-api/apis/token/v1alpha1"
	"kubeform.dev/provider-linode-controller/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// TokenReconciler reconciles a Token object
type TokenReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	Gvk              schema.GroupVersionKind // GVK of the Resource
	Provider         *tfschema.Provider      // returns a *schema.Provider from the provider package
	Resource         *tfschema.Resource      // returns *schema.Resource
	TypeName         string                  // resource type
	WatchOnlyDefault bool
}

// +kubebuilder:rbac:groups=token.linode.kubeform.com,resources=tokens,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=token.linode.kubeform.com,resources=tokens/status,verbs=get;update;patch

func (r *TokenReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("token", req.NamespacedName)

	if r.WatchOnlyDefault && req.Namespace != v1.NamespaceDefault {
		log.Info("Only default namespace is supported for Kubeform Community, Please upgrade to Kubeform Enterprise to use any namespace.")
		return ctrl.Result{}, nil
	}
	var unstructuredObj unstructured.Unstructured
	unstructuredObj.SetGroupVersionKind(r.Gvk)

	if err := r.Get(ctx, req.NamespacedName, &unstructuredObj); err != nil {
		log.Error(err, "unable to fetch Token")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	rClient := r.Client
	provider := r.Provider
	res := r.Resource
	gv := r.Gvk.GroupVersion()
	tName := r.TypeName
	jsonit := controllers.GetJSONItr(tokenv1alpha1.GetEncoder(), tokenv1alpha1.GetDecoder())
	err := controllers.StartProcess(rClient, provider, ctx, res, gv, &unstructuredObj, tName, jsonit)
	return ctrl.Result{}, err
}

func (r *TokenReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, auditor *auditlib.EventPublisher) error {
	if auditor != nil {
		if err := auditor.SetupWithManager(ctx, mgr, &tokenv1alpha1.Token{}); err != nil {
			klog.Error(err, "unable to set up auditor", tokenv1alpha1.Token{}.APIVersion, tokenv1alpha1.Token{}.Kind)
			return err
		}
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&tokenv1alpha1.Token{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return !meta_util.MustAlreadyReconciled(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return (e.ObjectNew.(metav1.Object)).GetDeletionTimestamp() != nil || !meta_util.MustAlreadyReconciled(e.ObjectNew)
			},
		}).
		Watches(
			&source.Kind{Type: &v1.Secret{}},
			handler.EnqueueRequestsFromMapFunc(r.SensitiveSecretWatch(ctx)),
		).
		Complete(r)
}

func (r *TokenReconciler) SensitiveSecretWatch(ctx context.Context) handler.MapFunc {
	log := ctrl.LoggerFrom(ctx)
	return func(o client.Object) []ctrl.Request {
		result := []ctrl.Request{}

		sensSec, ok := o.(*v1.Secret)
		if !ok {
			log.Error(errors.Errorf("expected a Secret but go a %T", o), "failed to get secret for Token")
			return nil
		}

		secName := sensSec.Name
		secNamespace := sensSec.Namespace

		resourceList := &tokenv1alpha1.TokenList{}
		if err := r.List(ctx, resourceList, client.InNamespace(secNamespace)); err != nil {
			log.Error(err, "failed to list Token")
			return nil
		}

		for _, res := range resourceList.Items {
			if res.Spec.SecretRef.Name == secName {
				name := client.ObjectKey{Namespace: res.Namespace, Name: res.Name}
				result = append(result, ctrl.Request{NamespacedName: name})
			}
		}
		return result
	}
}
