package boobot

import (
	"io"
	"net/http"
)

func HttpGet(url string, f func(io.ReadCloser) bool) (bool, error) {
	r, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer r.Body.Close()
	return f(r.Body), err
}
