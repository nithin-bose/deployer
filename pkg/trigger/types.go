package trigger

import "deployer/pkg/webhook"

type WebhookRequestBase struct {
	AccessKey   string `json:"access_key"`
	AccessToken string `json:"access_token"`
}

type WebhookDockerRequest struct {
	WebhookRequestBase
	webhook.DockerRequest
}

type WebhookK8sRequest struct {
	WebhookRequestBase
	webhook.K8sRequest
}

type WebhookResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message"`
}
