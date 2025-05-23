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

func (s *PostService) GetPosts(ctx context.Context, userUlid string) ([]*repository.Post, error) {
	userID, err := s.postRepo.GetUserIDByUlid(ctx, userUlid)
	if err != nil {
		return nil, err
	}

	posts, err := s.postRepo.GetPostsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetPost(ctx context.Context, userUlid string, postID string) (*repository.Post, error) {
	userID, err := s.postRepo.GetUserIDByUlid(ctx, userUlid)
	if err != nil {
		return nil, err
	}

	post, err := s.postRepo.GetPostByPostID(ctx, userID, postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) CreatePost(ctx context.Context, userUlid string, content string) error {
	userID, err := s.postRepo.GetUserIDByUlid(ctx, userUlid)
	if err != nil {
		return err
	}

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

func (s *PostService) DeletePost(ctx context.Context, userUlid string, postID string) error {
	userID, err := s.postRepo.GetUserIDByUlid(ctx, userUlid)
	if err != nil {
		return err
	}

	err = s.postRepo.DeletePost(ctx, userID, postID)
	if err != nil {
		return err
	}
	return nil
}
