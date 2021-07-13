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
	"fmt"

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
		if !framework.RunTest(framework.INSTANCE, whichController) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichController flag is either `all` or `%s` but got `%s`", framework.INSTANCE, framework.INSTANCE, whichController))
		}
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

			AfterEach(func() {
				By("Deleting Instance")
				err = f.DeleteInstance(instance.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting instance")
				f.EventuallyInstanceDeleted(instance.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
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
			})

			It("should create, update and delete instance successfully", func() {
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

				By("Updating Instance")
				err = f.UpdateInstance(instance)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Instance")
				f.EventuallyInstanceRunning(instance.ObjectMeta).Should(BeTrue())
			})

			It("should create, update (sensitive field) and delete instance successfully", func() {
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

				By("Updating Instance sensitive field")
				err, secret := f.UpdateInstanceSensitive(sensitiveData)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for updating sensitive secret")
				f.EventuallySensitiveSecretUpdating(sensitiveData, secret.Data["output"]).Should(BeTrue())

				By("Wait for Running Instance")
				f.EventuallyInstanceRunning(instance.ObjectMeta).Should(BeTrue())
			})
		})
	})
})
