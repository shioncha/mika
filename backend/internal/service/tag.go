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

func (s *TagService) GetTags(ctx context.Context, userID string) ([]*repository.Tag, error) {
	tags, err := s.tagRepo.GetTags(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (s *TagService) GetPostsByTag(ctx context.Context, userID string, tagID string, limit int, cursor string) ([]*repository.Post, string, error) {
	posts, err := s.tagRepo.GetPostsByTag(ctx, userID, tagID, limit+1, cursor)
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
