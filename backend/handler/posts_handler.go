package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/posts"
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

	user, err := getUserByUlid(c, client, uidStr)
	if err != nil && ent.IsNotFound(err) {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	postList, err := client.Posts.Query().Where(posts.UserIDEQ(user.ID)).All(ctx)
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

	user, err := getUserByUlid(c, client, uidStr)
	if err != nil && err.Error() == "unauthorized" {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	_, err = client.Posts.
		Create().
		SetUlid(id).
		SetUserID(user.ID).
		SetContent(req.Content).
		Save(ctx)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create post")
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

	user, err := getUserByUlid(c, client, uidStr)
	if err != nil && ent.IsNotFound(err) {
		respondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	res, err := client.Posts.Delete().Where(posts.UserIDEQ(user.ID), posts.UlidEQ(id)).Exec(ctx)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}

func getUserByUlid(ctx context.Context, client *ent.Client, uidStr string) (*ent.Users, error) {
	user, err := client.Users.Query().Where(users.UlidEQ(uidStr)).Select(users.FieldID).First(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, fmt.Errorf("unauthorized")
	}
	if err != nil {
		return nil, fmt.Errorf("Internal server error")
	}
	return user, nil
}
