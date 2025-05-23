package handler

import (
	"net/http"

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

func (h *TagHandler) GetTags(c *gin.Context) {
	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := h.tagService.GetTags(c.Request.Context(), uidStr)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *TagHandler) GetPostsByTag(c *gin.Context) {
	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	tag := c.Param("tag")
	if tag == "" {
		respondWithError(c, http.StatusBadRequest, "Tag is required")
		return
	}

	res, err := h.tagService.GetPostsByTag(c.Request.Context(), uidStr, tag)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, res)
}
