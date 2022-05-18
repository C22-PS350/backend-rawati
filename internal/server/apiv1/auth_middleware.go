package apiv1

import (
	"context"
	"net/http"
)

type ctxKey string

var key ctxKey = "key"

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), key, "value")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
