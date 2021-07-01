/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/linode/terraform-provider-linode/linode"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	clientSetScheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/scale/scheme"
	instanceScheme "kubeform.dev/provider-linode-api/apis/instance/v1alpha1"
	linodeclient "kubeform.dev/provider-linode-api/client/clientset/versioned"
	controllersinstance "kubeform.dev/provider-linode-controller/controllers/instance"
	"kubeform.dev/provider-linode-controller/tests/e2e/framework"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var (
	root *framework.Framework
)

var cfg *rest.Config
var k8sClient client.Client
var k8sManager ctrl.Manager
var testEnv *envtest.Environment

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)

	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "e2e Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {
	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crds")},
		ErrorIfCRDPathMissing: true,
	}
	testEnv.ControlPlaneStopTimeout = 2 * time.Minute
	cfg, err := testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = scheme.AddToScheme(clientSetScheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = instanceScheme.AddToScheme(clientSetScheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	k8sClient, err = client.New(cfg, client.Options{Scheme: clientSetScheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: clientSetScheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	gvk := schema.GroupVersionKind{
		Group:   "instance.linode.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Instance",
	}

	err = (&controllersinstance.InstanceReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("Instance"),
		Scheme:   k8sManager.GetScheme(),
		Gvk:      gvk,
		Provider: linode.Provider(),
		Resource: linode.Provider().ResourcesMap["linode_instance"],
		TypeName: "linode_instance",
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	// Clients
	kubeClient := kubernetes.NewForConfigOrDie(cfg)
	linodeClient := linodeclient.NewForConfigOrDie(cfg)

	// Framework
	root = framework.New(cfg, kubeClient, linodeClient)

	// Create namespace
	By("Using namespace " + root.Namespace())
	err = root.CreateNamespace()
	Expect(err).NotTo(HaveOccurred())
	root.EventuallyCRD().Should(Succeed())

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		Expect(err).ToNot(HaveOccurred())
	}()
}, 60)

var _ = AfterSuite(func() {
	By("Deleting Namespace")
	err := root.DeleteNamespace()
	Expect(err).NotTo(HaveOccurred())

	By("tearing down the test environment")
	err = testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})
