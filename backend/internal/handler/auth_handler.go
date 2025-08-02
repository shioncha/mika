package handler

import (
	"net/http"
	"os"
	"time"

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
	Token string `json:"token"`
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	deviceInfo := c.GetHeader("User-Agent")
	ipAddress := c.ClientIP()

	res, err := h.authService.SignUp(c.Request.Context(), service.SignUpParams{
		Email:           req.Email,
		Name:            req.Name,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	}, deviceInfo, ipAddress)
	if err != nil && err.Error() == "email already registered" {
		respondWithError(c, http.StatusConflict, "Email already registered")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	c.SetCookie("refresh_token", res.RefreshToken, int((7 * 24 * time.Hour).Seconds()), "/api", "", false, true)

	c.JSON(http.StatusOK, AuthResponse{
		Token: res.Token,
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	deviceInfo := c.GetHeader("User-Agent")
	ipAddress := c.ClientIP()

	res, err := h.authService.SignIn(c, service.SignInParams{
		Email:    req.Email,
		Password: req.Password,
	}, deviceInfo, ipAddress)
	if err != nil && err.Error() == "invalid credentials" {
		respondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to sign in")
		return
	}

	c.SetCookie("refresh_token", res.RefreshToken, int((7 * 24 * time.Hour).Seconds()), "/api", "", false, true)

	c.JSON(http.StatusOK, AuthResponse{
		Token: res.Token,
	})
}

func (h *AuthHandler) RefreshAccessToken(c *gin.Context) {
	oldRefreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	res, err := h.authService.RefreshAccessToken(c, oldRefreshToken)
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Invalid session")
		return
	}
	c.JSON(http.StatusOK, AuthResponse{
		Token: res.Token,
	})
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	if err := h.authService.SignOut(c.Request.Context(), refreshToken); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to sign out")
		return
	}

	c.SetCookie("refresh_token", "", -1, "/api", os.Getenv("DOMAIN"), false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Signed out successfully"})
}

func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
