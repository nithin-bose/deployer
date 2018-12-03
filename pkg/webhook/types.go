package webhook

type Auth struct {
	AccessKey   string `json:"access_key"`
	AccessToken string `json:"access_token"`
}

type DockerRequest struct {
	Auth
	App string `json:"app"`
}

type K8sRequest struct {
	Auth
	Environment string `json:"environment"`
	App         string `json:"app"`
	Version     string `json:"version"`
}
