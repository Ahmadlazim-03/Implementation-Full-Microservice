package http

import (
	"github.com/gin-gonic/gin"

	"backend/internal/places/interfaces/http/handler"
)

func RegisterRoutes(r *gin.Engine, place *handler.PlaceHandler, category *handler.CategoryHandler) {
	api := r.Group("/api")
	{
		api.GET("/places", place.List)
		api.GET("/places/nearby", place.Nearby) // declared before :id so it doesn't shadow
		api.POST("/places", place.Create)
		api.GET("/places/:id", place.Get)
		api.GET("/categories", category.List)
	}
}
