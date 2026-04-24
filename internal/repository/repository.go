package repository

import (
    "context"
    "time"

    "gorm.io/gorm"
    "sistem-pembukuan-kas-rt/internal/models"
)

// Interfaces
type UserRepository interface {
    FindByEmail(ctx context.Context, email string) (*models.User, error)
    List(ctx context.Context) ([]models.User, error)
    Create(ctx context.Context, u *models.User) error
}

type RoleRepository interface {
    GetOrCreateByName(ctx context.Context, name string) (*models.Role, error)
}

type KasRepository interface {
    List(ctx context.Context) ([]models.Kas, error)
    Create(ctx context.Context, k *models.Kas) error
}

type TransactionRepository interface {
    Create(ctx context.Context, t *models.Transaction) error
    SumByMonth(ctx context.Context, month string, ttype string) (float64, error)
}

// GORM implementations
type userRepo struct{ db *gorm.DB }
type roleRepo struct{ db *gorm.DB }
type kasRepo struct{ db *gorm.DB }
type trxRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) UserRepository { return &userRepo{db} }
func NewRoleRepo(db *gorm.DB) RoleRepository { return &roleRepo{db} }
func NewKasRepo(db *gorm.DB) KasRepository { return &kasRepo{db} }
func NewTransactionRepo(db *gorm.DB) TransactionRepository { return &trxRepo{db} }

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
    var u models.User
    if err := r.db.WithContext(ctx).Preload("Role").Where("email = ?", email).First(&u).Error; err != nil { return nil, err }
    return &u, nil
}
func (r *userRepo) List(ctx context.Context) ([]models.User, error) {
    var users []models.User
    if err := r.db.WithContext(ctx).Preload("Role").Order("id").Find(&users).Error; err != nil { return nil, err }
    return users, nil
}
func (r *userRepo) Create(ctx context.Context, u *models.User) error {
    u.CreatedAt = time.Now()
    return r.db.WithContext(ctx).Create(u).Error
}

func (r *roleRepo) GetOrCreateByName(ctx context.Context, name string) (*models.Role, error) {
    var role models.Role
    if err := r.db.WithContext(ctx).Where("name = ?", name).First(&role).Error; err == nil {
        return &role, nil
    }
    role = models.Role{Name: name, CreatedAt: time.Now()}
    if err := r.db.WithContext(ctx).Create(&role).Error; err != nil { return nil, err }
    return &role, nil
}

func (r *kasRepo) List(ctx context.Context) ([]models.Kas, error) {
    var items []models.Kas
    if err := r.db.WithContext(ctx).Order("id").Find(&items).Error; err != nil { return nil, err }
    return items, nil
}
func (r *kasRepo) Create(ctx context.Context, k *models.Kas) error {
    k.CreatedAt = time.Now()
    return r.db.WithContext(ctx).Create(k).Error
}

func (r *trxRepo) Create(ctx context.Context, t *models.Transaction) error {
    t.CreatedAt = time.Now()
    return r.db.WithContext(ctx).Create(t).Error
}
func (r *trxRepo) SumByMonth(ctx context.Context, month string, ttype string) (float64, error) {
    // month format: YYYY-MM; we filter by CreatedAt
    start := month + "-01 00:00:00"
    end := month + "-31 23:59:59"
    var sum float64
    q := r.db.WithContext(ctx).Model(&models.Transaction{}).
        Joins("JOIN kas ON kas.id = transactions.kas_id").
        Where("kas.type = ?", ttype).
        Where("transactions.created_at BETWEEN ? AND ?", start, end).
        Select("COALESCE(SUM(transactions.amount),0)")
    if err := q.Scan(&sum).Error; err != nil { return 0, err }
    return sum, nil
}

