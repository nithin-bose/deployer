package k8s

import (
	"deployer/pkg"
	"fmt"
)

func InfraApp(directory string, cloudPlatorm string, app string) {
	chart := GetInfraChart(directory, cloudPlatorm, app)
	command := fmt.Sprintf("helm upgrade --install %s %s", app, chart)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func SystemApp(directory string, app string, environment string) {
	chart := GetSystemChart(directory, app)
	valFile := GetValFilePath(chart, pkg.ProductionValsFile)
	if environment != "production" {
		valFile = GetValFilePath(chart, pkg.StagingValsFile)
	}
	command := fmt.Sprintf("helm upgrade -f %s --install %s-%s %s --namespace kube-system", valFile, app, environment, chart)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func ServiceApp(directory string, environment string, app string, version string) {
	var command string
	var err error
	chart := GetServiceChart(directory, app)
	valFile := GetValFilePath(chart, pkg.ProductionValsFile)
	if environment != "production" {
		valFile = GetValFilePath(chart, pkg.StagingValsFile)
	}

	command = fmt.Sprintf("helm upgrade -f %s --set image.tag=%s --install %s-%s %s", valFile, version, app, environment, chart)
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func AdminPanel(directory string, app string) {
	chart := GetAdminPanelChart(directory, app)
	command := fmt.Sprintf("helm upgrade --install %s %s", app, chart)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func DeleteAppWithEnvironment(environment string, app string) {
	command := fmt.Sprintf("helm delete %s-%s", app, environment)
	fmt.Println(command, "\n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func DeleteApp(app string) {
	command := fmt.Sprintf("helm delete %s", app)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}