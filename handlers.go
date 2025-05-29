package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// Struct to match the webhook JSON structure
type WebhookPayload struct {
	Repository struct {
		RepoURL string `json:"repo_url"`
	} `json:"repository"`
}

func updateDeployment(w http.ResponseWriter, r *http.Request) {
	// Validate request method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Only POST is supported.", http.StatusMethodNotAllowed)
		fmt.Println("Invalid request method")
		return
	}

	// Read the webhook request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		fmt.Println("Could not read request body:", err)
		return
	}

	// Parse the JSON payload
	var payload WebhookPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		fmt.Println("JSON parse error:", err)
		return
	}

	// Log the repo URL we received
	fmt.Printf("Webhook received from: %s\n", payload.Repository.RepoURL)

	// Compare repo from webhook with the expected one
	if payload.Repository.RepoURL != expectedRepoURL {
		http.Error(w, "Repo URL does not match configured repo", http.StatusForbidden)
		fmt.Println("Repo URL mismatch! Ignoring...")
		return
	}

	// Repo URL matches — time to restart the deployment
	fmt.Println("Repo URL matches, restarting deployment...")

	// Run the kubectl command to restart the deployment
	cmd := exec.Command("kubectl", "rollout", "restart", "-n", "default", "deployment", deploymentName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to restart deployment", http.StatusInternalServerError)
		fmt.Println("Failed to restart deployment:", err)
		fmt.Println("Command output:", string(output))
		return
	}

	// Command successful
	fmt.Println("Deployment restarted successfully!")
	fmt.Println("Output:", string(output))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deployment restarted successfully"))
}
