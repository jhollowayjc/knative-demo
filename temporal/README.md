# temporal


## Setup and Deploy Temporal
```bash
kubectl create ns temporal
kustomize build --enable-helm ./hack/deploy | kubectl apply -n temporal -f -
```


## Build and deploy worker

```bash
./worker/build
kubectl apply -n default -f ./worker/service.yaml
```

## Deploy functions
```bash
func deploy --push=false --path=./subscribe
func deploy --push=false --path=./details
func deploy --push=false --path=./unsubscribe
```


# Notes

The worker is a (somewhat crufty) experiment designed to see if we can emulate the same UX
by having the main package for the `worker` abstracted away.

Effectively, all a `worker` needs is the following `register.go`:
```go
package workflow

// TaskQueueName is the queue name used to identify the worker
const TaskQueueName string = "email_subscription"

// RegisterWorkflow the workflow and activity
func RegisterWorkflow(w worker.Worker) {
	w.RegisterWorkflow(YourWorkflowHere)
	w.RegisterActivity(YourActivity)
}
```

During build, we drop a `main.go` that wires up the worker and starts running it.
See `hack/build/framework` for this code.

We'd probably want to use buildpack and follow the same sort of pattern `func` do if we like this.
