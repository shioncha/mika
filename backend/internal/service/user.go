package service

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/shioncha/mika/backend/internal/auth"
	"github.com/shioncha/mika/backend/internal/repository"
)

type UserService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	publicKey   *rsa.PublicKey
	privateKey  *rsa.PrivateKey
}

func NewUserService(userRepo repository.UserRepository, sessionRepo repository.SessionRepository, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) *UserService {
	return &UserService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		publicKey:   publicKey,
		privateKey:  privateKey,
	}
}

func (s *UserService) GetByID(c context.Context, userID string) (*repository.User, error) {
	user, err := s.userRepo.GetByID(c, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUsername(c context.Context, userID string, name string) error {
	if err := s.userRepo.UpdateUsername(c, userID, name); err != nil {
		return fmt.Errorf("failed to update username: %w", err)
	}
	return nil
}

func (s *UserService) UpdateEmail(c context.Context, userID string, email string) error {
	email = auth.NormalizeEmail(email)

	exists, err := s.userRepo.EmailExists(c, email)
	if err != nil {
		return fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return fmt.Errorf("email already registered")
	}

	if err := s.userRepo.UpdateEmail(c, userID, email); err != nil {
		return fmt.Errorf("failed to update email: %w", err)
	}
	return nil
}

func (s *UserService) UpdatePassword(c context.Context, userID string, newPassword string) error {
	hashedPassword, err := auth.GenerateHashedPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	if err := s.userRepo.UpdatePassword(c, userID, hashedPassword); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}
	return nil
}
