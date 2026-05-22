package query

import (
	"context"

	"backend/internal/places/application/dto"
	"backend/internal/places/domain/repository"
)

type ListCategoriesHandler struct {
	categories repository.CategoryRepository
}

func NewListCategoriesHandler(categories repository.CategoryRepository) *ListCategoriesHandler {
	return &ListCategoriesHandler{categories: categories}
}

func (h *ListCategoriesHandler) Handle(ctx context.Context) ([]dto.CategoryResponse, error) {
	items, err := h.categories.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]dto.CategoryResponse, 0, len(items))
	for _, c := range items {
		out = append(out, dto.ToCategoryResponse(c))
	}
	return out, nil
}
