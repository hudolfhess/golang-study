package application

import "net/http"

type Middleware func(next http.Handler) http.HandlerFunc
