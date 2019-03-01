package export

import (
	"deployer/cmd/export/s3"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(s3.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "export",
	Short: "Export releases",
}
