package auth

import (
	"context"
	"net/http"
)

func Required() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			token := request.Header.Get("token")

			if token != "admin" && token != "user" {
				writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(request.Context(), "token", token)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}
