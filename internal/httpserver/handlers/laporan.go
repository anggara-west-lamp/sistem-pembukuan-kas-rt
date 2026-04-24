package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/service"
)

type ReportHandler struct{ svc *service.ReportService }

func NewReportHandler(s *service.ReportService) *ReportHandler { return &ReportHandler{svc:s} }

func (h *ReportHandler) Monthly(c *gin.Context) {
    month := c.Query("month") // YYYY-MM
    if month == "" { c.JSON(http.StatusBadRequest, gin.H{"error":"month required (YYYY-MM)"}); return }
    rpt, err := h.svc.Monthly(c, month)
    if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, rpt)
}
