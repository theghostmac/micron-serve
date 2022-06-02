package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, ref *http.Request) {
	log.Println("Hello world")
	d, err := ioutil.ReadAll(ref.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}