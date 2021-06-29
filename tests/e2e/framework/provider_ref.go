package framework

import (
	"os"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/constants/linode"
)

func (i *Invocation) LinodeProviderRef(name string) *core.Secret {
	return &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
		},
		StringData: map[string]string{
			"provider": `{
				"token": "` + os.Getenv(linode.LINODE_API_TOKEN) + `"
			}`,
		},
	}
}
