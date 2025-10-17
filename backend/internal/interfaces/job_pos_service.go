package interfaces

import (
    "context"
    "coop/internal/models"
    "mime/multipart"
)

type JobPositionService interface {
    // Job Position operations
    GetJobPositionByID(ctx context.Context, id int) (models.JobPositionWithDetails, error)
    GetAllJobPositions(ctx context.Context, page, pageSize int) (models.JobPositionListResponse, error)
    GetJobPositionsByStatus(ctx context.Context, status string, page, pageSize int) (models.JobPositionListResponse, error)
    GetJobPositionsByCompany(ctx context.Context, companyID int, page, pageSize int) (models.JobPositionListResponse, error)
    CreateJobPosition(ctx context.Context, req models.CreateJobPositionRequest, files []*multipart.FileHeader) (models.JobPositionWithDetails, error)
    UpdateJobPosition(ctx context.Context, id int, req models.UpdateJobPositionRequest) (models.JobPositionWithDetails, error)
    DeleteJobPosition(ctx context.Context, id int) error

    // File operations
    AddFilesToJobPosition(ctx context.Context, jobPositionID int, files []*multipart.FileHeader) error
    DeleteFileFromJobPosition(ctx context.Context, fileID int) error

    // Search operations
    SearchJobPositions(ctx context.Context, keyword string, page, pageSize int) (models.JobPositionListResponse, error)
}