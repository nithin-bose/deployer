package k8s

import (
	"deployer/pkg"
	"deployer/pkg/deploy"
	"errors"
	"fmt"
)

func InfraApp(directory string, cloudPlatorm string, app string) error {
	chart := GetInfraChart(directory, cloudPlatorm, app)
	command := fmt.Sprintf("helm upgrade --install %s %s", app, chart)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func SystemApp(directory string, app string, environment string) error {
	chart := GetSystemChart(directory, app)
	valFile := GetValFilePath(chart, pkg.ProductionValsFile)
	if environment != "production" {
		valFile = GetValFilePath(chart, pkg.StagingValsFile)
	}
	command := fmt.Sprintf("helm upgrade -f %s --install %s-%s %s --namespace kube-system", valFile, app, environment, chart)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func DeployServiceApp(directory string, force bool, ci bool, environment string, app string, version string) error {
	err := deploy.ValidateEnvironment(environment)
	if err != nil {
		return err
	}
	if !force && environment == "production" && version == "latest" {
		return errors.New("Only versioned releases should be deployed to production")
	}
	fmt.Sprintf("Deploying %s... ", environment)

	if ci {
		err = SetupKubeConfig(environment)
		if err != nil {
			return err
		}
	}

	var command string
	chart := GetServiceChart(directory, app)
	valFile := GetValFilePath(chart, pkg.ProductionValsFile)
	if environment != "production" {
		valFile = GetValFilePath(chart, pkg.StagingValsFile)
	}

	command = fmt.Sprintf("helm upgrade -f %s --set image.tag=%s --install %s-%s %s", valFile, version, app, environment, chart)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func AdminPanel(directory string, app string) error {
	chart := GetAdminPanelChart(directory, app)
	command := fmt.Sprintf("helm upgrade --install %s %s", app, chart)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func DeleteAppWithEnvironment(environment string, app string) error {
	command := fmt.Sprintf("helm delete %s-%s", app, environment)
	fmt.Println(command, "\n")
	return pkg.Execute(command)
}

func DeleteApp(app string) error {
	command := fmt.Sprintf("helm delete %s", app)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}
