package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(write http.ResponseWriter, request *http.Request) {
		log.Println("hello world!")
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(write, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(write, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":8080", nil)
}