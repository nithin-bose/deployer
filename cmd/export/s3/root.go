package s3

import (
	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "s3",
	Short: "Export to s3",
}
