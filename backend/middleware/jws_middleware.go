package middleware

import (
	"net/http"
	"product_manage/service" // Import service để dùng ValidateJWT
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Lấy token từ header: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		token := parts[1]

		// Validate JWT Token
		if !service.ValidateJWT(token) {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Nếu hợp lệ, chuyển tiếp request tới handler tiếp theo
		next.ServeHTTP(w, r)
	})
}
