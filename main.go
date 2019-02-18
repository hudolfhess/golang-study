package main

import (
	"log"
	"net/http"

	"github.com/hudolfhess/golang-study/cmd/server"
)

func main() {
	config := server.LoadConfig()
	mux := server.Routes(config)

	log.Println("Starting server: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
