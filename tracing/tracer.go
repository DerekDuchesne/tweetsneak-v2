package tracing

import (
	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

var (
	exporter *stackdriver.Exporter
)

func init() {
	// Don't export the logs to StackDriver if we're not in production.
	if !settings.UseProduction {
		return
	}

	// Create a StackDriver exporter.
	e, err := stackdriver.NewExporter(stackdriver.Options{})
	if err != nil {
		panic(err)
	}
	exporter = e

	// Register the exporter.
	trace.RegisterExporter(exporter)

	// Apply a config that will always sample the traces.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
