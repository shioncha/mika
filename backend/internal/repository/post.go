package repository

import (
	"context"

	"github.com/shioncha/mika/backend/ent"
)

type Post struct {
	ID        string
	Content   string
	CreatedAt string
	UpdatedAt string
}

type PostRepository interface {
	// ユーザーの投稿一覧を取得
	GetPostsByUserID(ctx context.Context, userID int) ([]*Post, error)

	// 投稿を取得
	GetPostByPostID(ctx context.Context, userID int, postID string) (*Post, error)

	// 投稿を作成
	CreatePost(ctx context.Context, userID int, post string, tags []int) error

	// 投稿を削除
	DeletePost(ctx context.Context, userID int, postID string) error

	// ユーザーのULIDからユーザーIDを取得
	GetUserIDByUlid(ctx context.Context, ulid string) (int, error)

	// タグを作成
	CreateTags(ctx context.Context, tx *ent.Tx, userID int, tag []string) ([]int, error)
}
