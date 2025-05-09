package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/jwtutil"
	"golang.org/x/crypto/bcrypt"
)

const (
	AuthTokenExpiryHours    = time.Hour * 24
	RefreshTokenExpiryHours = time.Hour * 24 * 14
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService struct {
	q      *repository.Queries
	secret string
}

func NewAuthService(q *repository.Queries, jwtSecret string) *AuthService {
	return &AuthService{q: q, secret: jwtSecret}
}

func (s *AuthService) Register(ctx context.Context, in RegisterInput) error {
	ph, err := hashPassword(in.Password)
	if err != nil {
		return err
	}

	params := repository.CreateUserParams{
		Username:     in.Username,
		Email:        in.Email,
		PasswordHash: ph,
	}

	return s.q.CreateUser(ctx, params)
}

// Login user and issue JWT
func (s *AuthService) Login(ctx context.Context, in LoginInput) (string, error) {
	user, err := s.q.GetUserByEmail(ctx, in.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrUserNotFound
		}
		return "", err
	}

	if err := checkPassword(in.Password, user.PasswordHash); err != nil {
		return "", ErrInvalidCredentials
	}

	return jwtutil.Sign(user.ID, jwtutil.TokenCfg{Secret: s.secret})
}

func (s *AuthService) Refresh(ctx context.Context, token string) (string, string, error) {
	rt, err := s.q.GetRefreshToken(ctx, token)
	if err != nil {
		return "", "", err
	}

	// Rotate
	_ = s.q.RevokeRefreshToken(ctx, token)

	newToken := uuid.NewString()
	expires := time.Now().Add(RefreshTokenExpiryHours)

	_, err = s.q.CreateRefreshToken(ctx, repository.CreateRefreshTokenParams{
		UserID:    rt.UserID,
		Token:     newToken,
		ExpiresAt: expires,
	})
	if err != nil {
		return "", "", err
	}

	at, err := jwtutil.Sign(rt.UserID, jwtutil.TokenCfg{
		Secret: s.secret,
		TTL:    AuthTokenExpiryHours,
	})
	if err != nil {
		return "", "", err
	}

	return at, newToken, nil
}

func (s *AuthService) RevokeRefreshToken(ctx context.Context, token string) error {
	return s.q.RevokeRefreshToken(ctx, token)
}

func hashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

func checkPassword(password, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(password),
		[]byte(passwordHash),
	)
	if err != nil {
		return err
	}

	return nil
}
