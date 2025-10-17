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

type jobPositionService struct {
    repo intInterfaces.JobPositionRepository
}

func NewJobPositionService(repo intInterfaces.JobPositionRepository) intInterfaces.JobPositionService {
    return &jobPositionService{repo: repo}
}

// GetJobPositionByID retrieves a job position by ID
func (s *jobPositionService) GetJobPositionByID(ctx context.Context, id int) (intModels.JobPositionWithDetails, error) {
    return s.repo.GetJobPositionByID(ctx, id)
}

// GetAllJobPositions retrieves all job positions with pagination
func (s *jobPositionService) GetAllJobPositions(ctx context.Context, page, pageSize int) (intModels.JobPositionListResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    jobPositions, total, err := s.repo.GetAllJobPositions(ctx, pageSize, offset)
    if err != nil {
        return intModels.JobPositionListResponse{}, err
    }

    totalPages := (total + pageSize - 1) / pageSize

    return intModels.JobPositionListResponse{
        JobPositions: jobPositions,
        Total:        total,
        Page:         page,
        PageSize:     pageSize,
        TotalPages:   totalPages,
    }, nil
}

// GetJobPositionsByStatus retrieves job positions filtered by status
func (s *jobPositionService) GetJobPositionsByStatus(ctx context.Context, status string, page, pageSize int) (intModels.JobPositionListResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    // Validate status
    validStatuses := map[string]bool{"open": true, "closed": true, "pending": true}
    if !validStatuses[status] {
        return intModels.JobPositionListResponse{}, fmt.Errorf("invalid status: %s", status)
    }

    offset := (page - 1) * pageSize
    jobPositions, total, err := s.repo.GetJobPositionsByStatus(ctx, status, pageSize, offset)
    if err != nil {
        return intModels.JobPositionListResponse{}, err
    }

    totalPages := (total + pageSize - 1) / pageSize

    return intModels.JobPositionListResponse{
        JobPositions: jobPositions,
        Total:        total,
        Page:         page,
        PageSize:     pageSize,
        TotalPages:   totalPages,
    }, nil
}

// GetJobPositionsByCompany retrieves job positions filtered by company
func (s *jobPositionService) GetJobPositionsByCompany(ctx context.Context, companyID int, page, pageSize int) (intModels.JobPositionListResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    jobPositions, total, err := s.repo.GetJobPositionsByCompany(ctx, companyID, pageSize, offset)
    if err != nil {
        return intModels.JobPositionListResponse{}, err
    }

    totalPages := (total + pageSize - 1) / pageSize

    return intModels.JobPositionListResponse{
        JobPositions: jobPositions,
        Total:        total,
        Page:         page,
        PageSize:     pageSize,
        TotalPages:   totalPages,
    }, nil
}

// CreateJobPosition creates a new job position with optional file attachments
func (s *jobPositionService) CreateJobPosition(ctx context.Context, req intModels.CreateJobPositionRequest, files []*multipart.FileHeader) (intModels.JobPositionWithDetails, error) {
    log.Println("Starting CreateJobPosition process...")

    // Validate required fields
    if req.Title == "" {
        return intModels.JobPositionWithDetails{}, fmt.Errorf("title is required")
    }

    // Set default status if not provided
    status := "open"
    if req.Status != nil {
        validStatuses := map[string]bool{"open": true, "closed": true, "pending": true}
        if !validStatuses[*req.Status] {
            return intModels.JobPositionWithDetails{}, fmt.Errorf("invalid status: %s", *req.Status)
        }
        status = *req.Status
    }

    var attachmentID *int

    // Handle file uploads if provided
    if len(files) > 0 {
        log.Printf("Processing %d file(s)...", len(files))

        // Create attachment record
        newAttachmentID, err := s.repo.CreateAttachment(ctx, 0)
        if err != nil {
            return intModels.JobPositionWithDetails{}, fmt.Errorf("failed to create attachment: %w", err)
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

    // Create job position
    jobPosition := intModels.JobPosition{
        Title:        req.Title,
        Description:  req.Description,
        CompanyID:    req.CompanyID,
        AttachmentID: attachmentID,
        Status:       status,
    }

    jobID, err := s.repo.CreateJobPosition(ctx, jobPosition)
    if err != nil {
        // Cleanup attachment if job position creation fails
        if attachmentID != nil {
            s.repo.DeleteAttachment(ctx, *attachmentID)
        }
        return intModels.JobPositionWithDetails{}, fmt.Errorf("failed to create job position: %w", err)
    }

    log.Printf("Job position created successfully with ID: %d", jobID)
    // Return the created job position with details
    return s.repo.GetJobPositionByID(ctx, jobID)
}

// UpdateJobPosition updates a job position
func (s *jobPositionService) UpdateJobPosition(ctx context.Context, id int, req intModels.UpdateJobPositionRequest) (intModels.JobPositionWithDetails, error) {
    // Get existing job position
    existing, err := s.repo.GetJobPositionByID(ctx, id)
    if err != nil {
        return intModels.JobPositionWithDetails{}, fmt.Errorf("job position not found: %w", err)
    }

    // Update fields if provided
    jobPosition := intModels.JobPosition{
        JobID:        existing.JobID,
        Title:        existing.Title,
        Description:  existing.Description,
        CompanyID:    existing.CompanyID,
        AttachmentID: existing.AttachmentID,
        Status:       existing.Status,
    }

    if req.Title != nil {
        jobPosition.Title = *req.Title
    }
    if req.Description != nil {
        jobPosition.Description = req.Description
    }
    if req.CompanyID != nil {
        jobPosition.CompanyID = req.CompanyID
    }
    if req.Status != nil {
        validStatuses := map[string]bool{"open": true, "closed": true, "pending": true}
        if !validStatuses[*req.Status] {
            return intModels.JobPositionWithDetails{}, fmt.Errorf("invalid status: %s", *req.Status)
        }
        jobPosition.Status = *req.Status
    }

    // Update job position
    if err := s.repo.UpdateJobPosition(ctx, id, jobPosition); err != nil {
        return intModels.JobPositionWithDetails{}, fmt.Errorf("failed to update job position: %w", err)
    }

    // Return updated job position
    return s.repo.GetJobPositionByID(ctx, id)
}

// DeleteJobPosition deletes a job position
func (s *jobPositionService) DeleteJobPosition(ctx context.Context, id int) error {
    // Get job position to find associated files
    jobPosition, err := s.repo.GetJobPositionByID(ctx, id)
    if err != nil {
        return fmt.Errorf("job position not found: %w", err)
    }

    // Delete associated files from filesystem
    if jobPosition.AttachmentID != nil {
        files, err := s.repo.GetAttachmentFiles(ctx, *jobPosition.AttachmentID)
        if err == nil {
            for _, file := range files {
                if err := os.Remove(file.StoragePath); err != nil {
                    log.Printf("Failed to delete file %s: %v", file.StoragePath, err)
                }
            }
        }

        // Delete attachment
        if err := s.repo.DeleteAttachment(ctx, *jobPosition.AttachmentID); err != nil {
            log.Printf("Failed to delete attachment: %v", err)
        }
    }

    // Delete job position
    return s.repo.DeleteJobPosition(ctx, id)
}

// AddFilesToJobPosition adds files to an existing job position
func (s *jobPositionService) AddFilesToJobPosition(ctx context.Context, jobPositionID int, files []*multipart.FileHeader) error {
    // Get job position
    jobPosition, err := s.repo.GetJobPositionByID(ctx, jobPositionID)
    if err != nil {
        return fmt.Errorf("job position not found: %w", err)
    }

    var attachmentID int

    // Create attachment if doesn't exist
    if jobPosition.AttachmentID == nil {
        newAttachmentID, err := s.repo.CreateAttachment(ctx, 0)
        if err != nil {
            return fmt.Errorf("failed to create attachment: %w", err)
        }
        attachmentID = newAttachmentID

        // Update job position with attachment ID
        jobPosition.AttachmentID = &attachmentID
        if err := s.repo.UpdateJobPosition(ctx, jobPositionID, intModels.JobPosition{
            JobID:        jobPosition.JobID,
            Title:        jobPosition.Title,
            Description:  jobPosition.Description,
            CompanyID:    jobPosition.CompanyID,
            AttachmentID: &attachmentID,
            Status:       jobPosition.Status,
        }); err != nil {
            return fmt.Errorf("failed to update job position with attachment: %w", err)
        }
    } else {
        attachmentID = *jobPosition.AttachmentID
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

// DeleteFileFromJobPosition deletes a file from a job position
func (s *jobPositionService) DeleteFileFromJobPosition(ctx context.Context, fileID int) error {
    // Get file info
    file, err := s.repo.GetFileByID(ctx, fileID)
    if err != nil {
        return fmt.Errorf("file not found: %w", err)
    }

    // Delete file from filesystem
    if err := os.Remove(file.StoragePath); err != nil {
        log.Printf("Failed to delete file from filesystem: %v", err)
    }

    // Delete file from database
    return s.repo.DeleteFile(ctx, fileID)
}

// SearchJobPositions searches job positions by keyword
func (s *jobPositionService) SearchJobPositions(ctx context.Context, keyword string, page, pageSize int) (intModels.JobPositionListResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    jobPositions, total, err := s.repo.SearchJobPositions(ctx, keyword, pageSize, offset)
    if err != nil {
        return intModels.JobPositionListResponse{}, err
    }

    totalPages := (total + pageSize - 1) / pageSize

    return intModels.JobPositionListResponse{
        JobPositions: jobPositions,
        Total:        total,
        Page:         page,
        PageSize:     pageSize,
        TotalPages:   totalPages,
    }, nil
}

// uploadFile handles file upload to filesystem
func (s *jobPositionService) uploadFile(ctx context.Context, attachmentID int, fileHeader *multipart.FileHeader) error {
    // Open uploaded file
    file, err := fileHeader.Open()
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    // Create upload directory if not exists
    uploadDir := "/app/uploads/documents"
    if err := os.MkdirAll(uploadDir, 0755); err != nil {
        return fmt.Errorf("failed to create upload directory: %w", err)
    }

    // Generate unique filename
    timestamp := time.Now().Unix()
    ext := filepath.Ext(fileHeader.Filename)
    nameWithoutExt := strings.TrimSuffix(fileHeader.Filename, ext)
    filename := fmt.Sprintf("%s_%d%s", nameWithoutExt, timestamp, ext)
    filePath := filepath.Join(uploadDir, filename)

    // Create destination file
    dst, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer dst.Close()

    // Copy file content
    if _, err := io.Copy(dst, file); err != nil {
        return fmt.Errorf("failed to copy file: %w", err)
    }

    // Get file type
    fileType := fileHeader.Header.Get("Content-Type")

    // Save file info to database
    fileDetail := intModels.AttachmentDetail{
        FileName:    fileHeader.Filename,
        FileType:    &fileType,
        StoragePath: filePath,
    }

    _, err = s.repo.AddFileToAttachment(ctx, attachmentID, fileDetail)
    if err != nil {
        // Cleanup file if database insert fails
        os.Remove(filePath)
        return fmt.Errorf("failed to save file info: %w", err)
    }

    log.Printf("File uploaded successfully: %s", filename)
    return nil
}