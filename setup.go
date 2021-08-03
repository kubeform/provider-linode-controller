/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gobuffalo/flect"
	linode "github.com/linode/terraform-provider-linode/linode"
	auditlib "go.bytebuilders.dev/audit/lib"
	arv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	informers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	domainv1alpha1 "kubeform.dev/provider-linode-api/apis/domain/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-linode-api/apis/firewall/v1alpha1"
	imagev1alpha1 "kubeform.dev/provider-linode-api/apis/image/v1alpha1"
	instancev1alpha1 "kubeform.dev/provider-linode-api/apis/instance/v1alpha1"
	lkev1alpha1 "kubeform.dev/provider-linode-api/apis/lke/v1alpha1"
	nodebalancerv1alpha1 "kubeform.dev/provider-linode-api/apis/nodebalancer/v1alpha1"
	objectv1alpha1 "kubeform.dev/provider-linode-api/apis/object/v1alpha1"
	rdnsv1alpha1 "kubeform.dev/provider-linode-api/apis/rdns/v1alpha1"
	sshkeyv1alpha1 "kubeform.dev/provider-linode-api/apis/sshkey/v1alpha1"
	stackscriptv1alpha1 "kubeform.dev/provider-linode-api/apis/stackscript/v1alpha1"
	tokenv1alpha1 "kubeform.dev/provider-linode-api/apis/token/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-linode-api/apis/user/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-linode-api/apis/volume/v1alpha1"
	controllersdomain "kubeform.dev/provider-linode-controller/controllers/domain"
	controllersfirewall "kubeform.dev/provider-linode-controller/controllers/firewall"
	controllersimage "kubeform.dev/provider-linode-controller/controllers/image"
	controllersinstance "kubeform.dev/provider-linode-controller/controllers/instance"
	controllerslke "kubeform.dev/provider-linode-controller/controllers/lke"
	controllersnodebalancer "kubeform.dev/provider-linode-controller/controllers/nodebalancer"
	controllersobject "kubeform.dev/provider-linode-controller/controllers/object"
	controllersrdns "kubeform.dev/provider-linode-controller/controllers/rdns"
	controllerssshkey "kubeform.dev/provider-linode-controller/controllers/sshkey"
	controllersstackscript "kubeform.dev/provider-linode-controller/controllers/stackscript"
	controllerstoken "kubeform.dev/provider-linode-controller/controllers/token"
	controllersuser "kubeform.dev/provider-linode-controller/controllers/user"
	controllersvolume "kubeform.dev/provider-linode-controller/controllers/volume"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var _provider = linode.Provider()

var runningControllers = struct {
	sync.RWMutex
	mp map[schema.GroupVersionKind]bool
}{mp: make(map[schema.GroupVersionKind]bool)}

func watchCRD(ctx context.Context, crdClient *clientset.Clientset, vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, stopCh <-chan struct{}, mgr manager.Manager, auditor *auditlib.EventPublisher, watchOnlyDefault bool) error {
	informerFactory := informers.NewSharedInformerFactory(crdClient, time.Second*30)
	i := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Informer()
	l := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Lister()

	i.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			var key string
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				klog.Error(err)
				return
			}

			_, name, err := cache.SplitMetaNamespaceKey(key)
			if err != nil {
				klog.Error(err)
				return
			}

			crd, err := l.Get(name)
			if err != nil {
				klog.Error(err)
				return
			}
			if strings.Contains(crd.Spec.Group, "linode.kubeform.com") {
				gvk := schema.GroupVersionKind{
					Group:   crd.Spec.Group,
					Version: crd.Spec.Versions[0].Name,
					Kind:    crd.Spec.Names.Kind,
				}

				// check whether this gvk came before, if no then start the controller
				runningControllers.RLock()
				_, ok := runningControllers.mp[gvk]
				runningControllers.RUnlock()

				if !ok {
					runningControllers.Lock()
					runningControllers.mp[gvk] = true
					runningControllers.Unlock()

					if enableValidatingWebhook {
						// add dynamic ValidatingWebhookConfiguration

						// create empty VWC if the group has come for the first time
						err := createEmptyVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						// update
						err = updateVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						err = SetupWebhook(mgr, gvk)
						if err != nil {
							setupLog.Error(err, "unable to enable webhook")
							os.Exit(1)
						}
					}

					err = SetupManager(ctx, mgr, gvk, auditor, watchOnlyDefault)
					if err != nil {
						setupLog.Error(err, "unable to start manager")
						os.Exit(1)
					}
				}
			}
		},
	})

	informerFactory.Start(stopCh)

	return nil
}

func createEmptyVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	_, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err == nil || !(errors.IsNotFound(err)) {
		return err
	}

	emptyVWC := &arv1.ValidatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ValidatingWebhookConfiguration",
			APIVersion: "admissionregistration.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-"),
			Labels: map[string]string{
				"app.kubernetes.io/instance": "linode.kubeform.com",
				"app.kubernetes.io/part-of":  "kubeform.com",
			},
		},
	}
	_, err = vwcClient.ValidatingWebhookConfigurations().Create(context.TODO(), emptyVWC, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func updateVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	vwc, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	path := "/validate-" + strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-") + "-v1alpha1-" + strings.ToLower(gvk.Kind)
	fail := arv1.Fail
	sideEffects := arv1.SideEffectClassNone
	admissionReviewVersions := []string{"v1beta1"}

	rules := []arv1.RuleWithOperations{
		{
			Operations: []arv1.OperationType{
				arv1.Delete,
			},
			Rule: arv1.Rule{
				APIGroups:   []string{strings.ToLower(gvk.Group)},
				APIVersions: []string{gvk.Version},
				Resources:   []string{strings.ToLower(flect.Pluralize(gvk.Kind))},
			},
		},
	}

	data, err := ioutil.ReadFile("/tmp/k8s-webhook-server/serving-certs/ca.crt")
	if err != nil {
		return err
	}

	name := strings.ToLower(gvk.Kind) + "." + gvk.Group
	for _, webhook := range vwc.Webhooks {
		if webhook.Name == name {
			return nil
		}
	}

	newWebhook := arv1.ValidatingWebhook{
		Name: name,
		ClientConfig: arv1.WebhookClientConfig{
			Service: &arv1.ServiceReference{
				Namespace: webhookNamespace,
				Name:      webhookName,
				Path:      &path,
			},
			CABundle: data,
		},
		Rules:                   rules,
		FailurePolicy:           &fail,
		SideEffects:             &sideEffects,
		AdmissionReviewVersions: admissionReviewVersions,
	}

	vwc.Webhooks = append(vwc.Webhooks, newWebhook)

	_, err = vwcClient.ValidatingWebhookConfigurations().Update(context.TODO(), vwc, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func SetupManager(ctx context.Context, mgr manager.Manager, gvk schema.GroupVersionKind, auditor *auditlib.EventPublisher, watchOnlyDefault bool) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "domain.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Domain",
	}:
		if err := (&controllersdomain.DomainReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Domain"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_domain"],
			TypeName:         "linode_domain",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Domain")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "domain.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Record",
	}:
		if err := (&controllersdomain.RecordReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Record"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_domain_record"],
			TypeName:         "linode_domain_record",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Record")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "firewall.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Firewall",
	}:
		if err := (&controllersfirewall.FirewallReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Firewall"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_firewall"],
			TypeName:         "linode_firewall",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Firewall")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "image.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Image",
	}:
		if err := (&controllersimage.ImageReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Image"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_image"],
			TypeName:         "linode_image",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Image")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "instance.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Instance",
	}:
		if err := (&controllersinstance.InstanceReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Instance"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_instance"],
			TypeName:         "linode_instance",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Instance")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "instance.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Ip",
	}:
		if err := (&controllersinstance.IpReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Ip"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_instance_ip"],
			TypeName:         "linode_instance_ip",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Ip")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "lke.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Cluster",
	}:
		if err := (&controllerslke.ClusterReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Cluster"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_lke_cluster"],
			TypeName:         "linode_lke_cluster",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Cluster")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "nodebalancer.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Nodebalancer",
	}:
		if err := (&controllersnodebalancer.NodebalancerReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Nodebalancer"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_nodebalancer"],
			TypeName:         "linode_nodebalancer",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Nodebalancer")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "nodebalancer.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Config",
	}:
		if err := (&controllersnodebalancer.ConfigReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Config"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_nodebalancer_config"],
			TypeName:         "linode_nodebalancer_config",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Config")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "nodebalancer.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Node",
	}:
		if err := (&controllersnodebalancer.NodeReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Node"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_nodebalancer_node"],
			TypeName:         "linode_nodebalancer_node",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Node")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "object.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "StorageBucket",
	}:
		if err := (&controllersobject.StorageBucketReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("StorageBucket"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_object_storage_bucket"],
			TypeName:         "linode_object_storage_bucket",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "StorageBucket")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "object.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "StorageKey",
	}:
		if err := (&controllersobject.StorageKeyReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("StorageKey"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_object_storage_key"],
			TypeName:         "linode_object_storage_key",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "StorageKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "object.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "StorageObject",
	}:
		if err := (&controllersobject.StorageObjectReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("StorageObject"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_object_storage_object"],
			TypeName:         "linode_object_storage_object",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "StorageObject")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "rdns.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rdns",
	}:
		if err := (&controllersrdns.RdnsReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Rdns"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_rdns"],
			TypeName:         "linode_rdns",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Rdns")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "sshkey.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Sshkey",
	}:
		if err := (&controllerssshkey.SshkeyReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Sshkey"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_sshkey"],
			TypeName:         "linode_sshkey",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Sshkey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "stackscript.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Stackscript",
	}:
		if err := (&controllersstackscript.StackscriptReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Stackscript"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_stackscript"],
			TypeName:         "linode_stackscript",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Stackscript")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "token.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Token",
	}:
		if err := (&controllerstoken.TokenReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Token"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_token"],
			TypeName:         "linode_token",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Token")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "User",
	}:
		if err := (&controllersuser.UserReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("User"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_user"],
			TypeName:         "linode_user",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "User")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Volume",
	}:
		if err := (&controllersvolume.VolumeReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Volume"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         _provider,
			Resource:         _provider.ResourcesMap["linode_volume"],
			TypeName:         "linode_volume",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Volume")
			return err
		}

	default:
		return fmt.Errorf("Invalid CRD")
	}

	return nil
}

func SetupWebhook(mgr manager.Manager, gvk schema.GroupVersionKind) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "domain.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Domain",
	}:
		if err := (&domainv1alpha1.Domain{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Domain")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "domain.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Record",
	}:
		if err := (&domainv1alpha1.Record{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Record")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "firewall.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Firewall",
	}:
		if err := (&firewallv1alpha1.Firewall{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Firewall")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "image.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Image",
	}:
		if err := (&imagev1alpha1.Image{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Image")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "instance.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Instance",
	}:
		if err := (&instancev1alpha1.Instance{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Instance")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "instance.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Ip",
	}:
		if err := (&instancev1alpha1.Ip{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Ip")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "lke.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Cluster",
	}:
		if err := (&lkev1alpha1.Cluster{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Cluster")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "nodebalancer.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Nodebalancer",
	}:
		if err := (&nodebalancerv1alpha1.Nodebalancer{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Nodebalancer")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "nodebalancer.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Config",
	}:
		if err := (&nodebalancerv1alpha1.Config{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Config")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "nodebalancer.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Node",
	}:
		if err := (&nodebalancerv1alpha1.Node{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Node")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "object.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "StorageBucket",
	}:
		if err := (&objectv1alpha1.StorageBucket{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "StorageBucket")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "object.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "StorageKey",
	}:
		if err := (&objectv1alpha1.StorageKey{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "StorageKey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "object.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "StorageObject",
	}:
		if err := (&objectv1alpha1.StorageObject{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "StorageObject")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "rdns.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rdns",
	}:
		if err := (&rdnsv1alpha1.Rdns{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Rdns")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "sshkey.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Sshkey",
	}:
		if err := (&sshkeyv1alpha1.Sshkey{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Sshkey")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "stackscript.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Stackscript",
	}:
		if err := (&stackscriptv1alpha1.Stackscript{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Stackscript")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "token.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Token",
	}:
		if err := (&tokenv1alpha1.Token{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Token")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "User",
	}:
		if err := (&userv1alpha1.User{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "User")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Volume",
	}:
		if err := (&volumev1alpha1.Volume{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Volume")
			return err
		}

	default:
		return fmt.Errorf("Invalid Webhook")
	}

	return nil
}
