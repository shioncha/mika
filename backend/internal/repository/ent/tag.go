package entrepogitory

import (
	"context"

	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/posts"
	"github.com/shioncha/mika/backend/ent/tags"
	"github.com/shioncha/mika/backend/internal/repository"
)

type TagRepository struct {
	client *ent.Client
}

func NewTagRepository(client *ent.Client) *TagRepository {
	return &TagRepository{
		client: client,
	}
}

func (r *TagRepository) GetTags(ctx context.Context, userID string) ([]*repository.Tag, error) {
	tags, err := r.client.Tags.Query().
		Where(tags.UserID(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var tagList []*repository.Tag
	for _, tag := range tags {
		tagList = append(tagList, &repository.Tag{
			ID:   tag.ID,
			Name: tag.Tag,
		})
	}
	return tagList, nil
}

func (r *TagRepository) GetPostsByTag(ctx context.Context, userID string, tag string) ([]*repository.Post, error) {
	posts, err := r.client.Posts.Query().
		Where(posts.UserIDEQ(userID), posts.HasTagsWith(tags.Tag(tag))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var postList []*repository.Post
	for _, post := range posts {
		postList = append(postList, &repository.Post{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.String(),
			UpdatedAt: post.UpdatedAt.String(),
		})
	}
	return postList, nil
}
