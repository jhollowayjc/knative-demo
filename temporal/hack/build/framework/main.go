package main

import (
	"flag"
	"fmt"
	"os"

	"log/slog"

	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"

	workflow "workflow"
)

var (
	usage = "run\n\nRuns a Temporal Workflow."
	addr  = flag.String("temporal-addr", client.DefaultHostPort, "Temporal Server Addr [$TEMPORAL_ADDR]")
	ns    = flag.String("temporal-ns", client.DefaultNamespace, "Temporal Namespace [$TEMPORAL_NS]")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}

	parseEnv()   // override static defaults with environment.
	flag.Parse() // override env vars with flags.

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run a cloudevents client in receive mode which invokes
// the user-defined function.Handler on receipt of an event.
func run() error {
	// create client and worker
	c, err := client.Dial(client.Options{
		HostPort:  *addr,
		Namespace: *ns,
		Logger: tlog.NewStructuredLogger(
			slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			}))),
	})

	if err != nil {
		return fmt.Errorf("Unable to create Temporal Client: %w", err)
	}

	defer c.Close()

	w := worker.New(c, workflow.TaskQueueName, worker.Options{})

	workflow.RegisterWorkflow(w)

	return w.Run(worker.InterruptCh())
}

// parseEnv parses environment variables, populating the destination flags
// prior to the builtin flag parsing.  Invalid values exit 1.
func parseEnv() {
	if val, ok := os.LookupEnv("TEMPORAL_ADDR"); ok {
		*addr = val
	}

	if val, ok := os.LookupEnv("TEMPORAL_NAMESPACE"); ok {
		*ns = val
	}
}
