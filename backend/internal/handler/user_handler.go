package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/service"
)

const (
	ContextKeyUserID = "user_id"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Name     string `json:"name" binding:"omitempty"`
	Password string `json:"password" binding:"omitempty,min=8"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func getUserIDFromContext(c *gin.Context) (string, bool) {
	uid, exists := c.Get(ContextKeyUserID)
	if !exists {
		return "", false
	}

	userID, ok := uid.(string)
	return userID, ok
}

// @Summary			Get User Info
// @Description	Get information about the authenticated user
// @Tags				User
// @Produce			json
// @Success			200  {object} UserResponse
// @Failure			500  {object} ErrorResponse
// @Router			/user [get]
func (h *UserHandler) Get(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := h.userService.GetByID(c.Request.Context(), userID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(200, UserResponse{
		ID:    res.ID,
		Email: res.Email,
		Name:  res.Name,
	})
}

// @Summary			Update User Info
// @Description	Update the authenticated user's information (email, name, password)
// @Tags				User
// @Accept			json
// @Produce			json
// @Param     	request body UpdateUserRequest true "Fields to update"
// @Success			200  {object} map[string]string
// @Failure			400  {object} ErrorResponse
// @Failure			500  {object} ErrorResponse
// @Router			/user [patch]
func (h *UserHandler) Update(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	if req.Name == "" && req.Email == "" && req.Password == "" {
		respondWithError(c, http.StatusBadRequest, "No fields to update")
		return
	}

	if req.Name != "" {
		if err := h.userService.UpdateUsername(c.Request.Context(), userID, req.Name); err != nil {
			respondWithError(c, http.StatusInternalServerError, "Failed to update username")
			return
		}
	}

	if req.Email != "" {
		if err := h.userService.UpdateEmail(c.Request.Context(), userID, req.Email); err != nil {
			respondWithError(c, http.StatusInternalServerError, "Failed to update email")
			return
		}
	}

	if req.Password != "" {
		if err := h.userService.UpdatePassword(c.Request.Context(), userID, req.Password); err != nil {
			respondWithError(c, http.StatusInternalServerError, "Failed to update password")
			return
		}
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}
