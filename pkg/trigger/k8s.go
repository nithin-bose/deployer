package trigger

import (
	"errors"
	"os"

	"github.com/parnurzeal/gorequest"
)

func getK8sAuthFields() map[string]string {
	auth := make(map[string]string)
	auth["access_key"] = os.Getenv("DEPLOYER_WEBHOOK_K8S_ACCESS_KEY")
	auth["access_token"] = os.Getenv("DEPLOYER_WEBHOOK_K8S_ACCESS_TOKEN")
	return auth
}

func K8sDeployApp(environment string, app string, version string) error {
	req := gorequest.New()

	body := getK8sAuthFields()
	body["environment"] = environment
	body["app"] = app
	body["version"] = version

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
