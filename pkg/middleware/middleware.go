package middleware

import (
	"net/http"

	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	"auth/auth_back/pkg/helpers/tokenHelper"
	"auth/auth_back/pkg/logger"
)

var l = logger.Logger{}

type UserMiddlewareResponse struct {
	StatusENUM string `json:"statusEnum"`
	Message    string `json:"message"`
}

func UserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenSrt := tokenHelper.ExtractToken(r)

		if tokenSrt == "" {
			httpServerHelper.ReturnErrUnauthorized(w, globalvars.ErrTokenAccess.Error, globalvars.ErrTokenAccess.Enum)
			return
		}

		token, err := tokenHelper.VerifyToken(tokenSrt)

		if err != nil {
			httpServerHelper.ReturnErrUnauthorized(w, globalvars.ErrTokenAccess.Error, globalvars.ErrTokenAccess.Enum)
			return
		}

		tokenData, err := tokenHelper.ExtractTokenMetadataUser(token)

		if err != nil {
			httpServerHelper.ReturnErrUnauthorized(w, globalvars.ErrTokenAccess.Error, globalvars.ErrTokenAccess.Enum)
			return
		}

		r.Header.Add("userId", tokenData.UserId.String())

		next(w, r)
	}
}

func AdminUserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenSrt := tokenHelper.ExtractToken(r)

		if tokenSrt == "" {
			httpServerHelper.ReturnErrUnauthorized(w, globalvars.ErrTokenAccess.Error, globalvars.ErrTokenAccess.Enum)
			return
		}

		token, err := tokenHelper.VerifyToken(tokenSrt)

		if err != nil {
			httpServerHelper.ReturnErrUnauthorized(w, globalvars.ErrTokenAccess.Error, globalvars.ErrTokenAccess.Enum)
			return
		}

		tokenData, err := tokenHelper.ExtractTokenMetadataAdmin(token)

		if err != nil {
			httpServerHelper.ReturnErrUnauthorized(w, globalvars.ErrTokenAccess.Error, globalvars.ErrTokenAccess.Enum)
			return
		}

		r.Header.Add("adminId", tokenData.AdminId.String())

		next(w, r)
	}
}
