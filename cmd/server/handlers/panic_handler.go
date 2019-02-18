package handlers

import (
	"net/http"

	"github.com/hudolfhess/golang-study/cmd/server/application"
)

type PanicHandler struct {
	Config application.Config
}

func (h *PanicHandler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	panic("Unexpected error!!!")
}
