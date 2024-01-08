//nolint:typecheck // disable lint since WriteToResponseBody still on same package
package utils

import (
	"net/http"

	"github.com/irfanandriansyah1997/learn-api/internal/model"
)

type AuthMiddleWare struct {
	handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleWare {
	return &AuthMiddleWare{handler: handler}
}

func (middleware *AuthMiddleWare) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "rahasia" {
		middleware.handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		response := model.APIResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		WriteToResponseBody(writer, response)
	}
}
