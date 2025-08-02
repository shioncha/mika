package service

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/shioncha/mika/backend/internal/auth"
	"github.com/shioncha/mika/backend/internal/repository"
)

type AuthService struct {
	authRepo    repository.AuthRepository
	sessionRepo repository.SessionRepository
	publicKey   *rsa.PublicKey
	privateKey  *rsa.PrivateKey
}

func NewAuthService(authRepo repository.AuthRepository, sessionRepo repository.SessionRepository, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) *AuthService {
	return &AuthService{
		authRepo:    authRepo,
		sessionRepo: sessionRepo,
		publicKey:   publicKey,
		privateKey:  privateKey,
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

func (s *AuthService) SignUp(c context.Context, params SignUpParams, deviceInfo string, ipAddress string) (*SignUpResult, error) {
	// メールアドレスの正規化
	email := auth.NormalizeEmail(params.Email)

	// メールアドレスの重複チェック
	exists, err := s.authRepo.EmailExists(c, email)
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

	// ユーザー作成
	user := &repository.User{
		Email:        email,
		Name:         params.Name,
		PasswordHash: hashedPassword,
	}

	if err := s.authRepo.Create(c, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// JWTトークン生成
	token, err := auth.GenerateJWT(user.ID, s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// リフレッシュトークン生成
	refreshToken, err := s.sessionRepo.CreateSession(c, user.ID, deviceInfo, ipAddress, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &SignUpResult{
		UserID:       user.ID,
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

func (s *AuthService) SignIn(c context.Context, params SignInParams, deviceInfo string, ipAddress string) (*SignInResult, error) {
	// メールアドレスの正規化
	email := auth.NormalizeEmail(params.Email)

	// ユーザー検索
	user, err := s.authRepo.FindByEmail(c, email)
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
	refreshToken, err := s.sessionRepo.CreateSession(c, user.ID, deviceInfo, ipAddress, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &SignInResult{
		UserID:       user.ID,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

type RefreshAccessTokenResult struct {
	Token string
}

func (s *AuthService) RefreshAccessToken(c context.Context, oldRefreshToken string) (*RefreshAccessTokenResult, error) {
	session, err := s.sessionRepo.GetSession(c, oldRefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	if err := s.sessionRepo.ExtendSession(c, oldRefreshToken, 7*24*time.Hour); err != nil {
		return nil, fmt.Errorf("failed to extend session: %w", err)
	}

	newAccessToken, err := auth.GenerateJWT(session.UserID, s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new access token: %w", err)
	}

	return &RefreshAccessTokenResult{
		Token: newAccessToken,
	}, nil
}

func (s *AuthService) SignOut(c context.Context, refreshToken string) error {
	if err := s.sessionRepo.DeleteSession(c, refreshToken); err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}
