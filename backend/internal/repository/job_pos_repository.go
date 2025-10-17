package repository

import (
    "context"
    "coop/internal/models"
    "fmt"

    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
)

type jobPositionRepository struct {
    db *pgxpool.Pool
}

func NewJobPositionRepository(db *pgxpool.Pool) *jobPositionRepository {
    return &jobPositionRepository{db: db}
}

// GetJobPositionByID retrieves a job position with details by ID
func (r *jobPositionRepository) GetJobPositionByID(ctx context.Context, id int) (models.JobPositionWithDetails, error) {
    var jobPosition models.JobPositionWithDetails

    // Get job position
    err := r.db.QueryRow(ctx, `
        SELECT job_id, title, description, company_id, attachment_id, status, created_at
        FROM Job_Position
        WHERE job_id = $1
    `, id).Scan(
        &jobPosition.JobID,
        &jobPosition.Title,
        &jobPosition.Description,
        &jobPosition.CompanyID,
        &jobPosition.AttachmentID,
        &jobPosition.Status,
        &jobPosition.CreatedAt,
    )

    if err != nil {
        if err == pgx.ErrNoRows {
            return jobPosition, fmt.Errorf("job position not found")
        }
        return jobPosition, fmt.Errorf("failed to get job position: %w", err)
    }

    // Get company info if exists
    if jobPosition.CompanyID != nil {
        company, err := r.getCompanyInfo(ctx, *jobPosition.CompanyID)
        if err == nil {
            jobPosition.Company = company
        }
    }

    // Get attachments if exists
    if jobPosition.AttachmentID != nil {
        attachments, err := r.GetAttachmentFiles(ctx, *jobPosition.AttachmentID)
        if err == nil {
            jobPosition.Attachments = attachments
        }
    }

    return jobPosition, nil
}

// GetAllJobPositions retrieves all job positions with pagination
func (r *jobPositionRepository) GetAllJobPositions(ctx context.Context, limit, offset int) ([]models.JobPositionWithDetails, int, error) {
    // Get total count
    var total int
    err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM Job_Position`).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count job positions: %w", err)
    }

    // Get job positions
    rows, err := r.db.Query(ctx, `
        SELECT job_id, title, description, company_id, attachment_id, status, created_at
        FROM Job_Position
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)

    if err != nil {
        return nil, 0, fmt.Errorf("failed to get job positions: %w", err)
    }
    defer rows.Close()

    var jobPositions []models.JobPositionWithDetails
    for rows.Next() {
        var jp models.JobPositionWithDetails

        err := rows.Scan(
            &jp.JobID,
            &jp.Title,
            &jp.Description,
            &jp.CompanyID,
            &jp.AttachmentID,
            &jp.Status,
            &jp.CreatedAt,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("failed to scan job position: %w", err)
        }

        // Get company info
        if jp.CompanyID != nil {
            company, err := r.getCompanyInfo(ctx, *jp.CompanyID)
            if err == nil {
                jp.Company = company
            }
        }

        // Get attachments
        if jp.AttachmentID != nil {
            attachments, err := r.GetAttachmentFiles(ctx, *jp.AttachmentID)
            if err == nil {
                jp.Attachments = attachments
            }
        }

        jobPositions = append(jobPositions, jp)
    }

    return jobPositions, total, nil
}

// GetJobPositionsByStatus retrieves job positions filtered by status
func (r *jobPositionRepository) GetJobPositionsByStatus(ctx context.Context, status string, limit, offset int) ([]models.JobPositionWithDetails, int, error) {
    // Get total count
    var total int
    err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM Job_Position WHERE status = $1`, status).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count job positions: %w", err)
    }

    // Get job positions
    rows, err := r.db.Query(ctx, `
        SELECT job_id, title, description, company_id, attachment_id, status, created_at
        FROM Job_Position
        WHERE status = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `, status, limit, offset)

    if err != nil {
        return nil, 0, fmt.Errorf("failed to get job positions: %w", err)
    }
    defer rows.Close()

    var jobPositions []models.JobPositionWithDetails
    for rows.Next() {
        var jp models.JobPositionWithDetails

        err := rows.Scan(
            &jp.JobID,
            &jp.Title,
            &jp.Description,
            &jp.CompanyID,
            &jp.AttachmentID,
            &jp.Status,
            &jp.CreatedAt,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("failed to scan job position: %w", err)
        }

        // Get company info
        if jp.CompanyID != nil {
            company, err := r.getCompanyInfo(ctx, *jp.CompanyID)
            if err == nil {
                jp.Company = company
            }
        }

        // Get attachments
        if jp.AttachmentID != nil {
            attachments, err := r.GetAttachmentFiles(ctx, *jp.AttachmentID)
            if err == nil {
                jp.Attachments = attachments
            }
        }

        jobPositions = append(jobPositions, jp)
    }

    return jobPositions, total, nil
}

// GetJobPositionsByCompany retrieves job positions filtered by company
func (r *jobPositionRepository) GetJobPositionsByCompany(ctx context.Context, companyID int, limit, offset int) ([]models.JobPositionWithDetails, int, error) {
    // Get total count
    var total int
    err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM Job_Position WHERE company_id = $1`, companyID).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count job positions: %w", err)
    }

    // Get job positions
    rows, err := r.db.Query(ctx, `
        SELECT job_id, title, description, company_id, attachment_id, status, created_at
        FROM Job_Position
        WHERE company_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `, companyID, limit, offset)

    if err != nil {
        return nil, 0, fmt.Errorf("failed to get job positions: %w", err)
    }
    defer rows.Close()

    var jobPositions []models.JobPositionWithDetails
    for rows.Next() {
        var jp models.JobPositionWithDetails

        err := rows.Scan(
            &jp.JobID,
            &jp.Title,
            &jp.Description,
            &jp.CompanyID,
            &jp.AttachmentID,
            &jp.Status,
            &jp.CreatedAt,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("failed to scan job position: %w", err)
        }

        // Get company info
        if jp.CompanyID != nil {
            company, err := r.getCompanyInfo(ctx, *jp.CompanyID)
            if err == nil {
                jp.Company = company
            }
        }

        // Get attachments
        if jp.AttachmentID != nil {
            attachments, err := r.GetAttachmentFiles(ctx, *jp.AttachmentID)
            if err == nil {
                jp.Attachments = attachments
            }
        }

        jobPositions = append(jobPositions, jp)
    }

    return jobPositions, total, nil
}

// CreateJobPosition creates a new job position
func (r *jobPositionRepository) CreateJobPosition(ctx context.Context, jobPosition models.JobPosition) (int, error) {
    var jobID int
    err := r.db.QueryRow(ctx, `
        INSERT INTO Job_Position (title, description, company_id, attachment_id, status)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING job_id
    `, jobPosition.Title, jobPosition.Description, jobPosition.CompanyID, jobPosition.AttachmentID, jobPosition.Status).Scan(&jobID)

    if err != nil {
        return 0, fmt.Errorf("failed to create job position: %w", err)
    }

    return jobID, nil
}

// UpdateJobPosition updates a job position
func (r *jobPositionRepository) UpdateJobPosition(ctx context.Context, id int, jobPosition models.JobPosition) error {
    _, err := r.db.Exec(ctx, `
        UPDATE Job_Position
        SET title = $1, description = $2, company_id = $3, status = $4
        WHERE job_id = $5
    `, jobPosition.Title, jobPosition.Description, jobPosition.CompanyID, jobPosition.Status, id)

    if err != nil {
        return fmt.Errorf("failed to update job position: %w", err)
    }

    return nil
}

// DeleteJobPosition deletes a job position
func (r *jobPositionRepository) DeleteJobPosition(ctx context.Context, id int) error {
    _, err := r.db.Exec(ctx, `DELETE FROM Job_Position WHERE job_id = $1`, id)
    if err != nil {
        return fmt.Errorf("failed to delete job position: %w", err)
    }
    return nil
}

// SearchJobPositions searches job positions by keyword
func (r *jobPositionRepository) SearchJobPositions(ctx context.Context, keyword string, limit, offset int) ([]models.JobPositionWithDetails, int, error) {
    searchPattern := "%" + keyword + "%"

    // Get total count
    var total int
    err := r.db.QueryRow(ctx, `
        SELECT COUNT(*) 
        FROM Job_Position 
        WHERE title ILIKE $1 OR description ILIKE $1
    `, searchPattern).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count search results: %w", err)
    }

    // Get job positions
    rows, err := r.db.Query(ctx, `
        SELECT job_id, title, description, company_id, attachment_id, status, created_at
        FROM Job_Position
        WHERE title ILIKE $1 OR description ILIKE $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `, searchPattern, limit, offset)

    if err != nil {
        return nil, 0, fmt.Errorf("failed to search job positions: %w", err)
    }
    defer rows.Close()

    var jobPositions []models.JobPositionWithDetails
    for rows.Next() {
        var jp models.JobPositionWithDetails

        err := rows.Scan(
            &jp.JobID,
            &jp.Title,
            &jp.Description,
            &jp.CompanyID,
            &jp.AttachmentID,
            &jp.Status,
            &jp.CreatedAt,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("failed to scan job position: %w", err)
        }

        // Get company info
        if jp.CompanyID != nil {
            company, err := r.getCompanyInfo(ctx, *jp.CompanyID)
            if err == nil {
                jp.Company = company
            }
        }

        // Get attachments
        if jp.AttachmentID != nil {
            attachments, err := r.GetAttachmentFiles(ctx, *jp.AttachmentID)
            if err == nil {
                jp.Attachments = attachments
            }
        }

        jobPositions = append(jobPositions, jp)
    }

    return jobPositions, total, nil
}

// Helper function to get company info
func (r *jobPositionRepository) getCompanyInfo(ctx context.Context, companyID int) (*models.CompanyInfo, error) {
    var company models.CompanyInfo
    err := r.db.QueryRow(ctx, `
        SELECT company_id, company_name, status, profile_image
        FROM Company
        WHERE company_id = $1
    `, companyID).Scan(
        &company.CompanyID,
        &company.CompanyName,
        &company.Status,
        &company.ProfileImage,
    )

    if err != nil {
        return nil, fmt.Errorf("failed to get company info: %w", err)
    }

    return &company, nil
}

// CreateAttachment creates a new attachment record
func (r *jobPositionRepository) CreateAttachment(ctx context.Context, attachmentID int) (int, error) {
    var newAttachmentID int
    err := r.db.QueryRow(ctx, `
        INSERT INTO Attachment (file_count)
        VALUES (0)
        RETURNING attachment_id
    `).Scan(&newAttachmentID)

    if err != nil {
        return 0, fmt.Errorf("failed to create attachment: %w", err)
    }

    return newAttachmentID, nil
}

// AddFileToAttachment adds a file to an attachment
func (r *jobPositionRepository) AddFileToAttachment(ctx context.Context, attachmentID int, file models.AttachmentDetail) (int, error) {
    tx, err := r.db.Begin(ctx)
    if err != nil {
        return 0, fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback(ctx)

    // Insert file
    var fileID int
    err = tx.QueryRow(ctx, `
        INSERT INTO File (storage_path, file_name, file_type)
        VALUES ($1, $2, $3)
        RETURNING file_id
    `, file.StoragePath, file.FileName, file.FileType).Scan(&fileID)

    if err != nil {
        return 0, fmt.Errorf("failed to insert file: %w", err)
    }

    // Link file to attachment
    _, err = tx.Exec(ctx, `
        INSERT INTO Attachment_File (attachment_id, file_id)
        VALUES ($1, $2)
    `, attachmentID, fileID)

    if err != nil {
        return 0, fmt.Errorf("failed to link file to attachment: %w", err)
    }

    // Update file count
    _, err = tx.Exec(ctx, `
        UPDATE Attachment
        SET file_count = file_count + 1
        WHERE attachment_id = $1
    `, attachmentID)

    if err != nil {
        return 0, fmt.Errorf("failed to update file count: %w", err)
    }

    if err = tx.Commit(ctx); err != nil {
        return 0, fmt.Errorf("failed to commit transaction: %w", err)
    }

    return fileID, nil
}

// GetAttachmentFiles retrieves all files for an attachment
func (r *jobPositionRepository) GetAttachmentFiles(ctx context.Context, attachmentID int) ([]models.AttachmentDetail, error) {
    rows, err := r.db.Query(ctx, `
        SELECT f.file_id, f.file_name, f.file_type, f.storage_path, f.uploaded_at
        FROM File f
        INNER JOIN Attachment_File af ON f.file_id = af.file_id
        WHERE af.attachment_id = $1
        ORDER BY f.uploaded_at DESC
    `, attachmentID)

    if err != nil {
        return nil, fmt.Errorf("failed to get attachment files: %w", err)
    }
    defer rows.Close()

    var files []models.AttachmentDetail
    for rows.Next() {
        var file models.AttachmentDetail
        err := rows.Scan(
            &file.FileID,
            &file.FileName,
            &file.FileType,
            &file.StoragePath,
            &file.UploadedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to scan file: %w", err)
        }
        files = append(files, file)
    }

    return files, nil
}

// DeleteAttachment deletes an attachment and its files
func (r *jobPositionRepository) DeleteAttachment(ctx context.Context, attachmentID int) error {
    // Get all files for this attachment
    files, err := r.GetAttachmentFiles(ctx, attachmentID)
    if err != nil {
        return fmt.Errorf("failed to get attachment files: %w", err)
    }

    // Delete attachment (cascade will delete Attachment_File entries)
    _, err = r.db.Exec(ctx, `DELETE FROM Attachment WHERE attachment_id = $1`, attachmentID)
    if err != nil {
        return fmt.Errorf("failed to delete attachment: %w", err)
    }

    // Delete orphaned files
    for _, file := range files {
        _, err = r.db.Exec(ctx, `DELETE FROM File WHERE file_id = $1`, file.FileID)
        if err != nil {
            // Log error but continue
            fmt.Printf("failed to delete file %d: %v\n", file.FileID, err)
        }
    }

    return nil
}

// GetFileByID retrieves a file by ID
func (r *jobPositionRepository) GetFileByID(ctx context.Context, fileID int) (models.AttachmentDetail, error) {
    var file models.AttachmentDetail
    err := r.db.QueryRow(ctx, `
        SELECT file_id, file_name, file_type, storage_path, uploaded_at
        FROM File
        WHERE file_id = $1
    `, fileID).Scan(
        &file.FileID,
        &file.FileName,
        &file.FileType,
        &file.StoragePath,
        &file.UploadedAt,
    )

    if err != nil {
        if err == pgx.ErrNoRows {
            return file, fmt.Errorf("file not found")
        }
        return file, fmt.Errorf("failed to get file: %w", err)
    }

    return file, nil
}

// DeleteFile deletes a file
func (r *jobPositionRepository) DeleteFile(ctx context.Context, fileID int) error {
    tx, err := r.db.Begin(ctx)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback(ctx)

    // Get attachment_id before deleting
    var attachmentID int
    err = tx.QueryRow(ctx, `
        SELECT attachment_id FROM Attachment_File WHERE file_id = $1
    `, fileID).Scan(&attachmentID)

    if err != nil && err != pgx.ErrNoRows {
        return fmt.Errorf("failed to get attachment_id: %w", err)
    }

    // Delete from Attachment_File
    _, err = tx.Exec(ctx, `DELETE FROM Attachment_File WHERE file_id = $1`, fileID)
    if err != nil {
        return fmt.Errorf("failed to delete from Attachment_File: %w", err)
    }

    // Delete file
    _, err = tx.Exec(ctx, `DELETE FROM File WHERE file_id = $1`, fileID)
    if err != nil {
        return fmt.Errorf("failed to delete file: %w", err)
    }

    // Update file count if attachment exists
    if attachmentID > 0 {
        _, err = tx.Exec(ctx, `
            UPDATE Attachment
            SET file_count = file_count - 1
            WHERE attachment_id = $1
        `, attachmentID)
        if err != nil {
            return fmt.Errorf("failed to update file count: %w", err)
        }
    }

    if err = tx.Commit(ctx); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}

// GetJobPositionIDByFileID retrieves job position ID by file ID
func (r *jobPositionRepository) GetJobPositionIDByFileID(ctx context.Context, fileID int) (int, error) {
    var jobID int
    err := r.db.QueryRow(ctx, `
        SELECT jp.job_id
        FROM Job_Position jp
        INNER JOIN Attachment_File af ON jp.attachment_id = af.attachment_id
        WHERE af.file_id = $1
    `, fileID).Scan(&jobID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return 0, fmt.Errorf("job position not found for file")
        }
        return 0, fmt.Errorf("failed to get job position ID: %w", err)
    }

    return jobID, nil
}