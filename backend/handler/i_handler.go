package handler

import "github.com/gin-gonic/gin"

func I(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	c.JSON(200, gin.H{"user_id": userID})
}
