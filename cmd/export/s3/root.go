package s3

import (
	"deployer/cmd/export/s3/all"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(all.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "s3",
	Short: "Export to s3",
}
