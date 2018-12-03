package webhook

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

func RenderSuccess(w http.ResponseWriter, data interface{}) {
	response := make(map[string]interface{})
	response["success"] = true
	response["data"] = data
	r, err := json.Marshal(response)
	if err != nil {
		RenderError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
}

func RenderError(w http.ResponseWriter, err error) {
	log.Println("Error to render", err)
	response := make(map[string]interface{})
	response["success"] = false
	response["error_message"] = err.Error()
	r, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
}

func authenticateDocker(key string, token string) error {
	if key != os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_KEY") || token != os.Getenv("DEPLOYER_WEBHOOK_DOCKER_ACCESS_TOKEN") {
		return errors.New("Authentication error")
	}
	return nil
}
