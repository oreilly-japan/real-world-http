package main

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
	"image"
	"image/png"
	"net/http"
	//"net/http/httputil"
)

var images [][]byte

func init() {
	images = make([][]byte, 4)
	texts := []string{
		"image for page1",
		"image for page2",
		"image for page1(p)",
		"image for page2(p)",
	}
	tt, _ := truetype.Parse(gomono.TTF)
	face := truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	for i := 0; i < 4; i++ {
		dst := image.NewRGBA(image.Rect(0, 0, 260, 25))
		d := font.Drawer{
			Dst:  dst,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.P(0, 20),
		}
		d.DrawString(texts[i])
		var buffer bytes.Buffer
		png.Encode(&buffer, dst)
		images[i] = buffer.Bytes()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	/*dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))*/
	switch r.URL.Path {
	case "/":
		pusher, ok := w.(http.Pusher)
		if ok {
			pusher.Push("/image1", nil)
			pusher.Push("/image2", nil)
		}
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, `<html><body>
                            <img src="/image1">
                            <a href="/next">next</a>
                        </body></html>`)
	case "/next":
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<html><body><img src="/image2"></body></html>`)
	case "/image1":
		w.Header().Set("Content-Type", "image/png")
		if r.Header.Get("User-Agent") != "" {
			w.Write(images[0])
		} else {
			w.Write(images[2])
		}
	case "/image2":
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Cache-Control", "max-age=20")
		//w.Header().Set("etag", "image2")
		if r.Header.Get("User-Agent") != "" {
			w.Write(images[1])
		} else {
			w.Write(images[3])
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
	fmt.Println(err)
}
