package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// sendData is a helper function that sends a JSON response to the client.
// It takes an HTTP response writer, the data to send, and the HTTP status code as parameters.
// It sets the "Content-Type" to "application/json", writes the status code,
// marshals the data to JSON format, and writes the JSON response to the client.
func sendData(rw http.ResponseWriter, data interface{}, status int) {
	// Set the Content-Type header to application/json.
	rw.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code for the response.
	rw.WriteHeader(status)

	// Marshal the data to JSON format. If marshaling fails, it silently ignores the error.
	output, _ := json.Marshal(&data)
	// Write the JSON response to the client.
	fmt.Fprintln(rw, string(output))
}

// sendError sends a simple error message as a response when a resource is not found.
// It takes an HTTP response writer and the status code as parameters.
// It writes the status code and a default error message "Resource not found" to the client.
func sendError(rw http.ResponseWriter, status int) {
	// Set the HTTP status code for the response.
	rw.WriteHeader(status)
	// Write a simple error message to the client.
	fmt.Fprintln(rw, "Resource not found")
}
