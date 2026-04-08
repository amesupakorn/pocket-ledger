package repository

import (
	"context"

	"github.com/amesupakorn/pocket-ledger/internal/database"
	"github.com/amesupakorn/pocket-ledger/internal/dto"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]dto.CategoryResponse, error) {
	rows, err := database.DB.Query(ctx, `
		SELECT id, name, type, created_at
		FROM categories
		ORDER BY id
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []dto.CategoryResponse

	for rows.Next() {
		var c dto.CategoryResponse
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Type,
			&c.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}
