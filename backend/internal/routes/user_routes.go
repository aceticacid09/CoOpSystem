package routes

import (
    "coop/internal/handler"

    "github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) {
    // Serve static files for profile images
    r.Static("/uploads/profile", "/app/uploads/profile")

    v1 := r.Group("/api/v1")
    {
        // Authentication routes
        auth := v1.Group("/auth")
        {
            auth.POST("/login", authHandler.Login)
            auth.POST("/refresh", authHandler.RefreshToken)
            auth.POST("/logout", authHandler.Logout)
        }

        // User routes
        users := v1.Group("/users")
        {
            users.GET("", userHandler.GetAllUsers)
            users.GET("/:id", userHandler.GetUserByID)
            users.GET("/username/:username", userHandler.GetUserByUsername)
            users.POST("", userHandler.CreateUser)
            users.PUT("/:id", userHandler.UpdateUser)
            users.DELETE("/:id", userHandler.DeleteUser)
        }

        // Department routes
        departments := v1.Group("/departments")
        {
            departments.GET("", userHandler.GetAllDepartments)
            departments.GET("/:id", userHandler.GetDepartmentByID)
        }
    }
}