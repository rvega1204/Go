package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response represents the standard structure for HTTP responses.
// It includes the status code, data to be returned, and any additional message.
type Response struct {
	Status      int                 `json:"status"`  // HTTP status code
	Data        interface{}         `json:"data"`    // Data to be returned in the response body
	Message     string              `json:"message"` // Message providing additional context (e.g., error message)
	contentType string              // Content type of the response (usually "application/json")
	respWrite   http.ResponseWriter // The response writer to send the response
}

// CreateDefaultResponse initializes a Response with default values.
// The status is set to OK (200), content type is set to "application/json",
// and it takes the ResponseWriter to send the response.
func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,      // Default status is 200 OK
		respWrite:   rw,                 // Response writer to send the response
		contentType: "application/json", // Default content type is JSON
	}
}

// Send sends the Response to the client.
// It sets the response headers, marshals the Response struct to JSON,
// and writes the response to the client.
func (resp *Response) Send() {
	// Set the Content-Type header for the response
	resp.respWrite.Header().Set("Content-Type", resp.contentType)
	// Set the HTTP status code for the response
	resp.respWrite.WriteHeader(resp.Status)

	// Marshal the response to JSON
	output, _ := json.Marshal(&resp)
	// Write the JSON output to the response body
	fmt.Fprintln(resp.respWrite, string(output))
}

// SendData creates a default Response, assigns the provided data to it,
// and sends it as the response to the client.
func SendData(rw http.ResponseWriter, data interface{}) {
	// Create a default Response with the provided ResponseWriter
	response := CreateDefaultResponse(rw)
	// Assign the data to the Response
	response.Data = data
	// Send the response to the client
	response.Send()
}

// NotFound sets the Response status to HTTP 404 (Not Found)
// and adds a default "Resource not found" message.
func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound   // Set status code to 404
	resp.Message = "Resource not found" // Set the default not found message
}

// SendNotFound creates a default Response, sets it to "Not Found" (404),
// and sends the response to the client.
func SendNotFound(rw http.ResponseWriter) {
	// Create a default Response
	response := CreateDefaultResponse(rw)
	// Set the response to "Not Found"
	response.NotFound()
	// Send the "Not Found" response to the client
	response.Send()
}

// UnprocessableEntity sets the Response status to HTTP 422 (Unprocessable Entity)
// and adds a default "UnprocessableEntity not found" message.
func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity   // Set status code to 422
	resp.Message = "UnprocessableEntity not found" // Set the default unprocessable entity message
}

// SendUnprocessableEntity creates a default Response, sets it to "Unprocessable Entity" (422),
// and sends the response to the client.
func SendUnprocessableEntity(rw http.ResponseWriter) {
	// Create a default Response
	response := CreateDefaultResponse(rw)
	// Set the response to "Unprocessable Entity"
	response.UnprocessableEntity()
	// Send the "Unprocessable Entity" response to the client
	response.Send()
}
