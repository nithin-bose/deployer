package cmd

import (
	"deployer/cmd/docker"
	"deployer/cmd/export"
	"deployer/cmd/k8s"
	"deployer/cmd/promote"
	"deployer/cmd/release"
	"deployer/cmd/webhook"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(k8s.RootCmd)
	RootCmd.AddCommand(docker.RootCmd)
	RootCmd.AddCommand(release.RootCmd)
	RootCmd.AddCommand(promote.RootCmd)
	RootCmd.AddCommand(webhook.RootCmd)
	RootCmd.AddCommand(export.RootCmd)
}

// RootCmd is the root CLI command
var RootCmd = &cobra.Command{
	Use:           "deployer",
	Short:         "Deployment related stuff that just needs to be easy",
	SilenceUsage:  true,
	SilenceErrors: true,
}
