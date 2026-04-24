package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"

    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/config"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/db"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/httpserver"
    "github.com/anggara-west-lamp/sistem-pembukuan-kas-rt/internal/models"
)

// @title Sistem Pembukuan Kas RT API
// @version 1.0
// @description Backend API untuk kas RT
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
    _ = godotenv.Load()

    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("config error: %v", err)
    }

    gdb, err := db.OpenGorm(cfg)
    if err != nil {
        log.Fatalf("db error: %v", err)
    }

    // Auto migrate schema
    if err := gdb.AutoMigrate(&models.Role{}, &models.User{}, &models.Kas{}, &models.Transaction{}, &models.Report{}); err != nil {
        log.Fatalf("migrate error: %v", err)
    }

    r := httpserver.SetupRouter(cfg, gdb)

    addr := ":" + cfg.AppPort
    if os.Getenv("PORT") != "" { // fly.io/heroku compat
        addr = ":" + os.Getenv("PORT")
    }
    log.Printf("listening on %s", addr)
    if err := r.Run(addr); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
