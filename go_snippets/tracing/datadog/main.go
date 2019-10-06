package main

import (
	"log"
	"net/http"

	goji "goji.io"
	"goji.io/pat"
)

func main() {
	InitAppTracer()

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hoge"), getHogeHandler)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
