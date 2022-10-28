package prowJobs

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/aws/eks-distro-build-tooling/golang/conformance-test-executor/pkg/constants"
)

const eksDRebuildProwJobTemplate = "pkg/prowJobs/templates/eks-distro-rebuild.yaml"

func NewEksDistroRebuildProwJob(kubernetesVersion string, jobName string, opts *EksDistroRebuildProwJobOptions) ([]byte, error) {
	if opts == nil {
		opts = &EksDistroRebuildProwJobOptions{}
	}
	opts.setEksDRebuildOptionsDefaults()
	templateData := EksDistroRebuildTemplateValues(*opts)

	templateData["startTime"] = ProwJobStartTime()
	templateData["kubernetesVersion"] = kubernetesVersion
	templateData["jobName"] = jobName

	temp, err := template.ParseFiles(eksDRebuildProwJobTemplate)
	if err != nil {
		return nil, fmt.Errorf("parsing template file: %v", err)
	}

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

	for k, v := range opts.TemplateValues() {
		templateValues[k] = v
	}

	return templateValues
}

type EksDistroRebuildProwJobOptions struct {
	*ProwJobOptions
	TestRoleArn                 string
	ArtifactsBucket             string
	ControlPlaneInstanceProfile string
	NodeInstanceProfile         string
	KopsStateStore              string
	ImageRepo                   string
	DockerConfig                string
}

func (b *EksDistroRebuildProwJobOptions) setEksDRebuildOptionsDefaults() {
	if b.ProwJobOptions == nil {
		b.ProwJobOptions = &ProwJobOptions{}
	}

	b.setDefaults()

	if b.TestRoleArn == "" {
		b.TestRoleArn = constants.TestRoleArn
	}

	if b.ArtifactsBucket == "" {
		b.ArtifactsBucket = constants.EksDPostSubmitArtifactsBucket
	}

	if b.ControlPlaneInstanceProfile == "" {
		b.ControlPlaneInstanceProfile = constants.ControlPlaneInstanceProfile
	}

	if b.NodeInstanceProfile == "" {
		b.NodeInstanceProfile = constants.KopsNodeInstanceProfile
	}

	if b.KopsStateStore == "" {
		b.KopsStateStore = constants.KopsStateStoreBucket
	}

	if b.ImageRepo == "" {
		b.ImageRepo = constants.ImageRepo
	}

	if b.DockerConfig == "" {
		b.DockerConfig = constants.DockerConfig
	}
}
