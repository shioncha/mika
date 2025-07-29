package service

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/shioncha/mika/backend/internal/auth"
	"github.com/shioncha/mika/backend/internal/repository"
	"github.com/shioncha/mika/backend/internal/utils"
)

type AuthService struct {
	userRepo   repository.UserRepository
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func NewAuthService(userRepo repository.UserRepository, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

type SignUpParams struct {
	Email           string
	Name            string
	Password        string
	PasswordConfirm string
}

type SignUpResult struct {
	UserID       string
	Token        string
	RefreshToken string
}

func (s *AuthService) SignUp(ctx context.Context, params SignUpParams) (*SignUpResult, error) {
	// メールアドレスの正規化
	email := auth.NormalizeEmail(params.Email)

	// メールアドレスの重複チェック
	exists, err := s.userRepo.EmailExists(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("email already registered")
	}

	// パスワードのハッシュ化
	hashedPassword, err := auth.GenerateHashedPassword(params.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// ユーザーID生成
	userID := utils.GenerateULID()

	// ユーザー作成
	user := &repository.User{
		ID:           userID,
		Email:        email,
		Name:         params.Name,
		PasswordHash: hashedPassword,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// JWTトークン生成
	token, err := auth.GenerateJWT(userID, s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// リフレッシュトークン生成
	refreshToken := "refresh_token" // TODO: Implement refresh token generation

	return &SignUpResult{
		UserID:       userID,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

type SignInParams struct {
	Email    string
	Password string
}

type SignInResult struct {
	UserID       string
	Token        string
	RefreshToken string
}

func (s *AuthService) SignIn(ctx context.Context, params SignInParams) (*SignInResult, error) {
	// メールアドレスの正規化
	email := auth.NormalizeEmail(params.Email)

	// ユーザー検索
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// パスワード検証
	if err := auth.ComparePassword(user.PasswordHash, params.Password); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// JWTトークン生成
	token, err := auth.GenerateJWT(user.ID, s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// リフレッシュトークン生成
	refreshToken := "refresh_token" // TODO: Implement refresh token generation

	return &SignInResult{
		UserID:       user.ID,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) GetByUlid(ctx context.Context, userUlid string) (*repository.User, error) {
	user, err := s.userRepo.GetByUlid(ctx, userUlid)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ulid: %w", err)
	}
	return user, nil
}
