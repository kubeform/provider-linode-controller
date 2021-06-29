GO_PKG   := kubeform.dev
BIN      := provider-linode-controller

# Generate Controllers
.PHONY: controllers
controllers:
	@echo "Generating Controllers"
	@generator --which-provider=linode --controller-gen=true --controller-path=$(GOPATH)/src/$(GO_PKG)/$(BIN)