package webhook

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/docker/deploy/app", DockerDeployAppHandler).Methods("POST")
	r.HandleFunc("/k8s/deploy/app", K8sDeployAppHandler).Methods("POST")

	port := os.Getenv("DEPLOYER_WEBHOOK_PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Listening on port", port+"...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
