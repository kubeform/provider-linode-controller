// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	"github.com/linode/terraform-provider-linode/linode"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
	controllersvlan "kubeform.dev/provider-linode-controller/controllers/vlan"
	controllersvolume "kubeform.dev/provider-linode-controller/controllers/volume"

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
	vlanv1alpha1 "kubeform.dev/provider-linode-api/apis/vlan/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-linode-api/apis/volume/v1alpha1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func SetupManager(mgr manager.Manager) error {
	if err := (&controllersdomain.DomainReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Domain"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "domain.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Domain",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_domain"],
		TypeName: "linode_domain",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Domain")
		return err
	}
	if err := (&controllersdomain.RecordReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Record"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "domain.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Record",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_domain_record"],
		TypeName: "linode_domain_record",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Record")
		return err
	}
	if err := (&controllersfirewall.FirewallReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Firewall"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "firewall.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Firewall",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_firewall"],
		TypeName: "linode_firewall",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Firewall")
		return err
	}
	if err := (&controllersimage.ImageReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Image"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "image.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Image",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_image"],
		TypeName: "linode_image",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Image")
		return err
	}
	if err := (&controllersinstance.InstanceReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Instance"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "instance.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Instance",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_instance"],
		TypeName: "linode_instance",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Instance")
		return err
	}
	if err := (&controllersinstance.IpReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Ip"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "instance.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Ip",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_instance_ip"],
		TypeName: "linode_instance_ip",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Ip")
		return err
	}
	if err := (&controllerslke.ClusterReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Cluster"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "lke.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Cluster",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_lke_cluster"],
		TypeName: "linode_lke_cluster",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Cluster")
		return err
	}
	if err := (&controllersnodebalancer.NodebalancerReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Nodebalancer"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "nodebalancer.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Nodebalancer",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_nodebalancer"],
		TypeName: "linode_nodebalancer",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Nodebalancer")
		return err
	}
	if err := (&controllersnodebalancer.ConfigReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Config"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "nodebalancer.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Config",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_nodebalancer_config"],
		TypeName: "linode_nodebalancer_config",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Config")
		return err
	}
	if err := (&controllersnodebalancer.NodeReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Node"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "nodebalancer.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Node",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_nodebalancer_node"],
		TypeName: "linode_nodebalancer_node",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Node")
		return err
	}
	if err := (&controllersobject.StorageBucketReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("StorageBucket"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "object.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "StorageBucket",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_object_storage_bucket"],
		TypeName: "linode_object_storage_bucket",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "StorageBucket")
		return err
	}
	if err := (&controllersobject.StorageKeyReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("StorageKey"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "object.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "StorageKey",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_object_storage_key"],
		TypeName: "linode_object_storage_key",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "StorageKey")
		return err
	}
	if err := (&controllersobject.StorageObjectReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("StorageObject"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "object.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "StorageObject",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_object_storage_object"],
		TypeName: "linode_object_storage_object",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "StorageObject")
		return err
	}
	if err := (&controllersrdns.RdnsReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Rdns"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "rdns.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Rdns",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_rdns"],
		TypeName: "linode_rdns",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Rdns")
		return err
	}
	if err := (&controllerssshkey.SshkeyReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Sshkey"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "sshkey.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Sshkey",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_sshkey"],
		TypeName: "linode_sshkey",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Sshkey")
		return err
	}
	if err := (&controllersstackscript.StackscriptReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Stackscript"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "stackscript.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Stackscript",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_stackscript"],
		TypeName: "linode_stackscript",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Stackscript")
		return err
	}
	if err := (&controllerstoken.TokenReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Token"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "token.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Token",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_token"],
		TypeName: "linode_token",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Token")
		return err
	}
	if err := (&controllersuser.UserReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("User"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "user.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "User",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_user"],
		TypeName: "linode_user",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "User")
		return err
	}
	if err := (&controllersvlan.VlanReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Vlan"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "vlan.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Vlan",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_vlan"],
		TypeName: "linode_vlan",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Vlan")
		return err
	}
	if err := (&controllersvolume.VolumeReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Volume"),
		Scheme: mgr.GetScheme(),
		Gvk: schema.GroupVersionKind{
			Group:   "volume.linode.kubeform.com",
			Version: "v1alpha1",
			Kind:    "Volume",
		},
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_volume"],
		TypeName: "linode_volume",
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Volume")
		return err
	}

	return nil
}

func SetupWebhook(mgr manager.Manager) error {
	if err := (&domainv1alpha1.Domain{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Domain")
		return err
	}
	if err := (&domainv1alpha1.Record{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Record")
		return err
	}
	if err := (&firewallv1alpha1.Firewall{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Firewall")
		return err
	}
	if err := (&imagev1alpha1.Image{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Image")
		return err
	}
	if err := (&instancev1alpha1.Instance{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Instance")
		return err
	}
	if err := (&instancev1alpha1.Ip{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Ip")
		return err
	}
	if err := (&lkev1alpha1.Cluster{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Cluster")
		return err
	}
	if err := (&nodebalancerv1alpha1.Nodebalancer{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Nodebalancer")
		return err
	}
	if err := (&nodebalancerv1alpha1.Config{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Config")
		return err
	}
	if err := (&nodebalancerv1alpha1.Node{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Node")
		return err
	}
	if err := (&objectv1alpha1.StorageBucket{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "StorageBucket")
		return err
	}
	if err := (&objectv1alpha1.StorageKey{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "StorageKey")
		return err
	}
	if err := (&objectv1alpha1.StorageObject{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "StorageObject")
		return err
	}
	if err := (&rdnsv1alpha1.Rdns{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Rdns")
		return err
	}
	if err := (&sshkeyv1alpha1.Sshkey{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Sshkey")
		return err
	}
	if err := (&stackscriptv1alpha1.Stackscript{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Stackscript")
		return err
	}
	if err := (&tokenv1alpha1.Token{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Token")
		return err
	}
	if err := (&userv1alpha1.User{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "User")
		return err
	}
	if err := (&vlanv1alpha1.Vlan{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Vlan")
		return err
	}
	if err := (&volumev1alpha1.Volume{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "Volume")
		return err
	}

	return nil
}
