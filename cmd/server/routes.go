package server

import (
	"net/http"

	"github.com/hudolfhess/golang-study/cmd/server/application"
	"github.com/hudolfhess/golang-study/cmd/server/handlers"
)

func Routes(config application.Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", runWithMiddewares(&handlers.IndexHandler{}, config.DefaultMiddlewares))
	mux.Handle("/panic", runWithMiddewares(&handlers.PanicHandler{}, config.DefaultMiddlewares))

	return mux
}

func runWithMiddewares(handler http.Handler, executableMiddlewares []application.Middleware) http.Handler {
	if len(executableMiddlewares) > 0 {
		currentMiddleware := executableMiddlewares[0]
		nextMiddlewares := executableMiddlewares[1:]
		return runWithMiddewares(currentMiddleware(handler), nextMiddlewares)
	}
	return handler
}
