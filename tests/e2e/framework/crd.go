package framework

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) EventuallyCRD() GomegaAsyncAssertion {
	return Eventually(
		func() error {
			// Check Instances CRD
			if _, err := f.linodeClient.InstanceV1alpha1().Instances(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
				return errors.New("CRD Instances is not ready")
			}

			return nil
		},
		time.Minute*2,
		time.Second*10,
	)
}
