package main

import (
	"log"
	"net/http"

	"github.com/lu-moreira/shouldgo/context/gocontextsample/server"
)

func main() {
	handler := server.HandleSearch

	if err := http.ListenAndServe(":8080", http.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}
