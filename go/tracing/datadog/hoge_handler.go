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
	defer tracer.Flush() // Handlerの処理が終わるタイミングでtracer.Flush()によってトレース情報を送信させる. (1req 1送信)

	// Traceを作成
	trace := tracer.Trace(r.Context())

	// Spanの作成
	span := trace.NewRootSpanFromRequest(r)
	defer span.SetResponse(200).Finish()

	// SpanのContextを渡してfunc 子を呼ぶ
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
		// 孫(span.Context()) // Spanからchildを作りたいがうまく認識されない
		孫(ctx) // これだとRootのSpanからそのまま派生してしまう
	}
}

func 孫(ctx context.Context) {
	span := tracer.Trace(ctx).NewChildSpan("孫")
	defer span.Finish()

	fmt.Println("孫")

	time.Sleep(100 * time.Millisecond)
}
