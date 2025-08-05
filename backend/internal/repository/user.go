package repository

import "context"

type User struct {
	ID           string
	Email        string
	Name         string
	PasswordHash string
}

type UserRepository interface {
	// メールアドレスからユーザーを検索
	FindByEmail(ctx context.Context, email string) (*User, error)

	// メールアドレスが存在確認
	EmailExists(ctx context.Context, email string) (bool, error)

	// ユーザーIDからユーザーを検索
	GetByID(ctx context.Context, id string) (*User, error)

	// ユーザー名を更新
	UpdateUsername(ctx context.Context, id string, name string) error

	// メールアドレスを更新
	UpdateEmail(ctx context.Context, id string, email string) error

	// パスワードを変更
	UpdatePassword(ctx context.Context, id string, newPassword string) error
}
