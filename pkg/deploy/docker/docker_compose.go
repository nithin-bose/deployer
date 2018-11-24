package docker

import (
	"deployer/pkg"
	"fmt"
)

func DeployServiceApp(composeFile string, app string) {
	var command string
	var err error

	command = fmt.Sprintf("docker-compose pull %s", app)
	if composeFile != "" {
		command = command + fmt.Sprintf(" -f %s", composeFile)
	}
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	command = fmt.Sprintf("docker-compose up -d %s", app)
	if composeFile != "" {
		command = command + fmt.Sprintf(" -f %s", composeFile)
	}
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}
