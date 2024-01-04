// Package repository - category repo
package repository

import (
	"context"
	"database/sql"

	"github.com/irfanandriansyah1997/learn-api/internal/model"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Delete(ctx context.Context, tx *sql.Tx, category model.Category)
	FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (model.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.Category
}
