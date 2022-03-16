package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello, micron server")
	})

	http.ListenAndServe(":8080", nil)
}
