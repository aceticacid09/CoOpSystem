package routes

import (
    "coop/internal/handler"

    "github.com/gin-gonic/gin"
)

func RegisterAnnouncementRoutes(r *gin.Engine, announcementHandler *handler.AnnouncementHandler) {
    // Serve static files for news attachments
    r.Static("/uploads/news", "/app/uploads/news")

    v1 := r.Group("/api/v1")
    {
        announcements := v1.Group("/announcements")
        {
            // Public routes
            announcements.GET("", announcementHandler.GetAllAnnouncements)
            announcements.GET("/search", announcementHandler.SearchAnnouncements)
            announcements.GET("/:id", announcementHandler.GetAnnouncementByID)

            //todo: Protected routes (add authentication middleware)
            announcements.POST("", announcementHandler.CreateAnnouncement)
            announcements.PUT("/:id", announcementHandler.UpdateAnnouncement)
            announcements.DELETE("/:id", announcementHandler.DeleteAnnouncement)
            announcements.POST("/:id/files", announcementHandler.AddFilesToAnnouncement)
        }
    }
}