GO_ARCH:=$(shell go env GOARCH)
GO_OS:=$(shell go env GOOS)

BINARY_NAME:=conformanceTestExecutor

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go build -ldflags "-s -w -extldflags -static" -o bin/darwin/amd64/$(BINARY_NAME) github.com/aws/eks-distro-build-tooling/golang/conformance-test-executor

.PHONY: clean
clean:
	rm -rf ./bin
