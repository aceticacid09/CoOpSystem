package interfaces

import (
    "context"
    "coop/internal/models"
)

type JobPositionRepository interface {
    // Job Position CRUD operations
    GetJobPositionByID(ctx context.Context, id int) (models.JobPositionWithDetails, error)
    GetAllJobPositions(ctx context.Context, limit, offset int) ([]models.JobPositionWithDetails, int, error)
    GetJobPositionsByStatus(ctx context.Context, status string, limit, offset int) ([]models.JobPositionWithDetails, int, error)
    GetJobPositionsByCompany(ctx context.Context, companyID int, limit, offset int) ([]models.JobPositionWithDetails, int, error)
    CreateJobPosition(ctx context.Context, jobPosition models.JobPosition) (int, error)
    UpdateJobPosition(ctx context.Context, id int, jobPosition models.JobPosition) error
    DeleteJobPosition(ctx context.Context, id int) error

    // Attachment operations
    CreateAttachment(ctx context.Context, attachmentID int) (int, error)
    AddFileToAttachment(ctx context.Context, attachmentID int, file models.AttachmentDetail) (int, error)
    GetAttachmentFiles(ctx context.Context, attachmentID int) ([]models.AttachmentDetail, error)
    DeleteAttachment(ctx context.Context, attachmentID int) error

    // File operations
    GetFileByID(ctx context.Context, fileID int) (models.AttachmentDetail, error)
    DeleteFile(ctx context.Context, fileID int) error
    GetJobPositionIDByFileID(ctx context.Context, fileID int) (int, error)

    // Search operations
    SearchJobPositions(ctx context.Context, keyword string, limit, offset int) ([]models.JobPositionWithDetails, int, error)
}