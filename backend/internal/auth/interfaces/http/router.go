package http

import (
	"github.com/gin-gonic/gin"

	"backend/internal/auth/interfaces/http/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.AuthHandler) {
	api := r.Group("/api/auth")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
	}
}
