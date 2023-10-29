package main

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the token is present in the request's cookies
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tokenString := cookie.Value

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) { // any is equivalent to interface{}
			return jwtSecretKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Attach the user information to the request context
		ctx := context.WithValue(r.Context(), "user", claims.Username)
		r = r.WithContext(ctx)
		// Call the next handler
		next.ServeHTTP(w, r)
	}
}
