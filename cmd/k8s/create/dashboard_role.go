package create

import (
	"deployer/pkg/deploy/k8s"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(dashboardRoleCmd)
}

var dashboardRoleCmd = &cobra.Command{
	Use:   "dashboard-role",
	Short: "Create a role binding for kube-system:kubernetes-dashboard to have full access",
	Run: func(cmd *cobra.Command, args []string) {
		k8s.SetRoleForDashboard()
	},
}
