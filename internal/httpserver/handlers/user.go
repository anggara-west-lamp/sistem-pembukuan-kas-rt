package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "sistem-pembukuan-kas-rt/internal/models"
    "sistem-pembukuan-kas-rt/internal/service"
)

type UserHandler struct{ svc *service.UserService }

func NewUserHandler(s *service.UserService) *UserHandler { return &UserHandler{svc:s} }

func (h *UserHandler) List(c *gin.Context) {
    users, err := h.svc.List(c)
    if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Create(c *gin.Context) {
    var u models.User
    if err := c.ShouldBindJSON(&u); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    if err := h.svc.Create(c, &u); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    c.JSON(http.StatusCreated, u)
}

