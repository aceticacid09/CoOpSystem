package repository

import (
	"context"
	"coop/internal/models"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type announcementRepository struct {
	db *pgxpool.Pool
}

func NewAnnouncementRepository(db *pgxpool.Pool) *announcementRepository {
	return &announcementRepository{db: db}
}

// GetAnnouncementByID retrieves an announcement with details by ID
func (r *announcementRepository) GetAnnouncementByID(ctx context.Context, id int) (models.AnnouncementWithDetails, error) {
	var announcement models.AnnouncementWithDetails

	// Get announcement
	err := r.db.QueryRow(ctx, `
        SELECT post_id, title, description, attachment_id, teacher_id, created_at, updated_at
        FROM Announ_News
        WHERE post_id = $1
    `, id).Scan(
		&announcement.PostID,
		&announcement.Title,
		&announcement.Description,
		&announcement.AttachmentID,
		&announcement.TeacherID,
		&announcement.CreatedAt,
		&announcement.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return announcement, fmt.Errorf("announcement not found")
		}
		return announcement, fmt.Errorf("failed to get announcement: %w", err)
	}

	// Get teacher info if exists
	if announcement.TeacherID != nil {
		teacher, err := r.getTeacherInfo(ctx, *announcement.TeacherID)
		if err == nil {
			announcement.Teacher = teacher
		}
	}

	// Get attachments if exists
	if announcement.AttachmentID != nil {
		attachments, err := r.GetAttachmentFiles(ctx, *announcement.AttachmentID)
		if err == nil {
			announcement.Attachments = attachments
		}
	}

	return announcement, nil
}

// GetAllAnnouncements retrieves all announcements with pagination
func (r *announcementRepository) GetAllAnnouncements(ctx context.Context, limit, offset int) ([]models.AnnouncementWithDetails, int, error) {
	// Get total count
	var total int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM Announ_News`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count announcements: %w", err)
	}

	// Get announcements
	rows, err := r.db.Query(ctx, `
        SELECT post_id, title, description, attachment_id, teacher_id, created_at, updated_at
        FROM Announ_News
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get announcements: %w", err)
	}
	defer rows.Close()

	var announcements []models.AnnouncementWithDetails
	for rows.Next() {
		var a models.AnnouncementWithDetails

		err := rows.Scan(
			&a.PostID,
			&a.Title,
			&a.Description,
			&a.AttachmentID,
			&a.TeacherID,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan announcement: %w", err)
		}

		// Get teacher info
		if a.TeacherID != nil {
			teacher, err := r.getTeacherInfo(ctx, *a.TeacherID)
			if err == nil {
				a.Teacher = teacher
			}
		}

		// Get attachments
		if a.AttachmentID != nil {
			attachments, err := r.GetAttachmentFiles(ctx, *a.AttachmentID)
			if err == nil {
				a.Attachments = attachments
			}
		}

		announcements = append(announcements, a)
	}

	return announcements, total, nil
}

// CreateAnnouncement creates a new announcement
func (r *announcementRepository) CreateAnnouncement(ctx context.Context, announcement models.Announcement) (int, error) {
	var postID int
	err := r.db.QueryRow(ctx, `
        INSERT INTO Announ_News (title, description, attachment_id, teacher_id)
        VALUES ($1, $2, $3, $4)
        RETURNING post_id
    `, announcement.Title, announcement.Description, announcement.AttachmentID, announcement.TeacherID).Scan(&postID)

	if err != nil {
		return 0, fmt.Errorf("failed to create announcement: %w", err)
	}

	return postID, nil
}

// UpdateAnnouncement updates an announcement
func (r *announcementRepository) UpdateAnnouncement(ctx context.Context, id int, announcement models.Announcement) error {
	_, err := r.db.Exec(ctx, `
        UPDATE Announ_News
        SET title = $1, description = $2, teacher_id = $3, updated_at = NOW()
        WHERE post_id = $4
    `, announcement.Title, announcement.Description, announcement.TeacherID, id)

	if err != nil {
		return fmt.Errorf("failed to update announcement: %w", err)
	}

	return nil
}

// DeleteAnnouncement deletes an announcement
func (r *announcementRepository) DeleteAnnouncement(ctx context.Context, id int) error {
	// Get attachment_id before deleting
	var attachmentID *int
	err := r.db.QueryRow(ctx, `
        SELECT attachment_id FROM Announ_News WHERE post_id = $1
    `, id).Scan(&attachmentID)

	if err != nil && err != pgx.ErrNoRows {
		return fmt.Errorf("failed to get announcement: %w", err)
	}

	// Delete announcement
	result, err := r.db.Exec(ctx, `DELETE FROM Announ_News WHERE post_id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete announcement: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("announcement not found")
	}

	// Delete associated attachments if exists
	if attachmentID != nil {
		r.DeleteAttachment(ctx, *attachmentID)
	}

	return nil
}

// CreateAttachment creates a new attachment record
func (r *announcementRepository) CreateAttachment(ctx context.Context, fileCount int) (int, error) {
	var attachmentID int
	err := r.db.QueryRow(ctx, `
        INSERT INTO Attachment (file_count)
        VALUES ($1)
        RETURNING attachment_id
    `, fileCount).Scan(&attachmentID)

	if err != nil {
		return 0, fmt.Errorf("failed to create attachment: %w", err)
	}

	return attachmentID, nil
}

// AddFileToAttachment adds a file to an attachment
func (r *announcementRepository) AddFileToAttachment(ctx context.Context, attachmentID int, file models.AttachmentDetail) (int, error) {
	var fileID int

	// Insert file
	err := r.db.QueryRow(ctx, `
        INSERT INTO File (storage_path, file_name, file_type)
        VALUES ($1, $2, $3)
        RETURNING file_id
    `, file.StoragePath, file.FileName, file.FileType).Scan(&fileID)

	if err != nil {
		return 0, fmt.Errorf("failed to create file: %w", err)
	}

	// Update attachment with file_id and increment file_count
	_, err = r.db.Exec(ctx, `
        UPDATE Attachment
        SET file_id = $1, file_count = file_count + 1
        WHERE attachment_id = $2
    `, fileID, attachmentID)

	if err != nil {
		return 0, fmt.Errorf("failed to update attachment: %w", err)
	}

	return fileID, nil
}

// GetAttachmentFiles retrieves all files for an attachment
func (r *announcementRepository) GetAttachmentFiles(ctx context.Context, attachmentID int) ([]models.AttachmentDetail, error) {
	rows, err := r.db.Query(ctx, `
        SELECT f.file_id, f.file_name, f.file_type, f.storage_path, f.uploaded_at
        FROM File f
        INNER JOIN Attachment a ON f.file_id = a.file_id
        WHERE a.attachment_id = $1
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
func (r *announcementRepository) DeleteAttachment(ctx context.Context, attachmentID int) error {
	// Get file_id
	var fileID *int
	err := r.db.QueryRow(ctx, `
        SELECT file_id FROM Attachment WHERE attachment_id = $1
    `, attachmentID).Scan(&fileID)

	if err != nil && err != pgx.ErrNoRows {
		return fmt.Errorf("failed to get attachment: %w", err)
	}

	// Delete attachment
	_, err = r.db.Exec(ctx, `DELETE FROM Attachment WHERE attachment_id = $1`, attachmentID)
	if err != nil {
		return fmt.Errorf("failed to delete attachment: %w", err)
	}

	// Delete file if exists
	if fileID != nil {
		_, err = r.db.Exec(ctx, `DELETE FROM File WHERE file_id = $1`, *fileID)
		if err != nil {
			return fmt.Errorf("failed to delete file: %w", err)
		}

	}

	return nil
}

// SearchAnnouncements searches announcements by keyword
func (r *announcementRepository) SearchAnnouncements(ctx context.Context, keyword string, limit, offset int) ([]models.AnnouncementWithDetails, int, error) {
	searchPattern := "%" + keyword + "%"

	// Get total count
	var total int
	err := r.db.QueryRow(ctx, `
        SELECT COUNT(*) FROM Announ_News
        WHERE title ILIKE $1 OR description ILIKE $1
    `, searchPattern).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count announcements: %w", err)
	}

	// Get announcements
	rows, err := r.db.Query(ctx, `
        SELECT post_id, title, description, attachment_id, teacher_id, created_at, updated_at
        FROM Announ_News
        WHERE title ILIKE $1 OR description ILIKE $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `, searchPattern, limit, offset)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to search announcements: %w", err)
	}
	defer rows.Close()

	var announcements []models.AnnouncementWithDetails
	for rows.Next() {

		var a models.AnnouncementWithDetails
		err := rows.Scan(
			&a.PostID,
			&a.Title,
			&a.Description,
			&a.AttachmentID,
			&a.TeacherID,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan announcement: %w", err)
		}

		// Get teacher info
		if a.TeacherID != nil {
			teacher, err := r.getTeacherInfo(ctx, *a.TeacherID)
			if err == nil {
				a.Teacher = teacher
			}
		}

		// Get attachments
		if a.AttachmentID != nil {
			attachments, err := r.GetAttachmentFiles(ctx, *a.AttachmentID)
			if err == nil {
				a.Attachments = attachments
			}
		}

		announcements = append(announcements, a)
	}

	return announcements, total, nil
}

// getTeacherInfo is a helper function to get teacher information
func (r *announcementRepository) getTeacherInfo(ctx context.Context, teacherID int) (*models.TeacherInfo, error) {
	var teacher models.TeacherInfo
	err := r.db.QueryRow(ctx, `
        SELECT teacher_id, name_th, name_en, email_contact, profile_image
        FROM Teacher
        WHERE teacher_id = $1
    `, teacherID).Scan(
		&teacher.TeacherID,
		&teacher.NameTH,
		&teacher.NameEN,
		&teacher.EmailContact,
		&teacher.ProfileImage,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get teacher info: %w", err)
	}

	return &teacher, nil
}
