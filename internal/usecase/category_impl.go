//nolint:typecheck // disable lint since CategoryUsecase still on same package
package usecase

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/irfanandriansyah1997/learn-api/internal/model"
	"github.com/irfanandriansyah1997/learn-api/internal/repository"
	"github.com/irfanandriansyah1997/learn-api/internal/utils"
)

type CategoryUsecaseImpl struct {
	repo     repository.CategoryRepository
	DB       *sql.DB
	validate *validator.Validate
}

func NewCategoryUsecase(repo repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryUsecase {
	return &CategoryUsecaseImpl{
		repo:     repo,
		DB:       db,
		validate: validate,
	}
}

func adapterCategoryResponse(category model.Category) model.CategoryResponse {
	return model.CategoryResponse(category)
}

func adapterCategoryResponses(categories []model.Category) []model.CategoryResponse {
	categoryResponses := []model.CategoryResponse{}

	for _, category := range categories {
		categoryResponses = append(categoryResponses, adapterCategoryResponse(category))
	}

	return categoryResponses
}

func (usecase *CategoryUsecaseImpl) Create(ctx context.Context, args model.CategoryCreateArgs) model.CategoryResponse {
	err := usecase.validate.Struct(args)
	utils.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	utils.PanicIfError(err)

	defer utils.CommitOrRollback(tx)

	result := usecase.repo.Save(ctx, tx, model.Category{
		Name: args.Name,
	})

	return adapterCategoryResponse(result)

}

func (usecase *CategoryUsecaseImpl) Update(ctx context.Context, args model.CategoryUpdateArgs) model.CategoryResponse {
	err := usecase.validate.Struct(args)
	utils.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	utils.PanicIfError(err)

	defer utils.CommitOrRollback(tx)

	category, err := usecase.repo.FindByID(ctx, tx, args.ID)
	utils.PanicIfNotFoundError(err)

	category.Name = args.Name
	result := usecase.repo.Update(ctx, tx, category)

	return adapterCategoryResponse(result)
}

func (usecase *CategoryUsecaseImpl) Delete(ctx context.Context, categoryID int) {
	tx, err := usecase.DB.Begin()
	utils.PanicIfError(err)

	defer utils.CommitOrRollback(tx)

	category, err := usecase.repo.FindByID(ctx, tx, categoryID)
	utils.PanicIfNotFoundError(err)

	usecase.repo.Delete(ctx, tx, category)
}

func (usecase *CategoryUsecaseImpl) FindByID(ctx context.Context, categoryID int) model.CategoryResponse {
	tx, err := usecase.DB.Begin()
	utils.PanicIfError(err)

	category, err := usecase.repo.FindByID(ctx, tx, categoryID)
	utils.PanicIfNotFoundError(err)

	return adapterCategoryResponse(category)
}

func (usecase *CategoryUsecaseImpl) FindAll(ctx context.Context) []model.CategoryResponse {
	tx, err := usecase.DB.Begin()
	utils.PanicIfError(err)

	defer utils.CommitOrRollback(tx)

	categories := usecase.repo.FindAll(ctx, tx)

	return adapterCategoryResponses(categories)
}
