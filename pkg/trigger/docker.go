package trigger

import (
	"errors"
	"os"

	"github.com/parnurzeal/gorequest"
)

func getDockerAuthFields() map[string]string {
	auth := make(map[string]string)
	auth["access_key"] = os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_KEY")
	auth["access_token"] = os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_TOKEN")
	return auth
}

func DockerDeployApp(dockerStacksDir string, app string, service string) error {
	req := gorequest.New()

	body := getDockerAuthFields()
	body["app"] = app
	body["service"] = service
	body["docker_stacks_dir"] = dockerStacksDir

	url := os.Getenv("DEPLOYER_WEBHOOK_DOCKER_URL") + "/docker/deploy/app"
	resp := WebhookResponse{}
	_, _, errs := req.Post(url).
		Send(body).
		EndStruct(&resp)

	if errs != nil {
		return errs[0]
	}

	if !resp.Success {
		return errors.New(resp.ErrorMessage)
	}
	return nil
}
