package service

import (
    "context"
    "testing"
)

func TestReportMonthly_ComputesTotals(t *testing.T) {
    f := &fakeTrxRepo{income: map[string]float64{"2026-04": 100000}, expense: map[string]float64{"2026-04": 40000}}
    svc := NewReportService(f)
    rpt, err := svc.Monthly(context.Background(), "2026-04")
    if err != nil { t.Fatalf("Monthly error: %v", err) }
    if rpt.TotalIncome != 100000 || rpt.TotalExpense != 40000 || rpt.Balance != 60000 {
        t.Fatalf("unexpected rpt: %+v", rpt)
    }
}

