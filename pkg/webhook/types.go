package webhook

type Auth struct {
	AccessKey   string `json:"access_key"`
	AccessToken string `json:"access_token"`
}
type DockerRequest struct {
	Auth
	ComposeFile string `json:"compose_file"`
	App         string `json:"app"`
}
