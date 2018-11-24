package delete

import (
	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deployment delete tasks. Helm is required to be installed and configured",
}
