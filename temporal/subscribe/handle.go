package function

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.temporal.io/sdk/client"
)

var temporalClient client.Client

func init() {
	var err error

	addr := client.DefaultHostPort
	if val, ok := os.LookupEnv("TEMPORAL_ADDR"); ok {
		addr = val
	}

	// create client and worker
	temporalClient, err = client.Dial(client.Options{
		HostPort: addr,
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create Temporal Client: %w", err))
	}
}

type RequestData struct {
	Email string `json:"email"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// TODO: this is from the worker itself, so how can we share these
type EmailDetails struct {
	EmailAddress      string `json:"emailAddress"`
	Message           string `json:"message"`
	IsSubscribed      bool   `json:"isSubscribed"`
	SubscriptionCount int    `json:"subscriptionCount"`
}

const TaskQueueName = "email_subscription"

// Handle an HTTP Request.
func Handle(w http.ResponseWriter, r *http.Request) {
	// ensure JSON request
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type, expecting application/json", http.StatusUnsupportedMediaType)
		return
	}

	var requestData RequestData

	// decode request into variable
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error processing request body", http.StatusBadRequest)
		return
	}

	// check if the email is blank
	if requestData.Email == "" {
		http.Error(w, "Email is blank", http.StatusBadRequest)
		return
	}

	// use the email as the id in the workflow.
	workflowOptions := client.StartWorkflowOptions{
		ID:                                       requestData.Email,
		TaskQueue:                                TaskQueueName,
		WorkflowExecutionErrorWhenAlreadyStarted: true,
	}

	// Define the EmailDetails struct
	subscription := EmailDetails{
		EmailAddress:      requestData.Email,
		Message:           "Welcome to the Subscription Workflow!",
		SubscriptionCount: 0,
		IsSubscribed:      true,
	}

	// Execute the Temporal Workflow to start the subscription.
	_, err = temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, "SubscriptionWorkflow", subscription)

	if err != nil {
		http.Error(w, "Couldn't sign up user. Please try again.", http.StatusInternalServerError)
		log.Print(err)
		return
	}

	// build response
	responseData := ResponseData{
		Status:  "success",
		Message: "Signed up.",
	}

	// send headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created status code

	// send response
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Print("Could not encode response JSON", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
