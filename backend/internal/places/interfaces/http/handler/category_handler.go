package handler

import (
	"github.com/gin-gonic/gin"

	"backend/internal/places/application/query"
	"backend/pkg/response"
)

type CategoryHandler struct {
	list *query.ListCategoriesHandler
}

func NewCategoryHandler(list *query.ListCategoriesHandler) *CategoryHandler {
	return &CategoryHandler{list: list}
}

func (h *CategoryHandler) List(c *gin.Context) {
	out, err := h.list.Handle(c.Request.Context())
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, out)
}
