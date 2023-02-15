package docker

import (
	"deployer/pkg"
	"fmt"
	"os"
)

func DeployServiceApp(dockerStacksDir string, composeFile string, app string, service string) error {
	var command string
	var err error

	composeFileDir := dockerStacksDir + string(os.PathSeparator) + app
	err = os.Chdir(composeFileDir)
	if err != nil {
		return err
	}

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
		command = fmt.Sprintf("docker compose -f %s up --remove-orphans -d %s", composeFile, service)
	} else {
		command = fmt.Sprintf("docker compose up --remove-orphans -d %s", service)
	}
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}
