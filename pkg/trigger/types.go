package trigger

type WebhookResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message"`
}
