package main

import (
	"./boobot"
	"os"
)

func main() {
	boobot.Start(os.Args[1])
}
