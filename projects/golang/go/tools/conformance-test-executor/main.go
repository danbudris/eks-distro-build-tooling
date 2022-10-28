package main

import (
	"os"

	"github.com/aws/eks-distro-build-tooling/golang/conformance-test-executor/cmd"
)

func main() {
	if cmd.Execute() == nil {
		os.Exit(0)
	}
	os.Exit(-1)
}