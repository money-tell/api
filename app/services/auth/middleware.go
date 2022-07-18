package auth

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
)

func (a *Auth) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			return
		}

		userId, ok := token.Get("id")
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		_, err = a.storage.FindUserByID(r.Context(), userId.(string))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
