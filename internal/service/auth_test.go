package service

import (
    "context"
    "testing"

    "golang.org/x/crypto/bcrypt"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/config"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/models"
)

func TestAuthLogin_IssuesJWT(t *testing.T) {
    hashed, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
    fu := &fakeUserRepo{byEmail: &models.User{ID: 1, Email: "tester@example.com", Password: string(hashed), Role: models.Role{Name: "warga"}}}
    cfg := &config.Config{JWTSecret: "test-secret", JWTTtlMinutes: 5}
    svc := NewAuthService(cfg, fu)

    token, err := svc.Login(context.Background(), "tester@example.com", "secret")
    if err != nil { t.Fatalf("Login error: %v", err) }
    if len(token) < 20 { t.Fatalf("expected non-empty token, got %q", token) }
}
