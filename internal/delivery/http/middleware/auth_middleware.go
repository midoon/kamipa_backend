package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/midoon/kamipa_backend/internal/util"
)

func AuthMiddleware(tokenUtil *util.TokenUtil, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.WriteJSON(w, http.StatusUnauthorized, model.MessageResponse{
				Status:  false,
				Message: "Auhorization header is required",
			})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			helper.WriteJSON(w, http.StatusUnauthorized, model.MessageResponse{
				Status:  false,
				Message: "Invalid Authorization header format",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := tokenUtil.ParseToken(r.Context(), token)
		if err != nil {
			customErr := err.(*helper.CustomError)

			helper.WriteJSON(w, customErr.Code, model.MessageResponse{
				Status:  false,
				Message: customErr.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), helper.UserIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
