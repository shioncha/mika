package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/service"
)

type TagHandler struct {
	tagService *service.TagService
}

func NewTagHandler(tagService *service.TagService) *TagHandler {
	return &TagHandler{
		tagService: tagService,
	}
}

// @Summary			Get User's Tags
// @Description	Get all tags associated with the authenticated user
// @Tags				Tags
// @Produce			json
// @Success			200  {array} string
// @Failure			500  {object} ErrorResponse
// @Router			/tags [get]
func (h *TagHandler) GetTags(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := h.tagService.GetTags(c.Request.Context(), userID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary			Get Posts having the Tag
// @Description	Get posts associated with a specific tag
// @Tags				Tags
// @Produce			json
// @Success			200  {object} GetPostsResponse
// @Failure			400  {object} ErrorResponse
// @Failure			500  {object} ErrorResponse
// @Router			/tags/:tag/posts [get]
func (h *TagHandler) GetPostsByTag(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	tag := c.Param("tag")
	if tag == "" {
		respondWithError(c, http.StatusBadRequest, "Tag is required")
		return
	}

	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 20
	}

	cursor := c.Query("cursor")

	res, nextCursor, err := h.tagService.GetPostsByTag(c.Request.Context(), userID, tag, limit, cursor)
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
