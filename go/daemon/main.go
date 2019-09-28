// inspired by https://gist.github.com/catatsuy/e627aaf118fbe001f2e7c665fda48146#isucon8%E6%9C%AC%E9%81%B8%E3%81%A7%E5%AE%9F%E9%9A%9B%E3%81%AB%E4%BD%BF%E3%81%A3%E3%81%9F%E3%82%A4%E3%83%B3%E3%83%A1%E3%83%A2%E3%83%AA%E3%82%AD%E3%83%A3%E3%83%83%E3%82%B7%E3%83%A5%E5%AE%9F%E8%A3%85

package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutexをfieldに持つことで、Mutexが使える.
type items struct {
	sync.Mutex // TODO ... RWMutex使うパターンもやってみる
	items      []string
}

func NewItems() *items {
	s := make([]string, 0, 100)
	i := &items{
		items: s,
	}
	return i
}

func (i *items) Append(s string) {
	i.Lock()
	i.items = append(i.items, s)
	i.Unlock()
}

func (i *items) Count() int {
	i.Lock()
	l := len(i.items)
	i.Unlock()

	return l
}

// グローバル変数にitemsGlobalを持ち、
var itemsGlobal = NewItems()

// 2つのデーモンから操作する
func main() {
	AppenderDaemon()
	AppenderDaemon2()

	CounterDaemon()

	// これ書かないとmainが終了して落ちる
	for {
	}
}

func AppenderDaemon() {
	c := time.Tick(1 * time.Millisecond)
	go func() {
		for {
			// fmt.Println("added! from appender daemon 1")
			itemsGlobal.Append("hoge")

			<-c // time.Tickを待ち受けることでsleepしている
		}
	}()
}

func AppenderDaemon2() {
	c := time.Tick(2 * time.Millisecond)
	go func() {
		for {
			// fmt.Println("added! from appender daemon 2")
			itemsGlobal.Append("fuga")

			<-c // time.Tickを待ち受けることでsleepしている
		}
	}()
}

func CounterDaemon() {
	c := time.Tick(5 * time.Second)
	go func() {
		for {
			cnt := itemsGlobal.Count()
			fmt.Println(cnt)

			<-c // time.Tickを待ち受けることでsleepしている
		}
	}()
}
