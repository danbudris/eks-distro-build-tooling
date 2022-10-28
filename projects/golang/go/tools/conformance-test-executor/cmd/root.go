package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "prow-job-runner",
	Short:            "Run custom prow-jobs on-demand",
	Long:             `Command to construct and run custom EKS-D prow-jobs on-demand`,
	PersistentPreRun: rootPersistentPreRun,
}

func init() {
}

func rootPersistentPreRun(cmd *cobra.Command, args []string) {
}

func Execute() error {
	return rootCmd.Execute()
}