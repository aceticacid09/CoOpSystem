// ../internal/service/announcement_service.go
package service

import (
	"context"
	intInterfaces "coop/internal/interfaces"
	intModels "coop/internal/models"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type announcementService struct {
	repo intInterfaces.AnnouncementRepository
}

func NewAnnouncementService(repo intInterfaces.AnnouncementRepository) intInterfaces.AnnouncementService {
	return &announcementService{repo: repo}
}

// GetAnnouncementByID retrieves an announcement by ID
func (s *announcementService) GetAnnouncementByID(ctx context.Context, id int) (intModels.AnnouncementWithDetails, error) {
	return s.repo.GetAnnouncementByID(ctx, id)
}

// GetAllAnnouncements retrieves all announcements with pagination
func (s *announcementService) GetAllAnnouncements(ctx context.Context, page, pageSize int) (intModels.AnnouncementListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	announcements, total, err := s.repo.GetAllAnnouncements(ctx, pageSize, offset)
	if err != nil {
		return intModels.AnnouncementListResponse{}, err
	}

	totalPages := (total + pageSize - 1) / pageSize

	return intModels.AnnouncementListResponse{
		Announcements: announcements,
		Total:         total,
		Page:          page,
		PageSize:      pageSize,
		TotalPages:    totalPages,
	}, nil
}

// CreateAnnouncement creates a new announcement with optional file attachments
func (s *announcementService) CreateAnnouncement(ctx context.Context, req intModels.CreateAnnouncementRequest, files []*multipart.FileHeader) (intModels.AnnouncementWithDetails, error) {
	log.Println("Starting CreateAnnouncement process...")

	// Validate required fields
	if req.Title == "" {
		return intModels.AnnouncementWithDetails{}, fmt.Errorf("title is required")
	}

	var attachmentID *int

	status := intModels.StatusDraft
	if req.Status != nil {
		status = *req.Status
	}

	// Handle file uploads if provided
	if len(files) > 0 {
		log.Printf("Processing %d file(s)...", len(files))

		// Create attachment record
		newAttachmentID, err := s.repo.CreateAttachment(ctx, 0)
		if err != nil {
			return intModels.AnnouncementWithDetails{}, fmt.Errorf("failed to create attachment: %w", err)
		}
		attachmentID = &newAttachmentID

		// Upload files
		for _, fileHeader := range files {
			if err := s.uploadFile(ctx, *attachmentID, fileHeader); err != nil {
				log.Printf("Failed to upload file %s: %v", fileHeader.Filename, err)
				// Continue with other files
			}
		}
	}

	// Create announcement
	announcement := intModels.Announcement{
		Title:        req.Title,
		Categories:   req.Categories,
		Description:  req.Description,
		AttachmentID: attachmentID,
		Status:       status,
		TeacherID:    req.TeacherID,
	}

	postID, err := s.repo.CreateAnnouncement(ctx, announcement)
	if err != nil {
		// Cleanup attachment if announcement creation fails
		if attachmentID != nil {
			s.repo.DeleteAttachment(ctx, *attachmentID)
		}
		return intModels.AnnouncementWithDetails{}, fmt.Errorf("failed to create announcement: %w", err)
	}

	log.Printf("Announcement created successfully with ID: %d", postID)
	// Return the created announcement with details
	return s.repo.GetAnnouncementByID(ctx, postID)
}

// UpdateAnnouncement updates an announcement
func (s *announcementService) UpdateAnnouncement(ctx context.Context, id int, req intModels.UpdateAnnouncementRequest) (intModels.AnnouncementWithDetails, error) {
	// Get existing announcement
	existing, err := s.repo.GetAnnouncementByID(ctx, id)
	if err != nil {
		return intModels.AnnouncementWithDetails{}, fmt.Errorf("announcement not found: %w", err)
	}

	// Update fields if provided
	announcement := intModels.Announcement{
		PostID:       existing.PostID,
		Title:        existing.Title,
		Categories:   existing.Categories,
		Description:  existing.Description,
		AttachmentID: existing.AttachmentID,
		TeacherID:    existing.TeacherID,
	}

	if req.Title != nil {
		announcement.Title = *req.Title
	}
	if req.Categories != nil {
		announcement.Categories = req.Categories
	}
	if req.Description != nil {
		announcement.Description = req.Description
	}
	if req.TeacherID != nil {
		announcement.TeacherID = req.TeacherID
	}

	// Update announcement
	if err := s.repo.UpdateAnnouncement(ctx, id, announcement); err != nil {
		return intModels.AnnouncementWithDetails{}, fmt.Errorf("failed to update announcement: %w", err)
	}

	// Return updated announcement
	return s.repo.GetAnnouncementByID(ctx, id)
}

// DeleteAnnouncement deletes an announcement
func (s *announcementService) DeleteAnnouncement(ctx context.Context, id int) error {
	// Get announcement to find associated files
	announcement, err := s.repo.GetAnnouncementByID(ctx, id)
	if err != nil {
		return fmt.Errorf("announcement not found: %w", err)
	}

	// Delete physical files if they exist
	if announcement.AttachmentID != nil {
		files, err := s.repo.GetAttachmentFiles(ctx, *announcement.AttachmentID)
		if err == nil {
			for _, file := range files {
				if err := os.Remove(file.StoragePath); err != nil {
					log.Printf("Failed to delete file %s: %v", file.StoragePath, err)
				}
			}
		}
	}

	return s.repo.DeleteAnnouncement(ctx, id)
}

// AddFilesToAnnouncement adds files to an existing announcement
func (s *announcementService) AddFilesToAnnouncement(ctx context.Context, announcementID int, files []*multipart.FileHeader) error {
	// Get announcement
	announcement, err := s.repo.GetAnnouncementByID(ctx, announcementID)
	if err != nil {
		return fmt.Errorf("announcement not found: %w", err)
	}

	var attachmentID int

	// Create attachment if doesn't exist
	if announcement.AttachmentID == nil {
		newAttachmentID, err := s.repo.CreateAttachment(ctx, 0)
		if err != nil {
			return fmt.Errorf("failed to create attachment: %w", err)
		}
		attachmentID = newAttachmentID

		// Update announcement with attachment_id
		updateAnnouncement := intModels.Announcement{
			PostID:       announcement.PostID,
			Title:        announcement.Title,
			Categories:   announcement.Categories,
			Description:  announcement.Description,
			AttachmentID: &attachmentID,
			TeacherID:    announcement.TeacherID,
		}
		if err := s.repo.UpdateAnnouncement(ctx, announcementID, updateAnnouncement); err != nil {
			return fmt.Errorf("failed to update announcement: %w", err)
		}
	} else {
		attachmentID = *announcement.AttachmentID
	}

	// Upload files
	for _, fileHeader := range files {
		if err := s.uploadFile(ctx, attachmentID, fileHeader); err != nil {
			log.Printf("Failed to upload file %s: %v", fileHeader.Filename, err)
			// Continue with other files
		}
	}

	return nil
}

// DeleteFileFromAnnouncement deletes a specific file from an announcement
func (s *announcementService) DeleteFileFromAnnouncement(ctx context.Context, fileID int) error {
	// Get file details
	file, err := s.repo.GetFileByID(ctx, fileID)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	// Delete file from database
	if err := s.repo.DeleteFile(ctx, fileID); err != nil {
		return fmt.Errorf("failed to delete file from database: %w", err)
	}

	// Delete physical file
	if err := os.Remove(file.StoragePath); err != nil {
		log.Printf("Warning: Failed to delete physical file %s: %v", file.StoragePath, err)
		// Don't return error as database deletion was successful
	}

	log.Printf("File deleted successfully: %s (ID: %d)", file.FileName, fileID)
	return nil
}

// SearchAnnouncements searches announcements by keyword
func (s *announcementService) SearchAnnouncements(ctx context.Context, keyword string, page, pageSize int) (intModels.AnnouncementListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	announcements, total, err := s.repo.SearchAnnouncements(ctx, keyword, pageSize, offset)
	if err != nil {
		return intModels.AnnouncementListResponse{}, err
	}

	totalPages := (total + pageSize - 1) / pageSize

	return intModels.AnnouncementListResponse{
		Announcements: announcements,
		Total:         total,
		Page:          page,
		PageSize:      pageSize,
		TotalPages:    totalPages,
	}, nil
}

// uploadFile is a helper function to upload a file
func (s *announcementService) uploadFile(ctx context.Context, attachmentID int, fileHeader *multipart.FileHeader) error {
	// Open uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Generate unique filename
	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), strings.ReplaceAll(fileHeader.Filename, ext, ""), ext)

	// Create upload directory
	uploadDir := "/app/uploads/news"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Create destination file
	dst := filepath.Join(uploadDir, filename)
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Copy file content
	if _, err = io.Copy(out, file); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	// Get file type
	fileType := fileHeader.Header.Get("Content-Type")

	// Add file to attachment
	fileDetail := intModels.AttachmentDetail{
		FileName:    fileHeader.Filename,
		FileType:    &fileType,
		StoragePath: dst,
	}

	_, err = s.repo.AddFileToAttachment(ctx, attachmentID, fileDetail)
	if err != nil {
		// Cleanup file if database insert fails
		os.Remove(dst)
		return fmt.Errorf("failed to add file to attachment: %w", err)
	}

	log.Printf("File uploaded successfully: %s", filename)
	return nil
}
