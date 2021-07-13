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

package framework

import (
	"bytes"
	"context"
	"time"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kubeform.dev/provider-linode-api/apis/instance/v1alpha1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

func (i *Invocation) Instance(name string, secretName string) *v1alpha1.Instance {
	image := "linode/ubuntu18.04"
	region := "us-east"

	return &v1alpha1.Instance{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.InstanceSpec{
			ProviderRef: corev1.LocalObjectReference{
				Name: secretName,
			},
			SecretRef: &corev1.LocalObjectReference{
				Name: InstanceSecretName,
			},
			Resource: v1alpha1.InstanceSpecResource{
				Image:  &image,
				Label:  &name,
				Region: &region,
			},
		},
	}
}

func (f *Framework) CreateInstance(obj *v1alpha1.Instance) error {
	_, err := f.kfClient.InstanceV1alpha1().Instances(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) UpdateInstance(obj *v1alpha1.Instance) error {
	obj, err := f.kfClient.InstanceV1alpha1().Instances(obj.Namespace).Get(context.TODO(), obj.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	updateName := *obj.Spec.Resource.Label + "-update"
	obj.Spec.Resource.Label = &updateName

	_, err = f.kfClient.InstanceV1alpha1().Instances(obj.Namespace).Update(context.TODO(), obj, metav1.UpdateOptions{})
	return err
}

func (f *Framework) UpdateInstanceSensitive(secret *corev1.Secret) (error, *corev1.Secret) {
	secret, err := f.kubeClient.CoreV1().Secrets(secret.Namespace).Get(context.TODO(), secret.Name, metav1.GetOptions{})
	if err != nil {
		return err, secret
	}

	secret.StringData = map[string]string{
		"input": `{
				"root_pass": "NewTestPassWord@123"
		}`,
	}

	secret, err = f.kubeClient.CoreV1().Secrets(secret.Namespace).Update(context.TODO(), secret, metav1.UpdateOptions{})
	return err, secret
}

func (f *Framework) EventuallySensitiveSecretUpdating(secret *corev1.Secret, data []byte) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			secret, err := f.kubeClient.CoreV1().Secrets(secret.Namespace).Get(context.TODO(), secret.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())

			eq := bytes.Compare(secret.Data["output"], data)
			return eq != 0
		},
		time.Minute*5,
		time.Second*30,
	)
}

func (f *Framework) DeleteInstance(meta metav1.ObjectMeta) error {
	return f.kfClient.InstanceV1alpha1().Instances(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInBackground())
}

func (f *Framework) EventuallyInstanceRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			instance, err := f.kfClient.InstanceV1alpha1().Instances(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return instance.Status.Phase == status.CurrentStatus
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyInstanceDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kfClient.InstanceV1alpha1().Instances(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
