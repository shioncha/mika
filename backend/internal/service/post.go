package service

import (
	"context"

	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/internal/repository"
	"github.com/shioncha/mika/backend/internal/utils"
)

type PostService struct {
	client   *ent.Client
	postRepo repository.PostRepository
}

func NewPostService(client *ent.Client, postRepo repository.PostRepository) *PostService {
	return &PostService{
		client:   client,
		postRepo: postRepo,
	}
}

func (s *PostService) GetPosts(ctx context.Context, userID string, limit int, cursor string) ([]*repository.Post, string, error) {
	posts, err := s.postRepo.GetPostsByUserID(ctx, userID, limit+1, cursor)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""
	if len(posts) > limit {
		nextCursor = posts[len(posts)-1].ID
		posts = posts[:limit]
	}

	return posts, nextCursor, nil
}

func (s *PostService) GetPost(ctx context.Context, userID string, postID string) (*repository.Post, error) {
	post, err := s.postRepo.GetPostByPostID(ctx, userID, postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) CreatePost(ctx context.Context, userID string, content string) error {
	tags, err := utils.ExtractHashtags(content)
	if err != nil {
		return err
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	tagIDs, err := s.postRepo.CreateTags(ctx, tx, userID, tags)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.postRepo.CreatePost(ctx, tx, userID, content, tagIDs)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *PostService) UpdateContent(ctx context.Context, userID string, postID string, content string) error {
	tags, err := utils.ExtractHashtags(content)
	if err != nil {
		return err
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	tagIDs, err := s.postRepo.CreateTags(ctx, tx, userID, tags)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.postRepo.UpdateContent(ctx, tx, userID, postID, content, tagIDs)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *PostService) UpdateCheckbox(ctx context.Context, userID string, postID string, isChecked bool) error {
	err := s.postRepo.UpdateCheckbox(ctx, userID, postID, isChecked)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(ctx context.Context, userID string, postID string) error {
	err := s.postRepo.DeletePost(ctx, userID, postID)
	if err != nil {
		return err
	}
	return nil
}
