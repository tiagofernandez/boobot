package goatbot

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"strings"
)

type ImageJson struct {
	Id      int    `json:"id"`
	Preview string `json:"preview"`
}
type ImagesJson []ImageJson

func RandomBoobsUrl() string {
	url := "https://cdn.boldomatic.com/content/post/AnZxFQ-7c08343dc0d8ff17bc01e946b0260d7ddf622e9e8cdbc549a2b9976985374fca/Error-404-Tits-not-found"
	HttpGet("http://api.oboobs.ru/noise/1", func(body io.ReadCloser) {
		var data ImagesJson
		err := json.NewDecoder(body).Decode(&data)
		if err == nil {
			url = "http://media.oboobs.ru/" + strings.Replace(data[0].Preview, "_preview", "", -1)
		} else {
			log.Println(err)
		}
	})
	log.Println(url)
	return url
}

func RandomBoobsImage() []byte {
	url := RandomBoobsUrl()
	img := make([]byte, 0)
	HttpGet(url, func(body io.ReadCloser) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
		img = buf.Bytes()
	})
	return img
}
