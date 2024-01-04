package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/irfanandriansyah1997/learn-api/internal/model"
	"github.com/irfanandriansyah1997/learn-api/internal/utils"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	SQL := "INSERT INTO customer(name) VALUE (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	utils.PanicIfError(err)

	id, err := result.LastInsertId()
	utils.PanicIfError(err)

	category.ID = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)
	utils.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category model.Category) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.ID)
	utils.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (model.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryID)
	utils.PanicIfError(err)

	category := model.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		utils.PanicIfError(err)

		return category, nil
	}

	return category, errors.New("Category is not found")

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	utils.PanicIfError(err)

	var categories []model.Category
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		utils.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
