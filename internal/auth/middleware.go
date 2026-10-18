package auth

import (
	"net/http"
	"strings"
)

func RequireAuth(tm *TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := r.Header.Get("Authorization")
			if h == "" || !strings.HasPrefix(h, "Bearer ") {
				http.Error(w, "Unauthorized", 401)
				return
			}
			raw := strings.TrimPrefix(h, "Bearer ")
			userID, err := tm.Parse(raw)
			if err != nil {
				http.Error(w, "Unauthorized", 401)
				return
			}
			ctx := ContextWithUserID(r.Context(), userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
