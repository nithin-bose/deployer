package k8s

import (
	"deployer/cmd/k8s/create"
	"deployer/cmd/k8s/delete"
	"deployer/cmd/k8s/deploy"
	"deployer/pkg/deploy/k8s"
	"log"

	"github.com/Songmu/prompter"
	"github.com/spf13/cobra"
)

var cluster string

func init() {
	RootCmd.PersistentFlags().StringVarP(&cluster, "cluster", "c", "", "Cluster to use. If not specified, uses default cluster configured by kubectl")
	RootCmd.AddCommand(create.RootCmd)
	RootCmd.AddCommand(delete.RootCmd)
	RootCmd.AddCommand(deploy.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "k8s",
	Short: "k8s tasks",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cluster != "" {
			if !prompter.YN("This will overwrite your ~/.kube. Are you sure?", false) {
				log.Fatal("Skipped")
			}
			err := k8s.SetupKubeConfig(cluster)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}
