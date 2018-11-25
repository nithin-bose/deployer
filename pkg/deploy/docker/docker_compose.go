package docker

import (
	"deployer/pkg"
	"fmt"
)

func DeployServiceApp(composeFile string, app string) {
	var command string
	var err error

	if composeFile != "" {
		command = fmt.Sprintf("docker-compose -f %s pull %s", composeFile, app)
	} else {
		command = fmt.Sprintf("docker-compose pull %s", app)
	}
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	if composeFile != "" {
		command = fmt.Sprintf("docker-compose -f %s up -d %s", composeFile, app)
	} else {
		command = fmt.Sprintf("docker-compose up -d %s", app)
	}
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}
