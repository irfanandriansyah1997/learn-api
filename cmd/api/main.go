package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/irfanandriansyah1997/learn-api/internal/controller"
	"github.com/irfanandriansyah1997/learn-api/internal/repository"
	"github.com/irfanandriansyah1997/learn-api/internal/usecase"
	"github.com/irfanandriansyah1997/learn-api/internal/utils"
	"github.com/irfanandriansyah1997/learn-api/pkg/database"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := database.NewDatabase()
	categoryRepository := repository.NewCatetegoryRepository()
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryUsecase)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindByID)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = utils.ErrorHandlerHTTP

	server := http.Server{
		Addr:              "localhost:3000",
		Handler:           utils.NewAuthMiddleware(router),
		ReadHeaderTimeout: 2 * time.Second,
	}

	err := server.ListenAndServe()
	utils.PanicIfError(err)
}
