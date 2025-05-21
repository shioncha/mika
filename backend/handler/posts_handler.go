package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/users"
	"github.com/shioncha/mika/backend/internal/service"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

type GetPostResponse struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := h.postService.GetPosts(c.Request.Context(), uidStr)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	err := h.postService.CreatePost(c.Request.Context(), uidStr, req.Content)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")

	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	err := h.postService.DeletePost(c.Request.Context(), uidStr, id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func getUserIdByUlid(ctx context.Context, client *ent.Client, uidStr string) (int, error) {
	user, err := client.Users.Query().Where(users.UlidEQ(uidStr)).Select(users.FieldID).First(ctx)
	if err != nil && ent.IsNotFound(err) {
		return 0, fmt.Errorf("unauthorized")
	}
	if err != nil {
		return 0, fmt.Errorf("Internal server error")
	}
	return user.ID, nil
}
