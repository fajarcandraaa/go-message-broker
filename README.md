# go-message-broker
## About The Project
Is a simple way to implement message-broker in the golang programming language.
In that case, i use Google Cloud Pub/Sub

### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Golang](https://golang.com)

## Getting Started

This is how you may give instructions on setting up your project locally. To get a local copy up and running follow these simple example steps.

### Prerequisites
1. Set the variable `subscribtionName` in `/handler/subscriber.go`, line 29 with your Pub/Sub Topic Subscription Name value
2. Run Publisher Service, that is in the main directory
    ```sh
   go run main.go
   ```
3. In other terminal, run Subscriber Service in go-message-broker/subscriber directory
    ```sh
   go run subscriber/main.go
   ```
4. Or You can go to the directory, than run
    ```sh
   go run main.go
   ```
   
## List of Endpoint to be tested

- URL : {host}/publish
- Method        : POST
- Body          : { "topic" : `your pubsub topic`, "message" : "" }
- Body type     : json

## Contact

Name - @fajarcandraaa - fajarcandraaa@gmail.com
