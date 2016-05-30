package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "👏👏👏")
	})
	http.HandleFunc("/boobs", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "( 🍥 Y 🍥 )")
	})
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}
