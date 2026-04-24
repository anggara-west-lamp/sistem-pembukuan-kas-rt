package service

import (
    "context"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/models"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/repository"
)

type ReportService struct{ trxs repository.TransactionRepository }

func NewReportService(t repository.TransactionRepository) *ReportService { return &ReportService{trxs:t} }

func (s *ReportService) Monthly(ctx context.Context, month string) (*models.Report, error) {
    inc, err := s.trxs.SumByMonth(ctx, month, "masuk")
    if err != nil { return nil, err }
    exp, err := s.trxs.SumByMonth(ctx, month, "keluar")
    if err != nil { return nil, err }
    rpt := &models.Report{Month: month, TotalIncome: inc, TotalExpense: exp, Balance: inc-exp}
    return rpt, nil
}
