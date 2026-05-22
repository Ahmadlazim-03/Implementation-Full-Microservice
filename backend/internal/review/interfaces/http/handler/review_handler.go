package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"backend/internal/review/application/command"
	"backend/internal/review/application/query"
	"backend/internal/review/interfaces/http/dto"
	shareddomain "backend/internal/shared/domain"
	"backend/pkg/response"
)

type ReviewHandler struct {
	create      *command.CreateReviewHandler
	listByPlace *query.ListReviewsByPlaceHandler
}

func NewReviewHandler(create *command.CreateReviewHandler, listByPlace *query.ListReviewsByPlaceHandler) *ReviewHandler {
	return &ReviewHandler{create: create, listByPlace: listByPlace}
}

func (h *ReviewHandler) Create(c *gin.Context) {
	var req dto.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	out, err := h.create.Handle(c.Request.Context(), command.CreateReviewInput{
		PlaceID: req.PlaceID, UserID: req.UserID,
		Rating: req.Rating, Comment: req.Comment,
	})
	if err != nil {
		if errors.Is(err, shareddomain.ErrNotFound) {
			response.NotFound(c, "place_id not found")
			return
		}
		response.BadRequest(c, err.Error())
		return
	}
	response.Created(c, out)
}

func (h *ReviewHandler) ListByPlace(c *gin.Context) {
	out, err := h.listByPlace.Handle(c.Request.Context(), c.Param("placeID"))
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, out)
}
