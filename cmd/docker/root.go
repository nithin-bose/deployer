package docker

import (
	"deployer/cmd/docker/deploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(deploy.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker tasks",
}
