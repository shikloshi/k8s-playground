package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

const numberOfMeetings int = 3

var meeting string

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
	skipHeader := r.Header.Get("x-skip-meeting")
	log.Infof("x-skip-meeting header value: %s", skipHeader)
	if len(skipHeader) != 0 {
		log.Info("skipped meeting due to header")
		json.NewEncoder(w).Encode(&Response{
			Message: "Skipped meeting due to header",
		})
		return
	}
	res, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", meeting, "meeting"), nil)
	log.Infof("response from meeting: %+v", res)
	if err != nil {
		log.Fatalf("Could not create a get request to meeting service: %+v", meeting)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading body - %+v", err)
	}

	log.Infof("response body: %s", body)
	res.Body.Close()
	json.NewEncoder(w).Encode(&Response{
		Message: "Encountered a meeting with response " + string(body),
	})
}

//func goToMeeting(w http.ResponseWriter, r *http.Request) {
//var message string
//skipHeader := r.Header.Get("x-skip-meeting")
//log.Infof("x-skip-meeting header value: %s", skipHeader)
//if len(skipHeader) != 0 {
//message = "Skipped meeting due to header"
//log.Info(message)
//} else {
//for i := 0; i < numberOfMeetings; i++ {
//message = "Going to a meeting"
//res, err := http.NewRequest("GET", meeting, nil)
//log.Infof("We are going to meeting: %d", i)
//if err != nil {
//log.Errorf("Error from meeting service: %+v ", err)
//message = "Error from meeting service, skipping"
//break
//} else {
//body, _ := ioutil.ReadAll(res.Body)
//log.Infof("response body: %s", body)
//res.Body.Close()
//}
//}
//}
//json.NewEncoder(w).Encode(&Response{
//Message: message,
//})
//}

func main() {

	port := getEnvWithDefault("PORT", "4000")
	meeting = getEnvWithDefault("MEETING_SERVICE_ADDRESS", "http://localhost:3000")

	log.Infof("Going to start worker-v1 service on port: %s", port)
	log.Infof("Meeting service address: %s", meeting)

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
