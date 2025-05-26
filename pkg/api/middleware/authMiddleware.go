package middleware

import (
	"go_final_project/pkg/api"
	"go_final_project/pkg/api/auth"
	"net/http"
	"os"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var jwt string

			cookie, err := req.Cookie("token")
			if err != nil {
				api.WriteJSON(res, http.StatusUnauthorized, map[string]string{"error": "токен не найден в куках"})
				return
			}
			jwt = cookie.Value
			var valid bool
			valid, err = auth.ValidateToken(jwt, pass)
			if !valid || err != nil {
				api.WriteJSON(res, http.StatusUnauthorized, map[string]string{"error": "токен не валиден"})
				return
			}
		}
		next(res, req)
	})
}
