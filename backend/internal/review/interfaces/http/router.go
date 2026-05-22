package http

import (
	"github.com/gin-gonic/gin"

	"backend/internal/review/interfaces/http/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.ReviewHandler) {
	api := r.Group("/api/reviews")
	{
		api.POST("", h.Create)
		api.GET("/place/:placeID", h.ListByPlace)
	}
}
