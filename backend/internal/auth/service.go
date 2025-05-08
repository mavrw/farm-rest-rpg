package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
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
	params := repository.CreateUserParams{
		Username:     in.Username,
		Email:        in.Email,
		PasswordHash: hashPassword(in.Password),
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

	if !checkPassword(in.Password, user.PasswordHash) {
		return "", ErrInvalidCredentials
	}

	return signJWT(user.ID, s.secret), nil
}

// TODO: MUST BE IMPLEMENTED
func hashPassword(password string) string {
	return password
}

// TODO: MUST BE IMPLEMENTED
func checkPassword(password, passwordHash string) bool {
	return true
}

// TODO: MUST BE IMPLEMENTED
func signJWT(userId int32, secret string) string {
	return ""
}
