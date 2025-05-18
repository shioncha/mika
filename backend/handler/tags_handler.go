package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/posts"
	"github.com/shioncha/mika/backend/ent/tags"
)

func GetTags(c *gin.Context, client *ent.Client) {
	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}
	ctx := c.Request.Context()

	userID, err := getUserIdByUlid(ctx, client, uidStr)
	if err != nil && err.Error() == "unauthorized" {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	tagList, err := client.Tags.Query().Where(tags.UserIDEQ(userID)).All(ctx)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	var tagsResponse []string
	for _, tag := range tagList {
		tagsResponse = append(tagsResponse, tag.Tag)
	}
	c.JSON(http.StatusOK, tagsResponse)
}

func GetPostsByTag(c *gin.Context, client *ent.Client) {
	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}
	ctx := c.Request.Context()

	userID, err := getUserIdByUlid(ctx, client, uidStr)
	if err != nil && err.Error() == "unauthorized" {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	tag := c.Param("tag")
	if tag == "" {
		respondWithError(c, http.StatusBadRequest, "Tag is required")
		return
	}

	postList, err := client.Posts.Query().Where(posts.UserIDEQ(userID), posts.HasTagsWith(tags.Tag(tag))).All(ctx)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	var postsResponse []GetPostResponse
	for _, post := range postList {
		postsResponse = append(postsResponse, GetPostResponse{
			Id:        post.Ulid,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.String(),
			UpdatedAt: post.UpdatedAt.String(),
		})
	}
	c.JSON(http.StatusOK, postsResponse)
}
