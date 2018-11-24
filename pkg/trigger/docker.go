package trigger

import (
	"deployer/pkg"
	"os"

	"github.com/parnurzeal/gorequest"
)

func getAuthFields() map[string]string {
	auth := make(map[string]string)
	auth["access_key"] = os.Getenv("DEPLOYER_WEBHOOK_ACCESS_KEY")
	auth["access_token"] = os.Getenv("DEPLOYER_WEBHOOK_ACCESS_TOKEN")
	return auth
}

func DockerDeployApp(composeFile string, app string) {
	req := gorequest.New()

	body := getAuthFields()
	body["app"] = app
	if composeFile != "" {
		body["compose_file"] = composeFile
	}
	url := os.Getenv("DEPLOYER_WEBHOOK_URL") + "/docker/deploy/app"
	resp := WebhookResponse{}
	_, _, errs := req.Post(url).
		Send(body).
		EndStruct(&resp)

	if errs != nil {
		pkg.FatalF("%s", errs[0])
	}

	if !resp.Success {
		pkg.FatalF("%s", resp.ErrorMessage)
	}
}
