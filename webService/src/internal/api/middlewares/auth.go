package middlewares

import "net/http"

func AuthMiddleware(nextHandler http.Handler) http.Handler {
	return nextHandler
}