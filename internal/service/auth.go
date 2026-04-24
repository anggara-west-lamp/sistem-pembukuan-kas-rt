package service

import (
    "context"
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/config"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/repository"
)

type AuthService struct {
    cfg *config.Config
    users repository.UserRepository
}

func NewAuthService(cfg *config.Config, users repository.UserRepository) *AuthService {
    return &AuthService{cfg: cfg, users: users}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
    u, err := s.users.FindByEmail(ctx, email)
    if err != nil { return "", errors.New("invalid credentials") }
    if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
        return "", errors.New("invalid credentials")
    }
    claims := jwt.MapClaims{
        "sub": u.ID,
        "email": u.Email,
        "role": u.Role.Name,
        "exp": time.Now().Add(time.Duration(s.cfg.JWTTtlMinutes) * time.Minute).Unix(),
        "iat": time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.cfg.JWTSecret))
}
