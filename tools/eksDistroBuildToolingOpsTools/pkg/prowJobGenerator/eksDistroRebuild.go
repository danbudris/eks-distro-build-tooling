package prowJobGenerator

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/aws/eks-distro-build-tooling/tools/eksDistroBuildToolingOpsTools/pkg/constants"
)

const (
	eksDRebuildProwJobTemplate = "pkg/prowJobs/templates/eks-distro-rebuild.yaml"
	eksDRebuildDefaultGitRepo = "eks-distro"
)

// NewEksDistroRebuildProwJob generates a yaml defining a triggered post-submit for the EKS dsitro main post-submit,
// using the provided options to populate the template. This job yaml can then be applied to a prow cluster to run the custom job.
func NewEksDistroRebuildProwJob(kubernetesVersion string, jobName string, baseSha string, headSha string, opts *EksDistroRebuildProwJobOptions) ([]byte, error) {
	if opts == nil {
		opts = &EksDistroRebuildProwJobOptions{}
	}
	opts.setEksDRebuildOptionsDefaults()

	temp, err := template.ParseFiles(eksDRebuildProwJobTemplate)
	if err != nil {
		return nil, fmt.Errorf("parsing template file: %v", err)
	}

	templateData := EksDistroRebuildTemplateValues(*opts)
	templateData["startTime"] = ProwJobStartTime(time.Now().UTC())
	templateData["kubernetesVersion"] = kubernetesVersion
	templateData["jobName"] = jobName
	templateData["baseSha"] = baseSha
	templateData["headSha"] = headSha
	templateData["gitRepo"] = opts.GitRepo

	var renderedTemplateData bytes.Buffer

	err = temp.ExecuteTemplate(&renderedTemplateData, "eks-distro-rebuild.yaml", templateData)
	if err != nil {
		return nil, fmt.Errorf("rendering builderBaseProwJob: %v", err)
	}

	return renderedTemplateData.Bytes(), nil
}

func EksDistroRebuildTemplateValues(opts EksDistroRebuildProwJobOptions) map[string]interface{} {
	templateValues := make(map[string]interface{})
	templateValues["testRoleArn"] = opts.TestRoleArn
	templateValues["artifactsBucket"] = opts.ArtifactsBucket
	templateValues["controlPlaneInstanceProfile"] = opts.ControlPlaneInstanceProfile
	templateValues["nodeInstanceProfile"] = opts.NodeInstanceProfile
	templateValues["kopsStateStore"] = opts.KopsStateStore
	templateValues["imageRepo"] = opts.ImageRepo
	templateValues["dockerConfig"] = opts.DockerConfig

	for k, v := range opts.prowJobCommonTemplateValues() {
		templateValues[k] = v
	}

	return templateValues
}

type EksDistroRebuildProwJobOptions struct {
	*ProwJobCommonOptions
	ArtifactsBucket             string
	ControlPlaneInstanceProfile string
	DockerConfig                string
	GitRepo                     string
	ImageRepo                   string
	KopsStateStore              string
	NodeInstanceProfile         string
	TestRoleArn                 string
}

func (b *EksDistroRebuildProwJobOptions) setEksDRebuildOptionsDefaults() {
	if b.ProwJobCommonOptions == nil {
		b.ProwJobCommonOptions = &ProwJobCommonOptions{}
	}

	b.setCommonDefaults()

	if b.TestRoleArn == "" {
		b.TestRoleArn = constants.EksDTestRoleArn
	}

	if b.ArtifactsBucket == "" {
		b.ArtifactsBucket = constants.EksDPostSubmitArtifactsBucket
	}

	if b.ControlPlaneInstanceProfile == "" {
		b.ControlPlaneInstanceProfile = constants.EksDKopsControlPlaneRole
	}

	if b.GitRepo == "" {
		b.GitRepo = eksDRebuildDefaultGitRepo
	}

	if b.NodeInstanceProfile == "" {
		b.NodeInstanceProfile = constants.EksDKopsNodeInstanceProfile
	}

	if b.KopsStateStore == "" {
		b.KopsStateStore = constants.EKsDKopsStateStoreBucket
	}

	if b.ImageRepo == "" {
		b.ImageRepo = constants.EKsDBuildToolingImageRepo
	}

	if b.DockerConfig == "" {
		b.DockerConfig = constants.DockerConfigPath
	}
}
