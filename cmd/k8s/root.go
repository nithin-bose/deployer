package k8s

import (
	"deployer/cmd/k8s/create"
	"deployer/cmd/k8s/delete"
	"deployer/cmd/k8s/deploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(create.RootCmd)
	RootCmd.AddCommand(delete.RootCmd)
	RootCmd.AddCommand(deploy.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "k8s",
	Short: "k8s tasks",
}
