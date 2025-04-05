package middleware

import (
	"net/http"

	userModels "medicine/pkg/user"
)

// idk.
func Authentification() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := userModels.User{
				ID:          userModels.VanekID,
				IsAnonymous: false,
			}

			newCtx := userModels.StoreInContext(r.Context(), user)
			r = r.WithContext(newCtx)

			next.ServeHTTP(w, r)
		})
	}
}
