package all

import (
	"deployer/pkg/deploy"
	"deployer/pkg/deploy/k8s"
	"fmt"
	"io/ioutil"
	"log"

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
			log.Fatal("Command must have exactly 1 argument, environment.  \n")
		}
		err := deploy.ValidateEnvironment(args[0])
		if err != nil {
			log.Fatal(err)
		}

		dir := chartsDir + "charts/services/"
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal("An error occurred:\n %s \n", err.Error())
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
			log.Fatal("Nothing found! \n")
		}
		if prompter.YN("Are you sure you want to deploy all of the above?", false) {
			for _, app := range apps {
				err = k8s.DeployServiceApp(chartsDir, true, false, args[0], app, "latest")
				if err != nil {
					log.Fatal("An error occurred:\n %s \n", err.Error())
				}
			}
		}
	},
}
