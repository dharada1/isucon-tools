package main

import (
	"context"

	"github.com/evalphobia/apptracer"
	"github.com/evalphobia/apptracer/platform/opencensus"
	"github.com/evalphobia/apptracer/platform/opencensus/datadog"
)

// global singleton tracer.
// AppTracer has multiple tracing platform clients to send tracing activity to the platforms.
// AppTracer -> opencensus client -> datadog Exporter
var tracer *apptracer.AppTracer

func InitAppTracer() {
	// Init tracer
	tracer = apptracer.New(apptracer.Config{
		Enable: true,
		// OpenCensusのみ使う場合はapptracer.Configは使われない?
		// ServiceName: "datch-apptracer-test-tracer", // https://app.datadoghq.com/apm/services のサービス名になる。
		// Version:     "v0.0.1",
		// Environment: "stage",
	})

	// Init ddExporter of OpenCensus
	ddExp, err := datadog.NewExporter(
		context.Background(),
		// datadog.Options{Service: "xxxxxxxxxxx"} がwrapされてる
		// https://app.datadoghq.com/apm/services のサービス名になる
		"datch-apptracer-test-ddexporter-9",
	)
	if err != nil {
		panic(err)
	}

	// OpenCensus client ( with Datadog exporter ) を AppTracer にもたせる.
	ocCli := opencensus.NewClient(ddExp) // (ocCli := opencensus.NewClient(ddExp, sdExp, xrayExp) みたく複数持てる)
	opencensus.SetSamplingRate(1)        // 100 % (テストなので)
	tracer.AddClient(ocCli)

	// datadog exporterに限ってはFlush() is dummy method (コード側でFlushの処理はやらずにDatadog Agentが溜めて送る感じなので)
	// だがXRayやStackDriver使う場合は必ず必要.
	tracer.Flush()
}
