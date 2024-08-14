package middlewares

import (
	"day06/configs"
	"net/http"
)

func AuthMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		insertedLogin, insertedPass, ok := r.BasicAuth()
		if ok {
			cfg := configs.GetConfig(nil)
			adminLogin := cfg.AdminUsername
			adminPass := cfg.AdminPassword
			if insertedLogin == adminLogin && insertedPass == adminPass {
				nextHandler.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
