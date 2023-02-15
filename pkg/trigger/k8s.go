package trigger

import (
	"errors"
	"os"

	"github.com/parnurzeal/gorequest"
)

func getBodyWithK8sAuthFields() *WebhookK8sRequest {
	auth := &WebhookK8sRequest{}
	auth.AccessKey = os.Getenv("DEPLOYER_WEBHOOK_K8S_ACCESS_KEY")
	auth.AccessToken = os.Getenv("DEPLOYER_WEBHOOK_K8S_ACCESS_TOKEN")
	return auth
}

func K8sDeployApp(environment string, app string, version string) error {
	req := gorequest.New()

	body := getBodyWithK8sAuthFields()
	body.Environment = environment
	body.App = app
	body.Version = version

	url := os.Getenv("DEPLOYER_WEBHOOK_K8S_URL") + "/k8s/deploy/app"
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
