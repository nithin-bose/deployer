package webhook

type Auth struct {
	AccessKey   string `json:"access_key"`
	AccessToken string `json:"access_token"`
}

type DockerRequest struct {
	Auth
	ComposeFile    string `json:"compose_file"`
	ComposeFileDir string `json:"compose_file_dir"`
	App            string `json:"app"`
}

type K8sRequest struct {
	Auth
	Environment string `json:"environment"`
	App         string `json:"app"`
	Version     string `json:"version"`
}
