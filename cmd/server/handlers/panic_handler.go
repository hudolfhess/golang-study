package handlers

import "net/http"

type PanicHandler struct {
}

func (h *PanicHandler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	panic("Unexpected error!!!")
}
