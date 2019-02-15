package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hudolfhess/golang-study/cmd/server"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Simple web application!")
}

func main() {
	mux := server.Routes()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
