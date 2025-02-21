package workflow

import (
	"go.temporal.io/sdk/worker"
)

// TaskQueueName is the queue name used to identify the worker
const TaskQueueName string = "email_subscription"

// RegisterWorkflow the workflow and activity
func RegisterWorkflow(w worker.Worker) {
	w.RegisterWorkflow(SubscriptionWorkflow)
	w.RegisterActivity(SendEmail)
}
