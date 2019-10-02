package main

import (
	"context"

	"github.com/evalphobia/apptracer"
	"github.com/evalphobia/apptracer/platform/opencensus"
	"github.com/evalphobia/apptracer/platform/opencensus/datadog"
)

// global singleton tracer.
var tracer *apptracer.AppTracer

// Init initializes trace setting.
func InitTracer() {
	// init tracer
	tracer = apptracer.New(apptracer.Config{
		Enable:      true,
		ServiceName: "datch-apptracer-test",
		Version:     "v0.0.1",
		Environment: "stage",
	})
	// datadog (installed agent is needed)
	ddExp, err := datadog.NewExporter(context.Background(), "datch-apptracer-test")
	if err != nil {
		panic(err)
	}

	ocCli := opencensus.NewClient(ddExp)
	opencensus.SetSamplingRate(1) // 100%
	tracer.AddClient(ocCli)

	tracer.Flush()
}
