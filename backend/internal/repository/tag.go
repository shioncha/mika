package repository

import "context"

type Tag struct {
	ID   string
	Name string
}

type TagRepository interface {
	// タグ一覧を取得
	GetTags(ctx context.Context, userID string) ([]*Tag, error)

	// タグからポストを取得
	GetPostsByTag(ctx context.Context, userID string, tag string) ([]*Post, error)
}
