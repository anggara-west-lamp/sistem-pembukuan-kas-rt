package service

import (
    "context"
    "errors"

    "golang.org/x/crypto/bcrypt"
    "sistem-pembukuan-kas-rt/internal/models"
    "sistem-pembukuan-kas-rt/internal/repository"
)

type UserService struct {
    users repository.UserRepository
    roles repository.RoleRepository
}

func NewUserService(u repository.UserRepository, r repository.RoleRepository) *UserService {
    return &UserService{users: u, roles: r}
}

func (s *UserService) List(ctx context.Context) ([]models.User, error) {
    return s.users.List(ctx)
}

func (s *UserService) Create(ctx context.Context, u *models.User) error {
    if u.Email == "" || u.Password == "" { return errors.New("email & password required") }
    role, err := s.roles.GetOrCreateByName(ctx, "warga")
    if err != nil { return err }
    u.RoleID = role.ID
    hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil { return err }
    u.Password = string(hashed)
    return s.users.Create(ctx, u)
}

