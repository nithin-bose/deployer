package all

import (
	"deployer/pkg"
	"deployer/pkg/deploy"
	"deployer/pkg/deploy/k8s"
	"fmt"
	"io/ioutil"

	"github.com/Songmu/prompter"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(appsCmd)
}

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Deploy all apps. Helm is required to be installed and configured",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			pkg.FatalF("Command must have exactly 1 argument, environment.  \n")
		}
		deploy.ValidateEnvironment(args[0])

		dir := chartsDir + "charts/services/"
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			pkg.FatalF("An error occurred:\n %s \n", err.Error())
		}

		var apps []string
		fmt.Println("Detected charts: ")
		for _, f := range files {
			if f.IsDir() {
				fmt.Println(f.Name())
				apps = append(apps, f.Name())
			}
		}
		if len(apps) == 0 {
			pkg.FatalF("Nothing found! \n")
		}
		if prompter.YN("Are you sure you want to deploy all of the above?", false) {
			for _, app := range apps {
				k8s.DeployServiceApp(chartsDir, true, false, args[0], app, "latest")
			}
		}
	},
}
