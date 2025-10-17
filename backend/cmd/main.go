package main

import (
    "log"
    "net/http"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

    "coop/config"
    "coop/database"

    intHandler "coop/internal/handler"
    intRepo "coop/internal/repository"
    intRoutes "coop/internal/routes"
    intService "coop/internal/service"
)

func main() {
    // โหลด configuration
    cfg, err := config.New()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    db, err := database.NewPostgresDB(cfg)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer db.Close()

    dbPool := db.GetPool()

    // User
    userRepo := intRepo.NewUserRepository(dbPool)
    userService := intService.NewUserService(userRepo)
    userHandler := intHandler.NewUserHandler(userService)

    // Auth (simplified without Keycloak)
    authHandler := intHandler.NewAuthHandler(userService)

    // Announcement&News
    announcementRepo := intRepo.NewAnnouncementRepository(db.GetPool())
    announcementService := intService.NewAnnouncementService(announcementRepo)
    announcementHandler := intHandler.NewAnnouncementHandler(announcementService)

    // Public Documents
    publicDocRepo := intRepo.NewPublicDocRepository(dbPool)
    publicDocService := intService.NewPublicDocService(publicDocRepo)
    publicDocHandler := intHandler.NewPublicDocHandler(publicDocService)

    // Job Position
    jobPositionRepo := intRepo.NewJobPositionRepository(dbPool)
    jobPositionService := intService.NewJobPositionService(jobPositionRepo)
    jobPositionHandler := intHandler.NewJobPositionHandler(jobPositionService)

    r := gin.Default()

    // Enable CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "healthy"})
    })

    // Register routes
    intRoutes.RegisterUserRoutes(r, userHandler, authHandler)
    intRoutes.RegisterAnnouncementRoutes(r, announcementHandler)
    intRoutes.RegisterPublicDocRoutes(r, publicDocHandler)
    intRoutes.RegisterJobPositionRoutes(r, jobPositionHandler)

    if err := r.Run(":8000"); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}