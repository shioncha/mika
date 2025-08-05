package entrepogitory

import (
	"context"
	"fmt"

	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/ent/users"
	"github.com/shioncha/mika/backend/internal/repository"
)

type AuthRepository struct {
	client *ent.Client
}

func NewAuthRepository(client *ent.Client) *AuthRepository {
	return &AuthRepository{
		client: client,
	}
}

func (r *AuthRepository) FindByEmail(ctx context.Context, email string) (*repository.User, error) {
	user, err := r.client.Users.Query().Where(users.EmailEQ(email)).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &repository.User{
		ID:           user.ID,
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (r *AuthRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	isExist, err := r.client.Users.Query().Where(users.EmailEQ(email)).Exist(ctx)

	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}

	return isExist, nil
}

func (r *AuthRepository) GetByID(ctx context.Context, id string) (*repository.User, error) {
	user, err := r.client.Users.Query().Where(users.IDEQ(id)).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &repository.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (r *AuthRepository) Create(ctx context.Context, user *repository.User) error {
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

	created, err := tx.Users.Create().
		SetEmail(user.Email).
		SetName(user.Name).
		SetPasswordHash(user.PasswordHash).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create user: %w", err)
	}
	user.ID = created.ID

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
