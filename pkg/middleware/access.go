package middleware

import (
	"net/http"
	"github.com/michee/micgram/pkg/model"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assuming you have some way to identify admin, like a role in User model
		userId := r.Header.Get("userId")
		user, _ := model.GetUserById(userId)
		if user.Role != true {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
