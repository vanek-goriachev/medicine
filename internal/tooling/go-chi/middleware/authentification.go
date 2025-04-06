package middleware

import (
	"net/http"

	userModels "medicine/pkg/user"
)

// I am not sure if i have to go to iam in this middleware
// Maybe i should use jwt tokens
// Maybe sessions

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
