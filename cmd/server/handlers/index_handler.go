package handlers

import "net/http"

type IndexHandler struct {
}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("<h1>Simple web application!</h1>"))
}