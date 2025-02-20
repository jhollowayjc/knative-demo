# knative demo

Demonstrating knative:

## Getting Started

### Dependencies


[!NOTE]: This has currently only been tested on docker-desktop with kubernetes enabled and using `kind` as the provisioning method.

```
brew tap knative-extensions/kn-plugins
brew install knative/client/kn
brew install func
```

## Install

### Knative operator:
```bash
kubectl apply -f https://github.com/knative/operator/releases/download/knative-v1.17.3/operator.yaml
```

### Knative Serving:
```
kubectl apply -f ./knative-serving.yaml
```

## Validation

### Deploy greeting function:

```bash
func deploy --push=false --path ./greeting
```
you should see output similar to the following:
```bash
Building function image
Still building
🙌 Function built: ko.local/knative-demo/greeting:latest
🎯 Creating Triggers on the cluster
✅ Function deployed in namespace "default" and exposed at URL:
   http://greeting.default.svc.cluster.local
```

### Invoke the function:
```bash
func invoke --path greeting
```
you should see output similar to the following:

```bash
{"Message":"Hello friend, nice to meet you"}
```

You can also provide a json payload with the `--data` argument
```bash
func invoke --path greeting --data '{"name": "Darth Vader"}'
```

## Next Steps

### Update `greeting`

1. Modify `greeting/handle.go`
2. Deploy the func with `func deploy --push=false --path ./greeting`
3. Invoke the func with `func invoke --path greeting`

### Create a new function

TODO
