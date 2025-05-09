package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
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

	return signJWT(user.ID, s.secret)
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

func signJWT(userId int32, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(AuthTokenExpiryHours).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
