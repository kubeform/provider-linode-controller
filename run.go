package main

import (
	"fmt"
	"os"

	// +kubebuilder:scaffold:imports

	"github.com/spf13/cobra"
	auditlib "go.bytebuilders.dev/audit/lib"
	licenseapi "go.bytebuilders.dev/license-verifier/apis/licenses/v1alpha1"
	license "go.bytebuilders.dev/license-verifier/kubernetes"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"kmodules.xyz/client-go/discovery"
	"kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/tools/cli"
	linodescheme "kubeform.dev/provider-linode-api/client/clientset/versioned/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

var (
	licenseFile             string
	enableValidatingWebhook bool
	webhookName             string
	webhookNamespace        string
	metricsAddr             string
	enableLeaderElection    bool
	probeAddr               string
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = linodescheme.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func NewCmdRun(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch Linode controller",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			klog.Infoln("Starting Linode controller...")

			ctrl.SetLogger(klogr.New())

			ctx := ctrl.SetupSignalHandler()

			mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
				Scheme:                 scheme,
				MetricsBindAddress:     metricsAddr,
				Port:                   9443,
				HealthProbeBindAddress: probeAddr,
				LeaderElection:         enableLeaderElection,
				LeaderElectionID:       "5b87adeb.linode.kubeform.com", // 5b87adeb needs to be dynamically generated for each controller
			})
			if err != nil {
				setupLog.Error(err, "unable to start manager")
				os.Exit(1)
			}
			cfg := mgr.GetConfig()

			info := license.NewLicenseEnforcer(cfg, licenseFile).LoadLicense()
			fmt.Printf("%+v\n", info)
			if info.Status != licenseapi.LicenseActive {
				fmt.Printf("License status %s", info.Status)
				os.Exit(1)
			}
			if sets.NewString(info.Features...).Has("kubeform-enterprise") {
				fmt.Println("watch all namespaces")
			} else if sets.NewString(info.Features...).Has("kubeform-community") {
				fmt.Println("watch default namespace")
			} else {
				fmt.Println("not a valid license for this product")
				os.Exit(1)
			}

			// audit event publisher
			var auditor *auditlib.EventPublisher
			if licenseFile != "" && cli.EnableAnalytics {
				kc, err := kubernetes.NewForConfig(cfg)
				if err != nil {
					setupLog.Error(err, "unable to create Kubernetes client")
					os.Exit(1)
				}
				mapper := discovery.NewResourceMapper(mgr.GetRESTMapper())
				fn := auditlib.BillingEventCreator{
					Mapper: mapper,
				}
				auditor = auditlib.NewResilientEventPublisher(func() (*auditlib.NatsConfig, error) {
					return auditlib.NewNatsConfig(kc.CoreV1().Namespaces(), licenseFile)
				}, mapper, fn.CreateEvent)
			}
			if err = auditor.SetupWithManager(ctx, mgr, nil /*&modulev1alpha1.Workflow{}*/); err != nil {
				setupLog.Error(err, "unable to set up auditor", "$group", "$kind")
				os.Exit(1)
			}

			crdClient := clientset.NewForConfigOrDie(cfg)
			vwcClient := admissionregistrationv1.NewForConfigOrDie(cfg)

			err = watchCRD(crdClient, vwcClient, ctx.Done(), mgr)
			if err != nil {
				setupLog.Error(err, "unable to watch crds")
				os.Exit(1)
			}
			// +kubebuilder:scaffold:builder

			// Start periodic license verification
			//nolint:errcheck
			go license.VerifyLicensePeriodically(mgr.GetConfig(), licenseFile, ctx.Done())

			setupLog.Info("starting manager")
			if err := mgr.Start(ctx); err != nil {
				setupLog.Error(err, "problem running manager")
				os.Exit(1)
			}
		},
	}

	meta.AddLabelBlacklistFlag(cmd.Flags())
	cmd.Flags().StringVar(&licenseFile, "license-file", licenseFile, "Path to license file")
	cmd.Flags().StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	cmd.Flags().StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	cmd.Flags().BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	cmd.Flags().BoolVar(&enableValidatingWebhook, "enable-validating-webhook", false, "Enable validating webhook")
	cmd.Flags().StringVar(&webhookName, "webhook-name", "webhook-service", "Webhook name")
	cmd.Flags().StringVar(&webhookNamespace, "webhook-namespace", "kube-system", "Webhook namespace")

	return cmd
}
