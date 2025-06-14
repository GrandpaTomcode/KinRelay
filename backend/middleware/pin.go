package middleware

import (
	"net/http"
)

const validPin = "123456" // Replace later with env-based value

func RequirePIN(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pin := r.Header.Get("X-PIN")
		if pin != validPin {
			http.Error(w, "Unauthorised", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}) 
}