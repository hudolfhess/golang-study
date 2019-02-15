package server

import (
	"net/http"

	"github.com/hudolfhess/golang-study/cmd/server/handlers"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", &handlers.IndexHandler{})

	return mux
}
