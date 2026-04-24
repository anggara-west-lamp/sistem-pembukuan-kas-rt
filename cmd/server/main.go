package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"

    "sistem-pembukuan-kas-rt/internal/config"
    "sistem-pembukuan-kas-rt/internal/db"
    "sistem-pembukuan-kas-rt/internal/httpserver"
    "sistem-pembukuan-kas-rt/internal/models"
)

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

