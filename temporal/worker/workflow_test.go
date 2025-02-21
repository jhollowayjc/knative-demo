package workflow_test

import (
	"testing"
	"time"
	"workflow"

	"go.temporal.io/sdk/testsuite"
)

func Test_CanceledSubscriptionWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := workflow.EmailDetails{
		EmailAddress:      "example@temporal.io",
		Message:           "This is a test to see if the Workflow cancels. This is dependent on the bool variable in the testDetails struct.",
		IsSubscribed:      true,
		SubscriptionCount: 12,
	}

	// set delayed callback to allow time for cancellation.
	env.RegisterDelayedCallback(func() {
		env.CancelWorkflow()
	}, 5*time.Second)

	env.RegisterWorkflow(workflow.SubscriptionWorkflow)
	env.RegisterActivity(workflow.SendEmail)

	env.ExecuteWorkflow(workflow.SubscriptionWorkflow, testDetails)
}
