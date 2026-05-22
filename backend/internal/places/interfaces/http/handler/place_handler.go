package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/internal/places/application/command"
	"backend/internal/places/application/query"
	"backend/internal/places/interfaces/http/dto"
	"backend/pkg/response"
)

type PlaceHandler struct {
	create *command.CreatePlaceHandler
	list   *query.ListPlacesHandler
}

func NewPlaceHandler(create *command.CreatePlaceHandler, list *query.ListPlacesHandler) *PlaceHandler {
	return &PlaceHandler{create: create, list: list}
}

func (h *PlaceHandler) Create(c *gin.Context) {
	var req dto.CreatePlaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	out, err := h.create.Handle(c.Request.Context(), command.CreatePlaceInput{
		CategoryID: req.CategoryID, Name: req.Name,
		Latitude: req.Latitude, Longitude: req.Longitude,
		Address: req.Address, Description: req.Description,
	})
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Created(c, out)
}

func (h *PlaceHandler) List(c *gin.Context) {
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "50"), 10, 64)
	skip, _ := strconv.ParseInt(c.DefaultQuery("skip", "0"), 10, 64)
	out, err := h.list.Handle(c.Request.Context(), query.ListPlacesInput{
		CategoryID: c.Query("category_id"),
		Search:     c.Query("search"),
		Limit:      limit,
		Skip:       skip,
	})
	if err != nil {
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, out)
}

func (h *PlaceHandler) Get(c *gin.Context) {
	response.OK(c, gin.H{"id": c.Param("id"), "note": "implement GetPlaceQuery similarly"})
}
