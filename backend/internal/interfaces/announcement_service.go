package interfaces

import (
	"context"
	"coop/internal/models"
	"mime/multipart"
)

type AnnouncementService interface {
	// Announcement operations
	GetAnnouncementByID(ctx context.Context, id int) (models.AnnouncementWithDetails, error)
	GetAllAnnouncements(ctx context.Context, page, pageSize int) (models.AnnouncementListResponse, error)
	CreateAnnouncement(ctx context.Context, req models.CreateAnnouncementRequest, files []*multipart.FileHeader) (models.AnnouncementWithDetails, error)
	UpdateAnnouncement(ctx context.Context, id int, req models.UpdateAnnouncementRequest) (models.AnnouncementWithDetails, error)
	DeleteAnnouncement(ctx context.Context, id int) error

	// File operations
	AddFilesToAnnouncement(ctx context.Context, announcementID int, files []*multipart.FileHeader) error
	DeleteFileFromAnnouncement(ctx context.Context, fileID int) error

	// Search operations
	SearchAnnouncements(ctx context.Context, keyword string, page, pageSize int) (models.AnnouncementListResponse, error)
}
