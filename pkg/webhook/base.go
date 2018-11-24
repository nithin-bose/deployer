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
	w.Write(r)
}

func RenderError(w http.ResponseWriter, err error) {
	log.Println("Error to render", err)
	response := make(map[string]interface{})
	response["success"] = false
	response["error_message"] = err.Error()
	r, _ := json.Marshal(response)
	w.Write(r)
}

func authenticate(key string, token string) error {
	if key != os.Getenv("DEPLOYER_WEBHOOK_ACCESS_KEY") || token != os.Getenv("DEPLOYER_WEBHOOK_ACCESS_TOKEN") {
		return errors.New("Authentication error")
	}
	return nil
}
