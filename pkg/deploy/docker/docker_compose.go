package docker

import (
	"deployer/pkg"
	"fmt"
)

func DeployServiceApp(composeFile string, service string) error {
	var command string
	var err error

	if composeFile != "" {
		command = fmt.Sprintf("docker compose -f %s pull %s", composeFile, service)
	} else {
		command = fmt.Sprintf("docker compose pull %s", service)
	}
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		return err
	}

	if composeFile != "" {
		command = fmt.Sprintf("docker compose -f %s up -d %s", composeFile, service)
	} else {
		command = fmt.Sprintf("docker compose up -d %s", service)
	}
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}
