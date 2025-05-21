package service

import (
	"context"

	"github.com/shioncha/mika/backend/internal/repository"
)

type TagService struct {
	tagRepo repository.TagRepository
}

func NewTagService(tagRepo repository.TagRepository) *TagService {
	return &TagService{
		tagRepo: tagRepo,
	}
}

func (s *TagService) GetTags(ctx context.Context, userUlid string) ([]*repository.Tag, error) {
	userID, err := s.tagRepo.GetUserIDByUlid(ctx, userUlid)
	if err != nil {
		return nil, err
	}

	tags, err := s.tagRepo.GetTags(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (s *TagService) GetPostsByTag(ctx context.Context, userUlid string, tagID string) ([]*repository.Post, error) {
	userID, err := s.tagRepo.GetUserIDByUlid(ctx, userUlid)
	if err != nil {
		return nil, err
	}

	posts, err := s.tagRepo.GetPostsByTag(ctx, userID, tagID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
