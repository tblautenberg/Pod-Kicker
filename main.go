package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Global vars to hold the repo URL and deployment name
var (
	expectedRepoURL string
	deploymentName  string
)

func main() {
	// Read config from environment variables instead of stdin
	expectedRepoURL = os.Getenv("REPO_URL")
	deploymentName = os.Getenv("DEPLOYMENT_NAME")

	if expectedRepoURL == "" || deploymentName == "" {
		log.Fatal("Environment variables REPO_URL and DEPLOYMENT_NAME must be set")
	}

	fmt.Printf("The CICD pipeline is now configured to listen to webhooks from ----->  %s\n", expectedRepoURL)
	fmt.Printf("When payload is received for the given repo, the app will restart -----> %s\n", deploymentName)
	fmt.Printf("The server will now start\n")

	// Setup HTTP handler
	http.HandleFunc("/updateDeployment", updateDeployment)

	// Start server
	port := ":8450"
	fmt.Printf("Server is now online and listening for incoming requests on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not start the server on %s\n %v", port, err))
	}
}
