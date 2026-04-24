package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"

    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/config"
)

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if !strings.HasPrefix(strings.ToLower(auth), "bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"missing bearer token"})
            return
        }
        tokenStr := strings.TrimSpace(auth[len("Bearer "):])
        _, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
            return []byte(cfg.JWTSecret), nil
        })
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
            return
        }
        c.Next()
    }
}
