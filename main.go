package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fajarcandraaa/go-message-broker/config"
	"github.com/fajarcandraaa/go-message-broker/handler"
	"github.com/fajarcandraaa/go-message-broker/ucase/publisher"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// func projectId(id string) string {
// 	return id
// }

func main() {
	_ = godotenv.Load()
	router := mux.NewRouter()
	env := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	jsonFile, err := os.Open(env)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	projectID := fmt.Sprintf("%v", result["project_id"])

	client := config.CreatePubSubClient(projectID)
	service := publisher.NewPublisher(client)
	handler := handler.NewPublisherHandler(service)

	router.HandleFunc("/publish", handler.PublishHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

	// publisher.NewPubSource().Subscriber(projectID)
}
