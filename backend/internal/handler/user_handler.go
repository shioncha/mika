package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Get(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := h.userService.GetByID(c.Request.Context(), userID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(200, gin.H{
		"id":    res.ID,
		"email": res.Email,
		"name":  res.Name,
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	var req struct {
		Email    string `json:"email,omitempty"`
		Name     string `json:"name,omitempty"`
		Password string `json:"password,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := h.userService.UpdateUsername(c.Request.Context(), userID, req.Name); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to update username")
		return
	}

	if err := h.userService.UpdateEmail(c.Request.Context(), userID, req.Email); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to update email")
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}
