package service

import (
    "context"
    "errors"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/models"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/repository"
)

// Fake repos for unit tests

type fakeUserRepo struct {
    users []models.User
    byEmail *models.User
}

func (f *fakeUserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
    if f.byEmail != nil && f.byEmail.Email == email { return f.byEmail, nil }
    return nil, errors.New("not found")
}
func (f *fakeUserRepo) List(ctx context.Context) ([]models.User, error) { return f.users, nil }
func (f *fakeUserRepo) Create(ctx context.Context, u *models.User) error { f.users = append(f.users, *u); return nil }

type fakeRoleRepo struct{ id uint; name string }
func (f *fakeRoleRepo) GetOrCreateByName(ctx context.Context, name string) (*models.Role, error) {
    if f.id == 0 { f.id = 7 }
    if f.name == "" { f.name = name }
    return &models.Role{ID: f.id, Name: f.name}, nil
}

type fakeKasRepo struct{ items []models.Kas }
func (f *fakeKasRepo) List(ctx context.Context) ([]models.Kas, error) { return f.items, nil }
func (f *fakeKasRepo) Create(ctx context.Context, k *models.Kas) error { f.items = append(f.items, *k); return nil }

type fakeTrxRepo struct {
    income map[string]float64
    expense map[string]float64
}
func (f *fakeTrxRepo) Create(ctx context.Context, t *models.Transaction) error { return nil }
func (f *fakeTrxRepo) SumByMonth(ctx context.Context, month string, ttype string) (float64, error) {
    if ttype == "masuk" { return f.income[month], nil }
    if ttype == "keluar" { return f.expense[month], nil }
    return 0, nil
}

// Compile-time assertions
var _ repository.UserRepository = (*fakeUserRepo)(nil)
var _ repository.RoleRepository = (*fakeRoleRepo)(nil)
var _ repository.KasRepository = (*fakeKasRepo)(nil)
var _ repository.TransactionRepository = (*fakeTrxRepo)(nil)
