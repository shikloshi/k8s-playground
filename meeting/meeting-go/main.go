package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Info("We are in the root path")
	json.NewEncoder(w).Encode(&Response{
		Message: "You are in the root path",
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Info("we are doing an health check")
	json.NewEncoder(w).Encode(&Response{
		Message: "Health check is okay",
	})
}

func goToMeeting(w http.ResponseWriter, r *http.Request) {
	log.Info("We are going to meeting")
	json.NewEncoder(w).Encode(&Response{
		Message: "You are going to a meeting now",
	})
}

func main() {

	port := getEnvWithDefault("PORT", "3001")

	log.Infof("Going to start meeting-v1 service on port: %s", port)

	router := mux.NewRouter()

	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/meeting", goToMeeting).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func getEnvWithDefault(key, defaultValue string) string {
	envVar := os.Getenv(key)
	if len(envVar) == 0 {
		return defaultValue
	}
	return envVar
}
