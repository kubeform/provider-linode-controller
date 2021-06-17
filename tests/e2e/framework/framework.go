package framework

import (
	linodeclient "kubeform.dev/provider-linode-api/client/clientset/versioned"

	"github.com/appscode/go/crypto/rand"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Framework struct {
	restConfig   *rest.Config
	kubeClient   kubernetes.Interface
	linodeClient linodeclient.Interface
	namespace    string
	name         string
}

func New(
	restConfig *rest.Config,
	kubeClient kubernetes.Interface,
	linodeClient linodeclient.Interface,
) *Framework {
	return &Framework{
		restConfig:   restConfig,
		kubeClient:   kubeClient,
		linodeClient: linodeClient,
		name:         "kfc",
		namespace:    rand.WithUniqSuffix("kubeform"),
	}
}

func (f *Framework) Invoke() *Invocation {
	return &Invocation{
		Framework: f,
		app:       rand.WithUniqSuffix("kfc-e2e"),
	}
}

func (fi *Invocation) GetRandomName(extraSuffix string) string {
	return rand.WithUniqSuffix(fi.name + extraSuffix)
}

type Invocation struct {
	*Framework
	app string
}
