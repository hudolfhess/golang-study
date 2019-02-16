package server

import (
	"net/http"

	"github.com/hudolfhess/golang-study/cmd/server/handlers"
	"github.com/hudolfhess/golang-study/cmd/server/middlewares"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", runWithMiddewares(&handlers.IndexHandler{}))
	mux.Handle("/panic", runWithMiddewares(&handlers.PanicHandler{}))

	return mux
}

func runWithMiddewares(handler http.Handler) http.Handler {
	return middlewares.ErrorMiddleware(middlewares.LogMiddleware(handler))
}
