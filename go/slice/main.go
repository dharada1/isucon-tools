// https://gist.github.com/catatsuy/e627aaf118fbe001f2e7c665fda48146#%E3%82%B9%E3%83%A9%E3%82%A4%E3%82%B9%E3%81%AElen%E3%81%A8cap
// https://blog.golang.org/go-slices-usage-and-internals
// https://speakerdeck.com/kaznishi/180713-lt?slide=10
package main

import "fmt"

func main() {
	capを適切に設定する()
	capの倍々ゲームサンプル()
}

func capを適切に設定する() {
	// capを適切に設定する.
	l := 0   // appendしてくやつは基本0
	c := 100 // 取りうる最大値が予測できれば。
	_ = make([]string, l, c)
}

func capの倍々ゲームサンプル() {
	fmt.Println("sample from https://gist.github.com/catatsuy/e627aaf118fbe001f2e7c665fda48146#%E3%82%B9%E3%83%A9%E3%82%A4%E3%82%B9%E3%81%AElen%E3%81%A8cap")
	a := make([]int, 0)
	prev := -1
	for i := 0; i < 10000; i++ {
		if prev != cap(a) {
			prev = cap(a)
			fmt.Println(cap(a))
		}
		a = append(a, i)
	}
}
