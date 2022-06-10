package main

import (
	"github.com/theghostmac/micron-serve/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
