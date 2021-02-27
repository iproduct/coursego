package rest

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/iproduct/coursego/12-modules-rest-jwtauth/model"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const BEARER_SCHEMA = "Bearer "
		//var header = r.Header.Get("x-access-token") //Grab the token from the header
		var header = r.Header.Get("Authorization") //Grab the token from the header

		var token string
		if header == "" || len(header) <= len(BEARER_SCHEMA) {
			respondWithError(w, http.StatusForbidden, "Missing auth token")
			return
		}
		token = header[len(BEARER_SCHEMA):]
		token = strings.TrimSpace(token)

		if token == "" {
			//Token is missing, returns with error code 403 Unauthorized
			respondWithError(w, http.StatusUnauthorized, "Missing auth token")
			return
		}
		claims := &model.UserToken{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			respondWithError(w, http.StatusForbidden, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
