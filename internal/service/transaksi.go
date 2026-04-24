package service

import (
    "context"
    "errors"

    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/models"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/repository"
)

type TransactionService struct {
    trxs repository.TransactionRepository
    kas  repository.KasRepository
}

func NewTransactionService(t repository.TransactionRepository, k repository.KasRepository) *TransactionService {
    return &TransactionService{trxs: t, kas: k}
}

func (s *TransactionService) Create(ctx context.Context, t *models.Transaction) error {
    if t.UserID == 0 || t.KasID == 0 || t.Amount <= 0 { return errors.New("user_id, kas_id, amount required") }
    if t.Status == "" { t.Status = "pending" }
    return s.trxs.Create(ctx, t)
}
