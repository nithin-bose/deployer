package cluster

import (
	"deployer/cmd/cluster/create"
	"deployer/cmd/cluster/delete"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(create.RootCmd)
	RootCmd.AddCommand(delete.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Cluster tasks",
}
