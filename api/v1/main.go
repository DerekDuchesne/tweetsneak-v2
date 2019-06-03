package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"

	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/restapi"
	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/restapi/operations"
)

func main() {
	// Load the API spec.
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Create the API.
	api := operations.NewSearchAPI(swaggerSpec)

	// Create the server.
	server := restapi.NewServer(api)

	// Shut down the server once the app terminates.
	defer server.Shutdown()

	// Configure the API.
	server.ConfigureAPI()

	// Set the port.
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8000
	}
	server.Port = port

	// Start the server.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
