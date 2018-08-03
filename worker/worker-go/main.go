package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const numberOfMeetings int = 3

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
	var message string
	skipHeader := r.Header.Get("x-skip-meeting")
	log.Infof("x-skip-meeting header value: %s", skipHeader)
	if len(skipHeader) != 0 {
		message = "Skipped meeting due to header"
		log.Info(message)
	} else {
		for i := 0; i < numberOfMeetings; i++ {
			message = "Going to a meeting"
			log.Infof("We are going to meeting: %d", i)
		}
	}
	json.NewEncoder(w).Encode(&Response{
		Message: message,
	})
}

func main() {

	port := getEnvWithDefault("PORT", "3001")

	log.Infof("Going to start worker-v1 service on port: %s", port)

	router := mux.NewRouter()

	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/work", goToMeeting).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func getEnvWithDefault(key, defaultValue string) string {
	envVar := os.Getenv(key)
	if len(envVar) == 0 {
		return defaultValue
	}
	return envVar
}
