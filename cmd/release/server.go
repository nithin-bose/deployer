package release

import (
	"deployer/pkg"
	"deployer/pkg/release"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Release for a server app. Command must have exactly 3 arguments, the release type, projectID and user ID.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 && len(args) > 4 {
			pkg.FatalF("Command must have atleast 3 arguments, the release type, projectID and user ID. It can also have an optional releaser name  \n")
		}

		if len(args) == 3 {
			args[3] = ""
		}
		release.CheckForGitRepo()
		release.ValidateType(args[0])

		release.Create(args[0], args[1], args[2], args[3])
	},
}
