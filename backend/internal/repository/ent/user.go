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
		ID:           user.ID,
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

func (r *UserRepository) GetByID(ctx context.Context, id string) (*repository.User, error) {
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

func (r *UserRepository) UpdateUsername(ctx context.Context, id string, name string) error {
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

	err = tx.Users.UpdateOneID(id).SetName(name).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update username: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *UserRepository) UpdateEmail(ctx context.Context, id string, email string) error {
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

	err = tx.Users.UpdateOneID(id).SetEmail(email).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update email: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id string, newPassword string) error {
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

	err = tx.Users.UpdateOneID(id).SetPasswordHash(newPassword).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update password: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
