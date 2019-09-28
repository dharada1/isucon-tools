package main

import (
	"fmt"
	"net/http"
)

// Goでnet/httpを使う時のこまごまとした注意
// https://qiita.com/ono_matope/items/60e96c01b43c64ed1d18

func main() {
	nethttpパッケージのパラメータチューニング()
	レスポンスを受け取ったら必ずBodyをCloseする()
}

// WIP ...
func nethttpパッケージのパラメータチューニング() {
	// http.Clientはパッケージ側にグローバル変数として存在する(スレッドセーフ).
	// const DefaultMaxIdleConnsPerHost = 2 がデフォルトなので、適切に変更してやる.
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000

	// DefaultTransport, Timeout, MaxIdleConns, MaxIdleConnsPerHost, IdleConnTimeout
	// など他にも

	// 事例 (適切な値いくつだろ)
	/*
	   外部 API へのリクエストをする際に HTTP Agent の MaxIdleConns と MaxIdleConnsPerHost を拡張、1 ホスト 3000 コネクションまで持てるようにした。
	   https://yosuke-furukawa.hatenablog.com/entry/2019/09/08/232231
	*/

	// 自分で用意したhttp.Clientの制限を変更する場合以下のようになるらしい(未確認)
	// c := http.Client{
	// 	Transport: http.Transport{MaxIdleConnsPerHost: 32},
	// }
	// _, _ = c.Get("http://example.com")

	// 参考
	// https://gist.github.com/catatsuy/e627aaf118fbe001f2e7c665fda48146#httpclient%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6
}

func レスポンスを受け取ったら必ずBodyをCloseする() {
	// 呼び出し側でCloseしてあげないとTCPコネクションが再利用されないため.
	// https://qiita.com/ono_matope/items/60e96c01b43c64ed1d18

	f := func(url string) error {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close() // 明示的にクローズ

		return nil
	}

	url := "https://pairs.lv/1.0/master"
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		f(url)
	}
}
