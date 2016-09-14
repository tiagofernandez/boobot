package goatbot

import (
	"io"
	"log"
	"net/http"
)

func HttpGet(url string, f func(io.ReadCloser)) {
	r, err := http.Get(url)
	if err == nil {
		defer r.Body.Close()
		f(r.Body)
	} else {
		log.Println(err)
	}
}
