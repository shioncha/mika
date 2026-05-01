package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/service"
)

const (
	CookieRefreshToken = "refresh_token"
	CookiePath         = "/api"
	RefreshTokenMaxAge = 7 * 24 * time.Hour
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

type ErrorResponse struct {
	Error string `json:"error"`
}

func setRefreshCookie(c *gin.Context, token string, maxAge int) {
	c.SetCookie(CookieRefreshToken, token, maxAge, CookiePath, os.Getenv("DOMAIN"), false, true)
}

func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, ErrorResponse{
		Error: message,
	})
}

// @Summary			Sign Up
// @Description	Create a new user
// @Tags				Auth
// @Accept			json
// @Produce			json
// @Param     	request body SignUpRequest true "User info"
// @Success			200  {object} AuthResponse
// @Failure			400  {object} ErrorResponse
// @Failure			409  {object} ErrorResponse
// @Failure			500  {object} ErrorResponse
// @Router			/sign-up [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request") // TODO: Add validation error
		return
	}

	res, err := h.authService.SignUp(c.Request.Context(), service.SignUpParams{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Device:   c.GetHeader("User-Agent"),
		IP:       c.ClientIP(),
	})
	if err != nil {
		if err.Error() == "email already registered" {
			respondWithError(c, http.StatusConflict, "Email already registered")
			return
		}
		respondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	setRefreshCookie(c, res.RefreshToken, int(RefreshTokenMaxAge.Seconds()))

	c.JSON(http.StatusOK, AuthResponse{
		Token: res.Token,
	})
}

// @Summary			Sign In
// @Description	Sign in to the application
// @Tags				Auth
// @Accept			json
// @Produce			json
// @Param     	request body SignInRequest true "User info"
// @Success			200  {object} AuthResponse
// @Failure			400  {object} ErrorResponse
// @Failure			401  {object} ErrorResponse
// @Failure			500  {object} ErrorResponse
// @Router			/sign-in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	res, err := h.authService.SignIn(c.Request.Context(), service.SignInParams{
		Email:    req.Email,
		Password: req.Password,
		Device:   c.GetHeader("User-Agent"),
		IP:       c.ClientIP(),
	})
	if err != nil {
		if err.Error() == "invalid credentials" {
			respondWithError(c, http.StatusUnauthorized, "Invalid email or password")
			return
		}
		respondWithError(c, http.StatusInternalServerError, "Failed to sign in")
		return
	}

	setRefreshCookie(c, res.RefreshToken, int(RefreshTokenMaxAge.Seconds()))

	c.JSON(http.StatusOK, AuthResponse{
		Token: res.Token,
	})
}

// @Summary			Refresh Access Token
// @Description	Refresh the access token
// @Tags				Auth
// @Produce			json
// @Success			200  {object} AuthResponse
// @Failure			401  {object} ErrorResponse
// @Router			/refresh-token [post]
func (h *AuthHandler) RefreshAccessToken(c *gin.Context) {
	oldRefreshToken, err := c.Cookie(CookieRefreshToken)
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	res, err := h.authService.RefreshAccessToken(c.Request.Context(), oldRefreshToken)
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Invalid session")
		return
	}
	c.JSON(http.StatusOK, AuthResponse{
		Token: res.Token,
	})
}

// @Summary			Sign Out
// @Description	Sign out the current user
// @Tags				Auth
// @Produce			json
// @Success			204
// @Failure			401  {object} ErrorResponse
// @Failure			500  {object} ErrorResponse
// @Router			/sign-out [post]
func (h *AuthHandler) SignOut(c *gin.Context) {
	refreshToken, err := c.Cookie(CookieRefreshToken)
	if err != nil {
		respondWithError(c, http.StatusUnauthorized, "Refresh token not found")
		return
	}

	if err := h.authService.SignOut(c.Request.Context(), refreshToken); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to sign out")
		return
	}

	setRefreshCookie(c, "", -1)
	c.Status(http.StatusNoContent)
}

func (h *AuthHandler) GetAllSessions(c *gin.Context) {
	userID := c.GetString("user_id")

	sessions, err := h.authService.GetAllSessions(c.Request.Context(), userID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to retrieve sessions")
		return
	}

	c.JSON(http.StatusOK, sessions)
}

func (h *AuthHandler) RevokeAllSessions(c *gin.Context) {
	userID := c.GetString("user_id")

	if err := h.authService.RevokeAllSessions(c.Request.Context(), userID); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to revoke sessions")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All sessions revoked successfully"})
}
