package server

import (
	"github.com/hudolfhess/golang-study/cmd/server/application"
	"github.com/hudolfhess/golang-study/cmd/server/middlewares"
)

func LoadConfig() application.Config {
	return application.Config{
		LogError: true,
		DefaultMiddlewares: []application.Middleware{
			middlewares.LogMiddleware,
			middlewares.ErrorMiddleware,
		},
	}
}
