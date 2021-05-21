module kubeform.dev/provider-linode-controller

go 1.16

require (
	github.com/fatih/structs v1.1.0
	github.com/go-logr/logr v0.4.0
	github.com/gobuffalo/flect v0.2.2
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/terraform-plugin-go v0.2.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.5.0
	github.com/imdario/mergo v0.3.12
	github.com/json-iterator/go v1.1.10
	github.com/linode/terraform-provider-linode v1.17.0
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
	kmodules.xyz/client-go v0.0.0-20210425191502-3a7296dae084
	kubeform.dev/apimachinery v0.0.0-20210507050445-8aadb2dc0a84
	kubeform.dev/provider-linode-api v0.0.0-20210508091551-1a5f34ccca8e
	sigs.k8s.io/controller-runtime v0.9.0-alpha.1.0.20210409054349-c7c85eb214f0
)

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.21.0-rc.0.0.20210405112358-ad4c2289ba4c

replace github.com/json-iterator/go => github.com/gomodules/json-iterator v1.1.12-0.20210506053207-2a3ea71074bc
