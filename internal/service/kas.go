package service

import (
    "context"
    "errors"

    "sistem-pembukuan-kas-rt/internal/models"
    "sistem-pembukuan-kas-rt/internal/repository"
)

type KasService struct { repo repository.KasRepository }

func NewKasService(r repository.KasRepository) *KasService { return &KasService{repo:r} }

func (s *KasService) List(ctx context.Context) ([]models.Kas, error) {
    return s.repo.List(ctx)
}
func (s *KasService) Create(ctx context.Context, k *models.Kas) error {
    if k.Type == "" || k.Category == "" { return errors.New("type & category required") }
    return s.repo.Create(ctx, k)
}

