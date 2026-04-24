package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "sistem-pembukuan-kas-rt/internal/models"
    "sistem-pembukuan-kas-rt/internal/service"
)

type TransactionHandler struct{ svc *service.TransactionService }

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler { return &TransactionHandler{svc:s} }

func (h *TransactionHandler) Create(c *gin.Context) {
    var t models.Transaction
    if err := c.ShouldBindJSON(&t); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    if err := h.svc.Create(c, &t); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusCreated, t)
}

