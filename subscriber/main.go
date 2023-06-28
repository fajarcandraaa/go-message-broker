package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fajarcandraaa/go-message-broker/handler"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	// router := mux.NewRouter()
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

	// subscriber := config.CreatePubSubClient(projectID)
	sub := handler.NewSubscriberHandler(projectID)
	sub.Subscriber()
}
