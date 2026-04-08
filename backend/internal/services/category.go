package services

import (
	"context"

	"github.com/amesupakorn/pocket-ledger/internal/dto"
	"github.com/amesupakorn/pocket-ledger/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		repo: repository.NewCategoryRepository(),
	}
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	return s.repo.GetAll(ctx)
}
