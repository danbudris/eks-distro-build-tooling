package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/aws/eks-distro-build-tooling/tools/eksDistroBuildToolingOpsTools/pkg/consumerUpdater"
	"github.com/aws/eks-distro-build-tooling/tools/eksDistroBuildToolingOpsTools/pkg/eksDistroRelease"
)

var (
	backportCmd = &cobra.Command{
		Use:   "update",
		Short: "Update consumers of EKS Distro",
		Long:  "Tool for updating consumers of EKS Distro generated artifacts",
		RunE: func(cmd *cobra.Command, args []string) error {
			var eksDReleases []*eksDistroRelease.Release
			for _, v := range viper.GetStringSlice(eksDistroReleasesFlag) {
				r, err := eksDistroRelease.NewEksDistroReleaseObject(v)
				if err != nil {
					return err
				}
				eksDReleases = append(eksDReleases, r)
			}
			bottlerocketUpdater := consumerUpdater.NewBottleRocketUpdater(eksDReleases)
			var err error
			for _, u := range bottlerocketUpdater.Updaters() {
				err = u.Update()
			}
			return err
		},
	}
)

func init() {
	rootCmd.AddCommand(backportCmd)
}
