package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type resHoge struct {
	Hoge string `json:"hoge"`
}

func getHogeHandler(w http.ResponseWriter, r *http.Request) {
	// defer tracer.Close() // ???
	defer tracer.Flush() // Handlerの処理が終わるタイミングでtracer.Flush()によってトレース情報を送信させる. (1req 1送信)

	// span := tracer.Trace(r.Context()).NewRootSpanFromRequest(r)
	trace := tracer.Trace(r.Context())
	span := trace.NewRootSpanFromRequest(r)
	defer func() {
		span.SetResponse(200).Finish()
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		子(span.Context())
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	res := resHoge{Hoge: "hoge"}
	json.NewEncoder(w).Encode(res)
}

func 子(ctx context.Context) {
	span := tracer.Trace(ctx).NewChildSpan("子")
	defer span.Finish()

	fmt.Println("子")

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		孫(ctx)
	}
}

func 孫(ctx context.Context) {
	span := tracer.Trace(ctx).NewChildSpan("孫")
	defer span.Finish()

	fmt.Println("孫")

	time.Sleep(100 * time.Millisecond)
}
