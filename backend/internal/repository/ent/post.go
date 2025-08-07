package entrepogitory

import (
	"context"
	"fmt"

	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/posts"
	"github.com/shioncha/mika/backend/ent/tags"
	"github.com/shioncha/mika/backend/internal/repository"
)

type PostRepository struct {
	client *ent.Client
}

func NewPostRepository(client *ent.Client) *PostRepository {
	return &PostRepository{
		client: client,
	}
}

func (r *PostRepository) GetPostsByUserID(ctx context.Context, userID string, limit int, cursor string) ([]*repository.Post, error) {
	query := r.client.Posts.Query().Order(ent.Desc(posts.FieldCreatedAt)).Where(posts.UserIDEQ(userID))
	if cursor != "" {
		query = query.Where(posts.IDLT(cursor))
	}
	postList, err := query.Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}

	var posts []*repository.Post
	for _, post := range postList {
		posts = append(posts, &repository.Post{
			ID:          post.ID,
			Content:     post.Content,
			HasCheckbox: post.HasCheckbox,
			IsChecked:   post.IsChecked,
			CreatedAt:   post.CreatedAt.String(),
			UpdatedAt:   post.UpdatedAt.String(),
		})
	}

	return posts, nil
}

func (r *PostRepository) GetPostByPostID(ctx context.Context, userID string, postID string) (*repository.Post, error) {
	post, err := r.client.Posts.Query().
		Where(posts.UserIDEQ(userID), posts.IDEQ(postID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &repository.Post{
		ID:          post.ID,
		Content:     post.Content,
		HasCheckbox: post.HasCheckbox,
		IsChecked:   post.IsChecked,
		CreatedAt:   post.CreatedAt.String(),
		UpdatedAt:   post.UpdatedAt.String(),
	}, nil
}

func (r *PostRepository) CreatePost(ctx context.Context, tx *ent.Tx, userID string, content string, tags []string, hasCheckbox bool) error {
	_, err := tx.Posts.Create().
		SetUserID(userID).
		SetContent(content).
		SetHasCheckbox(hasCheckbox).
		AddTagIDs(tags...).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) UpdateContent(ctx context.Context, tx *ent.Tx, userID string, postID string, content string, tags []string, hasCheckbox bool) error {
	_, err := tx.Posts.Update().
		Where(posts.UserIDEQ(userID), posts.IDEQ(postID)).
		SetContent(content).
		SetHasCheckbox(hasCheckbox).
		Save(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Posts.Update().
		Where(posts.UserIDEQ(userID), posts.IDEQ(postID)).
		ClearTags().
		AddTagIDs(tags...).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) UpdateCheckbox(ctx context.Context, userID string, postID string, isChecked bool) error {
	_, err := r.client.Posts.Update().
		Where(posts.UserIDEQ(userID), posts.IDEQ(postID)).
		SetIsChecked(isChecked).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) DeletePost(ctx context.Context, userID string, postID string) error {
	_, err := r.client.Posts.Delete().
		Where(posts.UserIDEQ(userID), posts.IDEQ(postID)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) CreateTags(ctx context.Context, tx *ent.Tx, userID string, tagList []string) ([]string, error) {
	var tagIDs []string

	existingTags, err := tx.Tags.Query().
		Where(tags.UserIDEQ(userID), tags.TagIn(tagList...)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query existing tags: %w", err)
	}

	existingTagMap := make(map[string]string)
	for _, tag := range existingTags {
		existingTagMap[tag.Tag] = tag.ID
	}

	for _, tag := range tagList {
		if id, ok := existingTagMap[tag]; ok {
			tagIDs = append(tagIDs, id)
			continue
		}

		res, err := tx.Tags.Create().
			SetTag(tag).
			SetUserID(userID).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create tag")
		}
		tagIDs = append(tagIDs, res.ID)
	}

	return tagIDs, nil
}
