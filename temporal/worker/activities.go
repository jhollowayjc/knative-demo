package workflow

import (
	"context"

	"go.temporal.io/sdk/activity"
)

// workflow activities

// SendEmail sends an email
func SendEmail(ctx context.Context, emailInfo EmailDetails) (string, error) {
	activity.GetLogger(ctx).Info("Sending email to customer", "EmailAddress", emailInfo.EmailAddress)
	return "Email sent to " + emailInfo.EmailAddress, nil
}
