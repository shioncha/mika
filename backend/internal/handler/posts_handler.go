package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (h *PostHandler) GetPosts(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}
	limitStr := c.DefaultQuery("limit", "20")
	cursor := c.Query("cursor")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 20
	}

	res, nextCursor, err := h.postService.GetPosts(c.Request.Context(), userID, limit, cursor)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	var response = gin.H{
		"posts":       res,
		"next_cursor": nextCursor,
	}

	c.JSON(http.StatusOK, response)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	postID := c.Param("id")

	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := h.postService.GetPost(c.Request.Context(), userID, postID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, res)
}

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	err := h.postService.CreatePost(c.Request.Context(), userID, req.Content)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postID := c.Param("id")

	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	err := h.postService.DeletePost(c.Request.Context(), userID, postID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}
