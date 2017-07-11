package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var image []byte

// 画像ファイルを準備
func init() {
	var err error
	// サーバーpush時に呼ばれる
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

// HTMLをブラウザに送信
// 画像をPUSHする
func handlerHtml(w http.ResponseWriter, r *http.Request) {
	// Pusherにキャスト可能であればPUSHする
	w.Header().Add("Content-Type", "text/html")
	pusher, ok := w.(http.Pusher)
	if ok {
		go func() {
			fmt.Println("Push /image")
			pusher.Push("/image", nil)
		}()
	}
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
	time.Sleep(time.Second)
}

// 画像ファイルをブラウザに送信
func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
	fmt.Println(err)
}
