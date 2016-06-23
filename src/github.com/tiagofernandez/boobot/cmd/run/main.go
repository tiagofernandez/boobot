package main

import (
	"github.com/tiagofernandez/boobot"
	"os"
)

func main() {
	boobot.Start(os.Args[1])
}
