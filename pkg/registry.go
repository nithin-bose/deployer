package pkg

import (
	"log"

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
		log.Fatal("Registry server IP/host name is required. \n")
	}
	details.User = prompter.Prompt("Enter docker registry user name", "")
	if details.User == "" {
		log.Fatal("Registry server user name is required. \n")
	}

	details.Password = prompter.Prompt("Enter docker registry password", "")
	if details.Password == "" {
		log.Fatal("Registry server password is required. \n")
	}

	details.Email = prompter.Prompt("Enter docker registry email", "")
	if details.Email == "" {
		log.Fatal("Registry server email is required. \n")
	}
	return details
}
