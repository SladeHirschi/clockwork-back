package api

import (
	"net/http"
	"strings"
)

func (c *ClockworkApi) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth check for login and signup
		if isPublicRoute(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		// Example auth check: look for Authorization header
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Real token verification logic would go here

		next.ServeHTTP(w, r)
	})
}

func isPublicRoute(path string) bool {
	switch path {
	case "/login", "/signup":
		return true
	default:
		return false
	}
}
