package deploy

import (
	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Trigger deploy web hooks",
}
