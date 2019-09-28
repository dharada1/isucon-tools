package main

import (
	"testing"
)

func Benchmark直列(b *testing.B) {
	_ = GetContent(url1)
	_ = GetContent(url2)
}

func Benchmark並列(b *testing.B) {
	c := make(chan string)
	go GetContentWithChannel(url1, c)
	go GetContentWithChannel(url2, c)
	_, _ = <-c, <-c
}
