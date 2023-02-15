package trigger

import (
	"errors"
	"os"

	"github.com/parnurzeal/gorequest"
)

func getBodyWithDockerAuthFields() *WebhookDockerRequest {
	auth := &WebhookDockerRequest{}
	auth.AccessKey = os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_KEY")
	auth.AccessToken = os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_TOKEN")
	return auth
}

func DockerDeployApp(dockerStacksDir string, app string, service string) error {
	req := gorequest.New()

	body := getBodyWithDockerAuthFields()
	body.App = app
	body.Service = service
	body.DockerStacksDir = dockerStacksDir

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
