package main

import (
	"fmt"
	"net/http"
)

func updateDeployment(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		fmt.Printf("We just recived a not valid HTTP request")
		http.Error(writer, "Method needs to be a HTTP POST - check your webhook to see if the configuration is setup correctly", http.StatusMethodNotAllowed)
		return
	}
}
