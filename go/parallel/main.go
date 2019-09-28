package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url1 = "https://www.bing.com/"
	url2 = "https://www.yahoo.co.jp"
)

func main() {
	fmt.Println("直列")
	_ = GetContent(url1)
	_ = GetContent(url2)
	fmt.Println("END!")

	fmt.Println("並列")
	c := make(chan string)
	go GetContentWithChannel(url1, c)
	go GetContentWithChannel(url2, c)
	_, _ = <-c, <-c // cから2つの値を待ち受け
	fmt.Println("END!")
}

func GetContent(url string) string {
	var s string
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return s
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	s = string(body[:])

	if err != nil {
		fmt.Println(err)
		return s
	}
	return s
}

// https://qiita.com/TsuyoshiUshio@github/items/6c04b7617db0062d3dee
func GetContentWithChannel(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	s := string(body[:])

	if err != nil {
		fmt.Println(err)
		return
	}
	c <- s
}
