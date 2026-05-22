package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"backend/internal/auth/application/command"
	"backend/internal/auth/interfaces/http/dto"
	shareddomain "backend/internal/shared/domain"
	"backend/pkg/response"
)

type AuthHandler struct {
	register *command.RegisterUserHandler
	login    *command.LoginUserHandler
}

func NewAuthHandler(register *command.RegisterUserHandler, login *command.LoginUserHandler) *AuthHandler {
	return &AuthHandler{register: register, login: login}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	out, err := h.register.Handle(c.Request.Context(), command.RegisterUserInput{
		Email: req.Email, Password: req.Password, Name: req.Name,
	})
	if err != nil {
		if errors.Is(err, shareddomain.ErrAlreadyExists) {
			response.Error(c, 409, "EMAIL_TAKEN", "email already registered")
			return
		}
		response.BadRequest(c, err.Error())
		return
	}
	response.Created(c, out)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	out, err := h.login.Handle(c.Request.Context(), command.LoginUserInput{
		Email: req.Email, Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, shareddomain.ErrUnauthorized) {
			response.Unauthorized(c, "invalid email or password")
			return
		}
		response.Internal(c, err.Error())
		return
	}
	response.OK(c, out)
}
