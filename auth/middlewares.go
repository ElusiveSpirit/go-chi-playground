package auth

import (
	"awesomeProject/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

/* Middleware handler to handle all requests for authentication */
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if err != nil {
					json.NewEncoder(w).Encode(types.Exception{Message: err.Error()})
					return
				}
				if token.Valid {
					log.Println("TOKEN WAS VALID")
					ctx := context.WithValue(r.Context(), "decoded", token.Claims)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					json.NewEncoder(w).Encode(types.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(types.Exception{Message: "An authorization header is required"})
		}
	})
}
