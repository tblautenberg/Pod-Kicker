package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	// Placeholder for API endpoints

	// Loading the URL of the image that we need to check webhooks from. We use bufio for I/O
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the full URL of the hub.docker.com repo you want to monitor:\n ")
	input, _ := reader.ReadString('\n')

	// Fast removal of the spaces in each end of user-input
	input = strings.TrimSpace(input)

	fmt.Printf("The CICD pipeline is now configured to litsen to ----->  %s\n", input)
	fmt.Printf("The server will now start\n")

	// Creation of the server
	port := ":8450"
	fmt.Printf("Server is now online and litsening for incomming requests on httpL://localhost%s 🖥️\n", port)
	err := http.ListenAndServe(port, nil)
	// If err is not nill, print the following line and exit the program
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not start the server on %s\n %v", port, err))
	}
}
