package handler

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/posts"
	"github.com/shioncha/mika/backend/ent/tags"
	"github.com/shioncha/mika/backend/ent/users"
	"github.com/shioncha/mika/backend/internal/utils"
)

type GetPostResponse struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

func GetPost(c *gin.Context, client *ent.Client) {
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

	postList, err := client.Posts.Query().Where(posts.UserIDEQ(userID)).All(ctx)
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

func CreatePost(c *gin.Context, client *ent.Client) {
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
	ctx := c.Request.Context()
	id := utils.GenerateULID()

	userID, err := getUserIdByUlid(ctx, client, uidStr)
	if err != nil && err.Error() == "unauthorized" {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	tagList, err := getTags(req.Content)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to start transaction")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	var tagIDs []int

	for _, tag := range tagList {
		existingTag, err := tx.Tags.Query().Where(tags.UserIDEQ(userID), tags.TagEQ(tag)).First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			tx.Rollback()
			respondWithError(c, http.StatusInternalServerError, "Failed to check tag existence")
			return
		}
		if existingTag != nil {
			tagIDs = append(tagIDs, existingTag.ID)
			continue
		}

		tagUlid := utils.GenerateULID()
		res, err := tx.Tags.Create().SetTag(tag).SetUserID(userID).SetUlid(tagUlid).Save(ctx)
		if err != nil {
			tx.Rollback()
			respondWithError(c, http.StatusInternalServerError, "Failed to create tag")
			return
		}
		tagIDs = append(tagIDs, res.ID)
	}

	_, err = tx.Posts.
		Create().
		SetUlid(id).
		SetUserID(userID).
		SetContent(req.Content).
		AddTagIDs(tagIDs...).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		respondWithError(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func DeletePost(c *gin.Context, client *ent.Client) {
	id := c.Param("id")

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

	_, err = client.Posts.Delete().Where(posts.UserIDEQ(userID), posts.UlidEQ(id)).Exec(ctx)
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

func getTags(content string) ([]string, error) {
	r, err := regexp.Compile(`#\S+`)
	if err != nil {
		return nil, err
	}
	matches := r.FindAllString(content, -1)
	var tags []string
	for _, match := range matches {
		tags = append(tags, match[1:])
	}
	return tags, nil
}
