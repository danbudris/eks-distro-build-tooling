package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var customProwJobCommand = &cobra.Command{
	Use:   "generate-custom-prow-job",
	Short: "Generate a custom prow job",
	Long:  "Given base options, generate a custom prow job template",
}

const (
	BranchFlag       = "branch"
	JobNameFlag      = "jobName"
	RuntimeImageFlag = "runtimeImage"
)

func init() {
	customProwJobCommand.PersistentFlags().String(BranchFlag, "main", "Branch to use as the head for the prow job comparison")
	customProwJobCommand.PersistentFlags().StringP(RuntimeImageFlag, "r", "", "Runtime image to use as the base of the prow job")
	customProwJobCommand.PersistentFlags().StringP(JobNameFlag, "n", "", "Name for the executed prow job")
	if err := viper.BindPFlags(customProwJobCommand.PersistentFlags()); err != nil {
		log.Fatalf("failed to bind flags for root: %v", err)
	}
	rootCmd.AddCommand(customProwJobCommand)
}