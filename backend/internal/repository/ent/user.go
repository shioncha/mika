package entrepogitory

import (
	"context"
	"fmt"

	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/users"
	"github.com/shioncha/mika/backend/internal/repository"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*repository.User, error) {
	user, err := r.client.Users.Query().Where(users.EmailEQ(email)).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &repository.User{
		ID:           user.Ulid,
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	isExist, err := r.client.Users.Query().Where(users.EmailEQ(email)).Exist(ctx)

	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}

	return isExist, nil
}

func (r *UserRepository) GetByUlid(ctx context.Context, id string) (*repository.User, error) {
	user, err := r.client.Users.Query().Where(users.UlidEQ(id)).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &repository.User{
		ID:    user.Ulid,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, user *repository.User) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	_, err = tx.Users.Create().
		SetUlid(user.ID).
		SetEmail(user.Email).
		SetName(user.Name).
		SetPasswordHash(user.PasswordHash).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
