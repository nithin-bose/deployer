package k8s

import (
	"deployer/pkg"
	"errors"
	"fmt"
	"strings"
)

func InfraApp(directory string, cloudPlatorm string, app string) error {
	chart := GetInfraChart(directory, cloudPlatorm, app)
	command := fmt.Sprintf("helm upgrade --install %s %s", app, chart)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func SystemApp(directory string, app string, environment string) error {
	chart := GetSystemChart(directory, app)
	
	defaultValFilePath, err := GetDefaultValFilePath(chart)
	if err != nil {
		return err
	}
	
	valFilePath, err := GetValFilePath(chart, environment)
	if err != nil {
		return err
	}
	command := fmt.Sprintf("helm upgrade -f %s -f %s --install %s-%s %s --namespace kube-system", defaultValFilePath, valFilePath, app, environment, chart)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func CommonApp(directory string, app string, environment string) error {
	chart := GetCommonChart(directory, app)

	defaultValFilePath, err := GetDefaultValFilePath(chart)
	if err != nil {
		return err
	}

	valFilePath, err := GetValFilePath(chart, environment)
	if err != nil {
		return err
	}
	command := fmt.Sprintf("helm upgrade -f %s -f %s --install %s-%s %s", defaultValFilePath, valFilePath, app, environment, chart)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func DeployServiceApp(directory string, force bool, ci bool, environment string, app string, version string) error {
	if !force && strings.Contains(environment, "production") && version == "latest" {
		return errors.New("Only versioned releases should be deployed to production")
	}
	fmt.Sprintf("Deploying %s... ", environment)

	var err error
	if ci {
		err = SetupKubeConfig(environment)
		if err != nil {
			return err
		}
	}

	var command string
	chart := GetServiceChart(directory, app)

	defaultValFilePath, err := GetDefaultValFilePath(chart)
	if err != nil {
		return err
	}

	valFilePath, err := GetValFilePath(chart, environment)
	if err != nil {
		return err
	}

	command = fmt.Sprintf("helm upgrade -f %s -f %s --set image.tag=%s --install %s-%s %s", defaultValFilePath, valFilePath, version, app, environment, chart)
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
