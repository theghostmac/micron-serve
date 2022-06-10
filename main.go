package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("server is running...")
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(writer, "Hello there %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye API fetched...")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
