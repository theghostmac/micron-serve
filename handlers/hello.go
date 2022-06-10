package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("server is running...")
	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer, "Hello there %s", d)
}
