package access


import (
	"net/http"
	"github.com/michee/micgram/pkg/model"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assuming you have some way to identify admin, like a role in User model
		userId := r.Header.Get("userId")
		user, err := model.GetUserById(userId)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}


		if !user.IsAdmin {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
