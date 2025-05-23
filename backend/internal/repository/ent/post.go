package entrepogitory

import (
	"context"
	"fmt"

	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/posts"
	"github.com/shioncha/mika/backend/ent/tags"
	"github.com/shioncha/mika/backend/ent/users"
	"github.com/shioncha/mika/backend/internal/repository"
	"github.com/shioncha/mika/backend/internal/utils"
)

type PostRepository struct {
	client *ent.Client
}

func NewPostRepository(client *ent.Client) *PostRepository {
	return &PostRepository{
		client: client,
	}
}

func (r *PostRepository) GetPostsByUserID(ctx context.Context, userID int) ([]*repository.Post, error) {
	postList, err := r.client.Posts.Query().Where(posts.UserIDEQ(userID)).All(ctx)
	if err != nil {
		return nil, err
	}

	var posts []*repository.Post
	for _, post := range postList {
		posts = append(posts, &repository.Post{
			ID:        post.Ulid,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.String(),
			UpdatedAt: post.UpdatedAt.String(),
		})
	}

	return posts, nil
}

func (r *PostRepository) GetPostByPostID(ctx context.Context, userID int, postID string) (*repository.Post, error) {
	post, err := r.client.Posts.Query().
		Where(posts.UserIDEQ(userID), posts.UlidEQ(postID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &repository.Post{
		ID:        post.Ulid,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.String(),
		UpdatedAt: post.UpdatedAt.String(),
	}, nil
}

func (r *PostRepository) CreatePost(ctx context.Context, tx *ent.Tx, userID int, content string, tags []int) error {
	id := utils.GenerateULID()
	_, err := tx.Posts.Create().
		SetUlid(id).
		SetUserID(userID).
		SetContent(content).
		AddTagIDs(tags...).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) DeletePost(ctx context.Context, userID int, postID string) error {
	_, err := r.client.Posts.Delete().
		Where(posts.UserIDEQ(userID), posts.UlidEQ(postID)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetUserIDByUlid(ctx context.Context, ulid string) (int, error) {
	user, err := r.client.Users.Query().
		Where(users.UlidEQ(ulid)).
		First(ctx)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *PostRepository) CreateTags(ctx context.Context, tx *ent.Tx, userID int, tagList []string) ([]int, error) {
	var tagIDs []int

	for _, tag := range tagList {
		existingTag, err := tx.Tags.Query().
			Where(tags.UserIDEQ(userID), tags.TagEQ(tag)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, fmt.Errorf("failed to check tag existence")
		}
		if existingTag != nil {
			tagIDs = append(tagIDs, existingTag.ID)
			continue
		}

		tagUlid := utils.GenerateULID()
		res, err := tx.Tags.Create().
			SetTag(tag).
			SetUserID(userID).
			SetUlid(tagUlid).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create tag")
		}
		tagIDs = append(tagIDs, res.ID)
	}

	return tagIDs, nil
}
