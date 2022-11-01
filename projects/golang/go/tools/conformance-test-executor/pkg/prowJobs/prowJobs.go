package prowJobs

import (
	"time"

	"github.com/google/uuid"

	"github.com/aws/eks-distro-build-tooling/golang/conformance-test-executor/pkg/constants"
)

const (
	defaultCpuRequest = "2"
	defaultGitOrg = "aws"
	defaultMemoryRequest = "8Gi"
	defaultTimeout = "6h0m0s"
)


type ProwJobCommonOptions struct {
	Architecture                string
	CpuRequest                  string
	GitOrg                      string
	JobIdentifier               string
	MemoryRequest               string
	RuntimeImage                string
	Timeout                     string
}

func (g *ProwJobCommonOptions) setCommonDefaults() {
	if g.Architecture == "" {
		g.Architecture = constants.AMD64Arch
	}

	if g.CpuRequest == "" {
		g.CpuRequest = defaultCpuRequest
	}

	if g.GitOrg == "" {
		g.GitOrg = defaultGitOrg
	}

	if g.JobIdentifier == "" {
		g.JobIdentifier = uuid.Must(uuid.NewRandom()).String()
	}

	if g.MemoryRequest == "" {
		g.MemoryRequest = defaultMemoryRequest
	}

	if g.Timeout == "" {
		g.Timeout = defaultTimeout
	}

	if g.RuntimeImage == "" {
		g.RuntimeImage = "public.ecr.aws/eks-distro-build-tooling/builder-base:8a2a9d01b95ee8f2bbacbb03d1d55b6c56615411.2"
	}
}

func (g *ProwJobCommonOptions) prowJobCommonTemplateValues() map[string]interface{}{
	templateValues := make(map[string]interface{})
	templateValues["architecture"] = g.Architecture
	templateValues["cpuRequest"] = g.CpuRequest
	templateValues["gitOrg"] = g.GitOrg
	templateValues["jobId"] = g.JobIdentifier
	templateValues["memoryRequest"] = g.MemoryRequest
	templateValues["runtimeImage"] = g.RuntimeImage
	templateValues["timeout"] = g.Timeout
	return templateValues
}

func ProwJobStartTime(startTime time.Time) string {
	return startTime.Format("2006-01-02T15:04:05Z")
}