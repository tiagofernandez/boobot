package main

import (
	"github.com/tiagofernandez/goatbot"
	"os"
)

func main() {
	goatbot.Start(os.Args[1])
}
