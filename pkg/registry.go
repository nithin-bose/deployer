package pkg

import (
	"github.com/Songmu/prompter"
)

//Docker registry details
type DockerRegistryDetails struct {
	Host     string
	User     string
	Password string
	Email    string
}

func GetDockerRegistryDetails() *DockerRegistryDetails {
	details := &DockerRegistryDetails{}
	details.Host = prompter.Prompt("Enter docker registry server IP/host name", "")
	if details.Host == "" {
		FatalF("Registry server IP/host name is required. \n")
	}
	details.User = prompter.Prompt("Enter docker registry user name", "")
	if details.User == "" {
		FatalF("Registry server user name is required. \n")
	}

	details.Password = prompter.Prompt("Enter docker registry password", "")
	if details.Password == "" {
		FatalF("Registry server password is required. \n")
	}

	details.Email = prompter.Prompt("Enter docker registry email", "")
	if details.Email == "" {
		FatalF("Registry server email is required. \n")
	}
	return details
}
