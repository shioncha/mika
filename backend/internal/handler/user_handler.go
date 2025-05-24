package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Get(c *gin.Context) {
	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := h.authService.GetByUlid(c.Request.Context(), uidStr)
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
