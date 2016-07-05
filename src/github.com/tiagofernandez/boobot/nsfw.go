package boobot

import (
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

func GimmeSomeBoobs() string {
	url := "http://tinyurl.com/tits-not-found"
	HttpGet("http://api.oboobs.ru/noise/1", func(body io.ReadCloser) {
		var data ImagesJson
		err := json.NewDecoder(body).Decode(&data)
		if err == nil {
			url = "http://media.oboobs.ru/" + strings.Replace(data[0].Preview, "_preview", "", -1)
		} else {
			log.Println(err)
		}
	})
	return url
}
