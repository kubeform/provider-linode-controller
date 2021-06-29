package e2e_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	linodeinstance "kubeform.dev/provider-linode-api/apis/instance/v1alpha1"
	"kubeform.dev/provider-linode-controller/tests/e2e/framework"
)

var _ = Describe("Test", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
	})

	Describe("Linode", func() {
		Context("InstanceController", func() {
			var (
				providerRef   *core.Secret
				sensitiveData *core.Secret
				secretName    string
				instanceName  string
				instance      *linodeinstance.Instance
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				instanceName = f.GetRandomName("")
				providerRef = f.LinodeProviderRef(secretName)
				sensitiveData = f.InstanceSensitiveData()
				instance = f.Instance(instanceName, secretName)
			})

			It("should create and delete instance successfully", func() {
				By("Creating LinodeProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Secret")
				err = f.CreateSecret(sensitiveData)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Instance")
				err = f.CreateInstance(instance)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Instance")
				f.EventuallyInstanceRunning(instance.ObjectMeta).Should(BeTrue())

				By("Deleting Instance")
				err = f.DeleteInstance(instance.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting instance")
				f.EventuallyInstanceDeleted(instance.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})
	})
})
