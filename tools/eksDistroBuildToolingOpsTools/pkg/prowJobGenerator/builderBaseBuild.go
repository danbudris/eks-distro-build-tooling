package prowJobGenerator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/aws/eks-distro-build-tooling/tools/eksDistroBuildToolingOpsTools/pkg/constants"
)

const eksDistroRebuildTemplate = "templates/eks-distro-rebuild-and-test.yaml"
const EksDistroEcrPublicPushRoleArn = "arn:aws:iam::832188789588:role/ECRPublicPushRole"

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

func (e *BuilderBaseBuildProwJobOptions) setBuilderBaseBuildDefaults() {
	if e.ProwJobCommonOptions == nil {
		e.ProwJobCommonOptions = &ProwJobCommonOptions{}
	}
	e.setCommonDefaults()
	if e.ArtifactDeploymentRoleArn == "" {
		e.ArtifactDeploymentRoleArn = constants.EksDArtifactDeploymentRoleArn
	}

	if e.ArtifactsBucket == "" {
		e.ArtifactsBucket = constants.EksDPostSubmitArtifactsBucket
	}

	if e.ImageRepo == "" {
		e.ImageRepo = constants.EKsDBuildToolingImageRepo
	}

	if e.AwsRegion == "" {
		e.AwsRegion = constants.DefaultAwsRegion
	}

	if !e.RebuildAll {
		e.RebuildAll = true
	}
}

// BuilderBaseBuildProwJob generates a yaml defining a triggered post-submit for the builder-base build,
// using the provided options to populate the template. This job yaml can then be applied to a prow cluster to run the custom job.
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