package db

import (
    "fmt"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/config"
)

func OpenGorm(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
    )
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
