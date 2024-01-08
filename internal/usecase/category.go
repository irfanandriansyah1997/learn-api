package usecase

import (
	"context"

	"github.com/irfanandriansyah1997/learn-api/internal/model"
)

type CategoryUsecase interface {
	Create(ctx context.Context, args model.CategoryCreateArgs) model.CategoryResponse
	Update(ctx context.Context, args model.CategoryUpdateArgs) model.CategoryResponse
	Delete(ctx context.Context, categoryID int)
	FindByID(ctx context.Context, categoryID int) model.CategoryResponse
	FindAll(ctx context.Context) []model.CategoryResponse
}
