package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/models"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/service"
)

type KasHandler struct{ svc *service.KasService }

func NewKasHandler(s *service.KasService) *KasHandler { return &KasHandler{svc:s} }

func (h *KasHandler) List(c *gin.Context) {
    items, err := h.svc.List(c)
    if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, items)
}

func (h *KasHandler) Create(c *gin.Context) {
    var k models.Kas
    if err := c.ShouldBindJSON(&k); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    if err := h.svc.Create(c, &k); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusCreated, k)
}
