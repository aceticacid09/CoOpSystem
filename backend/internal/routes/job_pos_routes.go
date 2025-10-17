package routes

import (
    "coop/internal/handler"

    "github.com/gin-gonic/gin"
)

func RegisterJobPositionRoutes(r *gin.Engine, jobPositionHandler *handler.JobPositionHandler) {
    // Serve static files for job position attachments
    r.Static("/uploads/documents", "/app/uploads/documents")

    v1 := r.Group("/api/v1")
    {
        jobPositions := v1.Group("/job-positions")
        {
            // Public routes
            jobPositions.GET("", jobPositionHandler.GetAllJobPositions)
            jobPositions.GET("/search", jobPositionHandler.SearchJobPositions)						//example: .../search?keyword=ระบบ
            jobPositions.GET("/:id", jobPositionHandler.GetJobPositionByID)
            jobPositions.GET("/status/:status", jobPositionHandler.GetJobPositionsByStatus)			//get job position by status (open, close, pending)
            jobPositions.GET("/company/:company_id", jobPositionHandler.GetJobPositionsByCompany)	

            // Protected routes - companies/teacher can create/update/delete
            // TODO: Add authentication middleware
            jobPositions.POST("", jobPositionHandler.CreateJobPosition)
            jobPositions.PUT("/:id", jobPositionHandler.UpdateJobPosition)
            jobPositions.DELETE("/:id", jobPositionHandler.DeleteJobPosition)
            jobPositions.POST("/:id/files", jobPositionHandler.AddFilesToJobPosition)
            jobPositions.DELETE("/files/:file_id", jobPositionHandler.DeleteFileFromJobPosition)
        }
    }
}