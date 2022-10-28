package prowJobs

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/aws/eks-distro-build-tooling/golang/conformance-test-executor/pkg/constants"
)

const eksDistroRebuildTemplate = "templates/eks-distro-rebuild-and-test.yaml"

type BuilderBaseBuildProwJobOptions struct {
	*ProwJobCommonOptions
	ArtifactDeploymentRoleArn   string
	ArtifactsBucket             string
	AwsRegion                   string
	ImageRepo                   string
	ReleaseEnvironment          string
	ReleaseBranch               string
	RebuildAll                  bool
}

func (e BuilderBaseBuildProwJobOptions) setBuilderBaseBuildDefaults() {
	if e.ProwJobCommonOptions == nil {
		e.ProwJobCommonOptions = &ProwJobCommonOptions{}
	}
	e.setCommonDefaults()
	if e.ArtifactDeploymentRoleArn == "" {
		e.ArtifactDeploymentRoleArn = constants.ArtifactDeploymentRoleArn
	}

	if e.ArtifactsBucket == "" {
		e.ArtifactsBucket = constants.EksDPostSubmitArtifactsBucket
	}

	if e.ImageRepo == "" {
		e.ImageRepo = constants.ImageRepo
	}

	if e.AwsRegion == "" {
		e.AwsRegion = constants.AwsRegion
	}

	if !e.RebuildAll {
		e.RebuildAll = true
	}
}

func BuilderBaseBuildProwJob(opts BuilderBaseBuildProwJobOptions) ([]byte, error) {
	opts.setBuilderBaseBuildDefaults()

	templateData := make(map[string]interface{})

	var renderedTemplateData bytes.Buffer

	temp, err := template.ParseFiles(eksDistroRebuildTemplate)
	if err != nil {
		return nil, fmt.Errorf("parsing template file: %v", err)
	}
	err = temp.ExecuteTemplate(&renderedTemplateData, "eks-distro-rebuild-and-test.yaml", templateData)
	if err != nil {
		return nil, fmt.Errorf("rendering builderBaseProwJob: %v", err)
	}

	return renderedTemplateData.Bytes(), nil
}