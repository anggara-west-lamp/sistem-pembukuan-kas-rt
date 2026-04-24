package httpserver

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "sistem-pembukuan-kas-rt/internal/config"
    "sistem-pembukuan-kas-rt/internal/httpserver/handlers"
    "sistem-pembukuan-kas-rt/internal/httpserver/middleware"
    "sistem-pembukuan-kas-rt/internal/repository"
    "sistem-pembukuan-kas-rt/internal/service"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
    r := gin.Default()

    r.GET("/api/v1/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status":"ok"}) })

    // repositories
    userRepo := repository.NewUserRepo(db)
    roleRepo := repository.NewRoleRepo(db)
    kasRepo := repository.NewKasRepo(db)
    trxRepo := repository.NewTransactionRepo(db)

    // services
    authSvc := service.NewAuthService(cfg, userRepo)
    userSvc := service.NewUserService(userRepo, roleRepo)
    kasSvc := service.NewKasService(kasRepo)
    trxSvc := service.NewTransactionService(trxRepo, kasRepo)
    rptSvc := service.NewReportService(trxRepo)

    // handlers
    authH := handlers.NewAuthHandler(authSvc)
    userH := handlers.NewUserHandler(userSvc)
    kasH := handlers.NewKasHandler(kasSvc)
    trxH := handlers.NewTransactionHandler(trxSvc)
    rptH := handlers.NewReportHandler(rptSvc)

    v1 := r.Group("/api/v1")
    {
        v1.POST("/auth/login", authH.Login)

        // protected
        authz := v1.Group("")
        authz.Use(middleware.JWTAuth(cfg))
        {
            authz.GET("/users", userH.List)
            authz.POST("/users", userH.Create)

            authz.GET("/kas", kasH.List)
            authz.POST("/kas", kasH.Create)

            authz.POST("/transaksi", trxH.Create)

            authz.GET("/laporan", rptH.Monthly)
        }
    }

    return r
}

