package interfaces

import (
    "context"
    "coop/internal/models"
)

type AnnouncementRepository interface {
    // Announcement CRUD operations
    GetAnnouncementByID(ctx context.Context, id int) (models.AnnouncementWithDetails, error)
    GetAllAnnouncements(ctx context.Context, limit, offset int) ([]models.AnnouncementWithDetails, int, error)
    CreateAnnouncement(ctx context.Context, announcement models.Announcement) (int, error)
    UpdateAnnouncement(ctx context.Context, id int, announcement models.Announcement) error
    DeleteAnnouncement(ctx context.Context, id int) error

    // Attachment operations
    CreateAttachment(ctx context.Context, attachmentID int) (int, error)
    AddFileToAttachment(ctx context.Context, attachmentID int, file models.AttachmentDetail) (int, error)
    GetAttachmentFiles(ctx context.Context, attachmentID int) ([]models.AttachmentDetail, error)
    DeleteAttachment(ctx context.Context, attachmentID int) error
    
    // Search operations
    SearchAnnouncements(ctx context.Context, keyword string, limit, offset int) ([]models.AnnouncementWithDetails, int, error)
}