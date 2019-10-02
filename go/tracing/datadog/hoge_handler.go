package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type resHoge struct {
	Hoge string `json:"hoge"`
}

func getHogeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()

	defer tracer.Flush()

	span := tracer.Trace(ctx).NewRootSpanFromRequest(r)
	defer func() {
		span.SetResponse(200).Finish()
	}()

	for i := 0; i < 10; i++ {
		子(ctx)
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	res := resHoge{Hoge: "hoge"}
	json.NewEncoder(w).Encode(res)
}

func 子(ctx context.Context) int {
	defer tracer.Trace(ctx).NewChildSpan("子").Finish()

	time.Sleep(500 * time.Millisecond)

	return 孫(ctx) + 孫(ctx)
}

func 孫(ctx context.Context) int {
	defer tracer.Trace(ctx).NewChildSpan("孫").Finish()

	time.Sleep(100 * time.Millisecond)

	return rand.Intn(10)
}
