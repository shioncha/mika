package repository

import "context"

type AuthRepository interface {
	// メールアドレスからユーザーを検索
	FindByEmail(ctx context.Context, email string) (*User, error)

	// メールアドレスが存在確認
	EmailExists(ctx context.Context, email string) (bool, error)

	// ユーザーIDからユーザーを検索
	GetByID(ctx context.Context, id string) (*User, error)

	// 新規ユーザーを作成
	Create(ctx context.Context, user *User) error
}
