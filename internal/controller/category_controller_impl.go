//nolint:typecheck // disable lint since CategoryController still on same package
package controller

import (
	"net/http"
	"strconv"

	"github.com/irfanandriansyah1997/learn-api/internal/model"
	"github.com/irfanandriansyah1997/learn-api/internal/usecase"
	"github.com/irfanandriansyah1997/learn-api/internal/utils"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	service usecase.CategoryUsecase
}

func NewCategoryController(service usecase.CategoryUsecase) CategoryController {
	return &CategoryControllerImpl{
		service: service,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	args := model.CategoryCreateArgs{}
	utils.ReadFromRequestBody(request, &args)

	response := controller.service.Create(request.Context(), args)
	apiResponse := model.APIResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	utils.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	args := model.CategoryUpdateArgs{}
	utils.ReadFromRequestBody(request, &args)

	categoryID, err := strconv.Atoi(params.ByName("categoryId"))
	utils.PanicIfError(err)
	args.ID = categoryID

	response := controller.service.Update(request.Context(), args)
	apiResponse := model.APIResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	utils.WriteToResponseBody(writer, apiResponse)

}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID, err := strconv.Atoi(params.ByName("categoryId"))
	utils.PanicIfError(err)

	controller.service.Delete(request.Context(), categoryID)
	apiResponse := model.APIResponse{
		Code:   200,
		Status: "Success",
	}

	utils.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID, err := strconv.Atoi(params.ByName("categoryId"))
	utils.PanicIfError(err)

	response := controller.service.FindByID(request.Context(), categoryID)
	apiResponse := model.APIResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	utils.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response := controller.service.FindAll(request.Context())
	apiResponse := model.APIResponse{
		Code:   200,
		Status: "Success",
		Data:   response,
	}

	utils.WriteToResponseBody(writer, apiResponse)
}
