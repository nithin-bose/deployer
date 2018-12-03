package trigger

import (
	"log"
	"os"

	"github.com/parnurzeal/gorequest"
)

func getDockerAuthFields() map[string]string {
	auth := make(map[string]string)
	auth["access_key"] = os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_KEY")
	auth["access_token"] = os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_TOKEN")
	return auth
}

func DockerDeployApp(app string) {
	req := gorequest.New()

	body := getDockerAuthFields()
	body["app"] = app

	url := os.Getenv("DEPLOYER_WEBHOOK_DOCKER_URL") + "/docker/deploy/app"
	resp := WebhookResponse{}
	_, _, errs := req.Post(url).
		Send(body).
		EndStruct(&resp)

	if errs != nil {
		log.Printf("%s", errs[0])
		os.Exit(2)
	}

	if !resp.Success {
		log.Printf("%s", resp.ErrorMessage)
		os.Exit(2)
	}
}
