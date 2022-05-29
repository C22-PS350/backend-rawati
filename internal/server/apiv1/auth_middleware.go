package apiv1

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/C22-PS350/backend-rawati/internal/utils"
	"github.com/patrickmn/go-cache"
)

type ctxKey string

var key ctxKey = "userID"

var (
	authFindUserByToken = `
		SELECT user_id
		FROM user u
		JOIN user_token ut
		WHERE ut.token = ?
	`
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), " ")
		if len(authHeader) != 2 {
			utils.RespondErr(w, http.StatusBadRequest, errors.New("invalid authorization header form"))
			return
		}

		if authHeader[0] != "Bearer" {
			utils.RespondErr(w, http.StatusBadRequest, errors.New("authorization header must be bearer token"))
			return
		}

		d, ok := h.C.Get(authHeader[1])
		if ok {
			ctx := context.WithValue(r.Context(), key, *(d.(*uint32)))
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		var u struct {
			UserID uint32
		}
		if err := h.DB.Raw(authFindUserByToken, authHeader[1]).Scan(&u).Error; err != nil {
			utils.RespondErr(w, http.StatusInternalServerError, err)
			return
		}

		if u.UserID == 0 {
			utils.RespondErr(w, http.StatusUnauthorized, errors.New("invalid user token"))
			return
		}

		h.C.Set(authHeader[1], &u.UserID, cache.DefaultExpiration)
		ctx := context.WithValue(r.Context(), key, u.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
