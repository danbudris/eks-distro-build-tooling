package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"log"

	"github.com/spf13/cobra"
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
	K8sVersionFlag = "kubernetesVersion"
)

func init() {
	customProwJobCommand.AddCommand(eksDistroRebuildCustomProwJob)
	eksDistroRebuildCustomProwJob.Flags().StringP(K8sVersionFlag, "v", "", "EKS D Kubernetes version to rebuild")

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
		ProwJobOptions: &prowJobs.ProwJobOptions{},
	}

	runtimeImage := viper.GetString(RuntimeImageFlag)
	k8sVersion := viper.GetString(K8sVersionFlag)
	fmt.Printf("\n\n K8s Version: %s \n\n", k8sVersion)

	if runtimeImage != "" {
		eksDistroRebuildOpts.RuntimeImage = runtimeImage
	}

	jobBytes, err := prowJobs.NewEksDistroRebuildProwJob(k8sVersion, "testJob", eksDistroRebuildOpts)
	if err != nil {
		return fmt.Errorf("building EKS Distro Rebuild Prow Job: %v", err)
	}
	fmt.Println(string(jobBytes))
	return nil
}
