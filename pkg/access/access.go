package access

import (
	"net/http"
	"github.com/go-chi/jwtauth/v5"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, "Vous n'etes pas autoriser", http.StatusUnauthorized)
			return
		}

		if !claims["isAdmin"].(bool) {
			http.Error(w, "You are not an administrator", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
