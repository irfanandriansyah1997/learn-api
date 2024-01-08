//nolint:typecheck // disable lint since WriteToResponseBody still on same package
package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/irfanandriansyah1997/learn-api/internal/model"
)

func ErrorHandlerHTTP(writer http.ResponseWriter, request *http.Request, err any) {
	if validationError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := model.APIResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		WriteToResponseBody(writer, response)
		return true
	}

	return false
}

func validationError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := model.APIResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		WriteToResponseBody(writer, response)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := model.APIResponse{
		Code:   http.StatusInternalServerError,
		Status: "Failure",
		Data:   err,
	}

	WriteToResponseBody(writer, response)
}
