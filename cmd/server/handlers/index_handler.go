package handlers

import (
	"net/http"

	"github.com/hudolfhess/golang-study/cmd/server/application"
)

type IndexHandler struct {
	Config application.Config
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("<h1>Simple web application!</h1><p>...</p>"))
}
