package goatbot

import (
	"encoding/json"
	"io"
	"testing"
)

type RandomJoke struct {
	Type  string `json:"type"`
	Value struct {
		Categories []interface{} `json:"categories"`
		ID         int           `json:"id"`
		Joke       string        `json:"joke"`
	} `json:"value"`
}

func TestHttpGet(t *testing.T) {
	HttpGet("http://api.icndb.com/jokes/random", func(body io.ReadCloser) {
		var data RandomJoke
		err := json.NewDecoder(body).Decode(&data)
		if err != nil {
			t.Error("Test failed")
		} else {
			b, _ := json.Marshal(data)
			t.Log(string(b))
		}
	})
}
