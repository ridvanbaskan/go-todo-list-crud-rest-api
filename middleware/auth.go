package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"todo-list/auth"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode("Missing authentication token")
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.ParseWithClaims(tokenString, &auth.Handler{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode("Invalid auth token")
			return
		}

		fmt.Println(token.Claims.(*auth.Handler))

		ctx := context.WithValue(request.Context(), "userID", token.Claims.(*auth.Handler).ID)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
