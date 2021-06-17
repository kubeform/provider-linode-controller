package framework

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	InstanceSecretName = "instance-secret"
)

func (i *Invocation) InstanceSensitiveData() *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      InstanceSecretName,
			Namespace: i.Namespace(),
		},
		StringData: map[string]string{
			"input": `{
				"root_pass": "thisIsAPassword123!"
			}`,
		},
	}
}
