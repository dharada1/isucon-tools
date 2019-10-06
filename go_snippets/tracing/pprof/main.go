package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/profile"
	goji "goji.io"
	"goji.io/pat"
)

func main() {
	// pprof
	defer profile.Start(profile.ProfilePath(".")).Stop()

	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/hoge"), getHoge)
	mux.HandleFunc(pat.Get("/hoge2"), getHoge2)
	mux.HandleFunc(pat.Get("/hoge3"), getHoge3)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

type resHoge struct {
	Hoge  string `json:"hoge"`
	Hoge2 string `json:"hoge2"`
	Hoge3 string `json:"hoge3"`
}

func getHoge(w http.ResponseWriter, r *http.Request) {

	fmt.Println("start")
	for i := 0; i < 1000; i++ {
		fib(30)
	}

	childOfGetHoge()

	res := resHoge{
		Hoge:  "hoge",
		Hoge2: "fuga",
		Hoge3: "pyoyo",
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

func getHoge2(w http.ResponseWriter, r *http.Request) {

	fmt.Println("start")
	for i := 0; i < 1000; i++ {
		fib(30)
	}

	childOfGetHoge()

	res := resHoge{
		Hoge:  "hoge",
		Hoge2: "fuga",
		Hoge3: "pyoyo",
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

func getHoge3(w http.ResponseWriter, r *http.Request) {

	fmt.Println("start")
	for i := 0; i < 1000; i++ {
		fib(30)
	}

	childOfGetHoge()

	res := resHoge{
		Hoge:  "hoge",
		Hoge2: "fuga",
		Hoge3: "pyoyo",
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

func childOfGetHoge() {
	for i := 0; i < 100; i++ {
		fmt.Println("childOfGetHoge")
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
