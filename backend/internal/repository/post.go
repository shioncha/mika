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
	GetPostsByUserID(ctx context.Context, userID string) ([]*Post, error)

	// 投稿を取得
	GetPostByPostID(ctx context.Context, userID string, postID string) (*Post, error)

	// 投稿を作成
	CreatePost(ctx context.Context, tx *ent.Tx, userID string, post string, tags []string) error

	// 投稿を削除
	DeletePost(ctx context.Context, userID string, postID string) error

	// タグを作成
	CreateTags(ctx context.Context, tx *ent.Tx, userID string, tag []string) ([]string, error)
}
