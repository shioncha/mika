package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type SignUpRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	res, err := h.authService.SignUp(c.Request.Context(), service.SignUpParams{
		Email:           req.Email,
		Name:            req.Name,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil && err.Error() == "email already registered" {
		respondWithError(c, http.StatusConflict, "Email already registered")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	res, err := h.authService.SignIn(c.Request.Context(), service.SignInParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil && err.Error() == "invalid credentials" {
		respondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to sign in")
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
	})
}

func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
