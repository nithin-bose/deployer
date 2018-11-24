package promote

import (
	"deployer/pkg"
	"deployer/pkg/release"

	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote for release. Command must have exactly 4 arguments, source branch, target branch, projectID and user ID.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 && len(args) > 5 {
			pkg.FatalF("Command must have atleast 4 arguments, source branch, target branch, projectID and user ID. Optional promoter name \n")
		}

		if len(args) == 4 {
			args[4] = ""
		}
		release.Promote(args[0], args[1], args[2], args[3], args[4])
	},
}
