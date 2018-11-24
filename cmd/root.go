package cmd

import (
	"deployer/cmd/cluster"
	"deployer/cmd/delete"
	"deployer/cmd/deploy"
	"deployer/cmd/promote"
	"deployer/cmd/release"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(cluster.RootCmd)
	RootCmd.AddCommand(release.RootCmd)
	RootCmd.AddCommand(promote.RootCmd)
	RootCmd.AddCommand(deploy.RootCmd)
	RootCmd.AddCommand(delete.RootCmd)
}

// RootCmd is the root CLI command
var RootCmd = &cobra.Command{
	Use:           "alt",
	Short:         "AltShifter stuff that just needs to be easy",
	SilenceUsage:  true,
	SilenceErrors: true,
}
