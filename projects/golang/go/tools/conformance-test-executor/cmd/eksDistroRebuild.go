package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/aws/eks-distro-build-tooling/golang/conformance-test-executor/pkg/prowJobs"
)

var eksDistroRebuildCustomProwJob = &cobra.Command{
	Use:   "eks-distro-rebuild",
	Short: "Rebuild EKS Distro",
	Long:  "Generate a prow job to re-build all of EKS Distro for the given K8s version",
	PreRun: preRunEksDistroRebuild,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := eksDistroRebuildProwJob(cmd.Context())
		if err != nil {
			log.Fatalf("Error getting image versions: %v", err)
		}
		return nil
	},
}

const (
	ArtifactsBucketFlag = "artifactsBucket"
	K8sVersionFlag      = "kubernetesVersion"
)

func init() {
	customProwJobCommand.AddCommand(eksDistroRebuildCustomProwJob)
	eksDistroRebuildCustomProwJob.Flags().StringP(K8sVersionFlag, "v", "", "EKS D Kubernetes version to rebuild")
	eksDistroRebuildCustomProwJob.Flags().StringP(ArtifactsBucketFlag, "b", "", "EKS-D artifacts bucket to use for generated artifacts")

	requiredFlags := []string{
		K8sVersionFlag,
	}

	for _, flag := range requiredFlags {
		if err := eksDistroRebuildCustomProwJob.MarkFlagRequired(flag); err != nil {
			log.Fatalf("failed to mark flag %v as required: %v", flag, err)
		}
	}
}

func preRunEksDistroRebuild(cmd *cobra.Command, args []string) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		err := viper.BindPFlag(flag.Name, flag)
		if err != nil {
			log.Fatalf("Error initializing flags: %v", err)
		}
	})
}

func eksDistroRebuildProwJob(ctx context.Context) error {
	eksDistroRebuildOpts := &prowJobs.EksDistroRebuildProwJobOptions{
		ProwJobCommonOptions: &prowJobs.ProwJobCommonOptions{},
	}

	artifactsBucket := viper.GetString(ArtifactsBucketFlag)
	runtimeImage := viper.GetString(RuntimeImageFlag)
	k8sVersion := viper.GetString(K8sVersionFlag)

	if runtimeImage != "" {
		eksDistroRebuildOpts.RuntimeImage = runtimeImage
	}

	if artifactsBucket != "" {
		eksDistroRebuildOpts.ArtifactsBucket = artifactsBucket
	}

	headSha := "5e5bbfc56809daec14982b258412a589e97f82a8"
	baseSha := "36f6355201f6c244eec34a7eafb6a5673928900b"
	jobName := fmt.Sprintf("build-%s-postsubmit-custom", k8sVersion)

	jobBytes, err := prowJobs.NewEksDistroRebuildProwJob(k8sVersion, jobName, headSha, baseSha, eksDistroRebuildOpts)
	if err != nil {
		return fmt.Errorf("building EKS Distro Rebuild Prow Job: %v", err)
	}
	fmt.Println(string(jobBytes))
	return nil
}
