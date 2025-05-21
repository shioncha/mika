package repository

import "context"

type Tag struct {
	ID   string
	Name string
}

type TagRepository interface {
	// タグ一覧を取得
	GetTags(ctx context.Context, userID int) ([]*Tag, error)

	// タグからポストを取得
	GetPostsByTag(ctx context.Context, userID int, tag string) ([]*Post, error)

	// ユーザーのULIDからユーザーIDを取得
	GetUserIDByUlid(ctx context.Context, ulid string) (int, error)
}
