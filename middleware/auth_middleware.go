package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Omaroma/drone-backend/services"
)

type ctxKey string

const ClaimsKey ctxKey = "claims"

func Auth(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := services.ValidateToken(token)
		if err != nil || claims.Role != requiredRole {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next(w, r.WithContext(ctx))
	}
}
