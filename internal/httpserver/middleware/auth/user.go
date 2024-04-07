package auth

import "net/http"

func UserOnly() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			token := request.Context().Value("token")

			if token != "user" {
				writer.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(writer, request)
		})
	}
}
