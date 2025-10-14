package routes

import (
    "coop/internal/handler"

    "github.com/gin-gonic/gin"
)

func RegisterPublicDocRoutes(r *gin.Engine, publicDocHandler *handler.PublicDocHandler) {
    // Serve static files for public documents
    r.Static("/uploads/public_docs", "/app/uploads/public_docs")

    v1 := r.Group("/api/v1")
    {
        publicDocs := v1.Group("/public-documents")
        {
            // Public routes
            publicDocs.GET("", publicDocHandler.GetAllPublicDocs)
            publicDocs.GET("/search", publicDocHandler.SearchPublicDocs)
            publicDocs.GET("/:id", publicDocHandler.GetPublicDocByID)
            publicDocs.GET("/:id/download", publicDocHandler.DownloadPublicDoc)

            // Protected routes - only teachers/admins can create/update/delete
            // TODO: Add authentication middleware
            publicDocs.POST("", publicDocHandler.CreatePublicDoc)
            publicDocs.PUT("/:id", publicDocHandler.UpdatePublicDoc)	
            publicDocs.DELETE("/:id", publicDocHandler.DeletePublicDoc)
        }
    }
}