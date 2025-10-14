package service

import (
    "context"
    "coop/internal/interfaces"
    "coop/internal/models"
    "fmt"
    "io"
    "math"
    "mime/multipart"
    "os"
    "path/filepath"
    "strings"
    "time"
)

type publicDocService struct {
    repo       interfaces.PublicDocRepository
    uploadPath string
}

func NewPublicDocService(repo interfaces.PublicDocRepository) *publicDocService {
    return &publicDocService{
        repo:       repo,
        uploadPath: "/app/uploads/public_docs",
    }
}

// GetPublicDocByID retrieves a public document by ID
func (s *publicDocService) GetPublicDocByID(ctx context.Context, id int) (models.PublicDocWithDetails, error) {
    return s.repo.GetPublicDocByID(ctx, id)
}

// GetAllPublicDocs retrieves all public documents with pagination
func (s *publicDocService) GetAllPublicDocs(ctx context.Context, page, pageSize int) (models.PublicDocListResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    documents, total, err := s.repo.GetAllPublicDocs(ctx, pageSize, offset)
    if err != nil {
        return models.PublicDocListResponse{}, err
    }

    totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

    return models.PublicDocListResponse{
        Documents:  documents,
        Total:      total,
        Page:       page,
        PageSize:   pageSize,
        TotalPages: totalPages,
    }, nil
}

// CreatePublicDoc creates a new public document with file upload
func (s *publicDocService) CreatePublicDoc(ctx context.Context, req models.CreatePublicDocRequest, file *multipart.FileHeader) (models.PublicDocWithDetails, error) {
    if file == nil {
        return models.PublicDocWithDetails{}, fmt.Errorf("file is required")
    }

    // Create upload directory if it doesn't exist
    if err := os.MkdirAll(s.uploadPath, 0755); err != nil {
        return models.PublicDocWithDetails{}, fmt.Errorf("failed to create upload directory: %w", err)
    }

    // Generate unique filename
    timestamp := time.Now().Unix()
    filename := fmt.Sprintf("%d_%s", timestamp, file.Filename)
    filePath := filepath.Join(s.uploadPath, filename)

    // Save file
    src, err := file.Open()
    if err != nil {
        return models.PublicDocWithDetails{}, fmt.Errorf("failed to open uploaded file: %w", err)
    }
    defer src.Close()

    dst, err := os.Create(filePath)
    if err != nil {
        return models.PublicDocWithDetails{}, fmt.Errorf("failed to create file: %w", err)
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return models.PublicDocWithDetails{}, fmt.Errorf("failed to save file: %w", err)
    }

    // Create database record
    doc := models.PublicDoc{
        Title:     req.Title,
        FilePath:  filePath,
        TeacherID: req.TeacherID,
    }

    pdocID, err := s.repo.CreatePublicDoc(ctx, doc)
    if err != nil {
        // Clean up file if database insert fails
        os.Remove(filePath)
        return models.PublicDocWithDetails{}, err
    }

    // Get the created document with details
    return s.repo.GetPublicDocByID(ctx, pdocID)
}

// UpdatePublicDoc updates a public document
func (s *publicDocService) UpdatePublicDoc(ctx context.Context, id int, req models.UpdatePublicDocRequest) (models.PublicDocWithDetails, error) {
    // Get existing document
    existing, err := s.repo.GetPublicDocByID(ctx, id)
    if err != nil {
        return models.PublicDocWithDetails{}, err
    }

    // Update fields
    doc := models.PublicDoc{
        Title:     existing.Title,
        FilePath:  existing.FilePath,
        TeacherID: existing.TeacherID,
    }

    if req.Title != nil {
        doc.Title = *req.Title
    }
    if req.TeacherID != nil {
        doc.TeacherID = req.TeacherID
    }

    if err := s.repo.UpdatePublicDoc(ctx, id, doc); err != nil {
        return models.PublicDocWithDetails{}, err
    }

    return s.repo.GetPublicDocByID(ctx, id)
}

// DeletePublicDoc deletes a public document and its file
func (s *publicDocService) DeletePublicDoc(ctx context.Context, id int) error {
    // Get document to get file path
    doc, err := s.repo.GetPublicDocByID(ctx, id)
    if err != nil {
        return err
    }

    // Delete from database
    if err := s.repo.DeletePublicDoc(ctx, id); err != nil {
        return err
    }

    // Delete file from filesystem
    if err := os.Remove(doc.FilePath); err != nil {
        // Log error but don't fail the operation
        fmt.Printf("Warning: failed to delete file %s: %v\n", doc.FilePath, err)
    }

    return nil
}

// SearchPublicDocs searches public documents by title
func (s *publicDocService) SearchPublicDocs(ctx context.Context, query string, page, pageSize int) (models.PublicDocListResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    documents, total, err := s.repo.SearchPublicDocs(ctx, query, pageSize, offset)
    if err != nil {
        return models.PublicDocListResponse{}, err
    }

    totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

    return models.PublicDocListResponse{
        Documents:  documents,
        Total:      total,
        Page:       page,
        PageSize:   pageSize,
        TotalPages: totalPages,
    }, nil
}

// DownloadPublicDoc returns the file path for downloading
func (s *publicDocService) DownloadPublicDoc(ctx context.Context, id int) (string, error) {
    doc, err := s.repo.GetPublicDocByID(ctx, id)
    if err != nil {
        return "", err
    }

    var filePath string
    
    // Handle different path formats
    if strings.HasPrefix(doc.FilePath, "/uploads/public_docs/") {
        fileName := strings.TrimPrefix(doc.FilePath, "/uploads/public_docs/")
        filePath = filepath.Join(s.uploadPath, fileName)
    } else if strings.HasPrefix(doc.FilePath, "/app/uploads/public_docs/") {
        // Already has the correct prefix
        filePath = doc.FilePath
    } else if filepath.IsAbs(doc.FilePath) {
        // Some other absolute path - use as is
        filePath = doc.FilePath
    } else {
        // Relative path - just the filename
        filePath = filepath.Join(s.uploadPath, doc.FilePath)
    }

    // Check if file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return "", fmt.Errorf("file not found on server: %s (original path: %s)", filePath, doc.FilePath)
    }

    return filePath, nil
}