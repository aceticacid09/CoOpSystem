package repository

import (
    "context"
    "coop/internal/models"
    "fmt"
    "path/filepath"

    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
)

type publicDocRepository struct {
    db *pgxpool.Pool
}

func NewPublicDocRepository(db *pgxpool.Pool) *publicDocRepository {
    return &publicDocRepository{db: db}
}

// GetPublicDocByID retrieves a public document by ID
func (r *publicDocRepository) GetPublicDocByID(ctx context.Context, id int) (models.PublicDocWithDetails, error) {
    var doc models.PublicDocWithDetails

    err := r.db.QueryRow(ctx, `
        SELECT pdoc_id, title, file_path, teacher_id, uploaded_at
        FROM PublicDoc
        WHERE pdoc_id = $1
    `, id).Scan(
        &doc.PdocID,
        &doc.Title,
        &doc.FilePath,
        &doc.TeacherID,
        &doc.UploadedAt,
    )

    if err != nil {
        if err == pgx.ErrNoRows {
            return doc, fmt.Errorf("public document not found")
        }
        return doc, fmt.Errorf("failed to get public document: %w", err)
    }

    // Extract file name from path
    doc.FileName = filepath.Base(doc.FilePath)
    doc.FileType = filepath.Ext(doc.FilePath)

    // Get teacher info if exists
    if doc.TeacherID != nil {
        teacher, err := r.getTeacherInfo(ctx, *doc.TeacherID)
        if err == nil {
            doc.Teacher = teacher
        }
    }

    return doc, nil
}

// GetAllPublicDocs retrieves all public documents with pagination
func (r *publicDocRepository) GetAllPublicDocs(ctx context.Context, limit, offset int) ([]models.PublicDocWithDetails, int, error) {
    // Get total count
    var total int
    err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM PublicDoc`).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count public documents: %w", err)
    }

    // Get documents
    rows, err := r.db.Query(ctx, `
        SELECT pdoc_id, title, file_path, teacher_id, uploaded_at
        FROM PublicDoc
        ORDER BY uploaded_at DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)

    if err != nil {
        return nil, 0, fmt.Errorf("failed to get public documents: %w", err)
    }
    defer rows.Close()

    var documents []models.PublicDocWithDetails
    for rows.Next() {
        var doc models.PublicDocWithDetails

        err := rows.Scan(
            &doc.PdocID,
            &doc.Title,
            &doc.FilePath,
            &doc.TeacherID,
            &doc.UploadedAt,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("failed to scan public document: %w", err)
        }

        // Extract file name from path
        doc.FileName = filepath.Base(doc.FilePath)
        doc.FileType = filepath.Ext(doc.FilePath)

        // Get teacher info
        if doc.TeacherID != nil {
            teacher, err := r.getTeacherInfo(ctx, *doc.TeacherID)
            if err == nil {
                doc.Teacher = teacher
            }
        }

        documents = append(documents, doc)
    }

    return documents, total, nil
}

// CreatePublicDoc creates a new public document
func (r *publicDocRepository) CreatePublicDoc(ctx context.Context, doc models.PublicDoc) (int, error) {
    var pdocID int
    err := r.db.QueryRow(ctx, `
        INSERT INTO PublicDoc (title, file_path, teacher_id)
        VALUES ($1, $2, $3)
        RETURNING pdoc_id
    `, doc.Title, doc.FilePath, doc.TeacherID).Scan(&pdocID)

    if err != nil {
        return 0, fmt.Errorf("failed to create public document: %w", err)
    }

    return pdocID, nil
}

// UpdatePublicDoc updates a public document
func (r *publicDocRepository) UpdatePublicDoc(ctx context.Context, id int, doc models.PublicDoc) error {
    _, err := r.db.Exec(ctx, `
        UPDATE PublicDoc
        SET title = $1, teacher_id = $2
        WHERE pdoc_id = $3
    `, doc.Title, doc.TeacherID, id)

    if err != nil {
        return fmt.Errorf("failed to update public document: %w", err)
    }

    return nil
}

// DeletePublicDoc deletes a public document
func (r *publicDocRepository) DeletePublicDoc(ctx context.Context, id int) error {
    result, err := r.db.Exec(ctx, `
        DELETE FROM PublicDoc
        WHERE pdoc_id = $1
    `, id)

    if err != nil {
        return fmt.Errorf("failed to delete public document: %w", err)
    }

    if result.RowsAffected() == 0 {
        return fmt.Errorf("public document not found")
    }

    return nil
}

// SearchPublicDocs searches public documents by title
func (r *publicDocRepository) SearchPublicDocs(ctx context.Context, query string, limit, offset int) ([]models.PublicDocWithDetails, int, error) {
    searchQuery := "%" + query + "%"

    // Get total count
    var total int
    err := r.db.QueryRow(ctx, `
        SELECT COUNT(*) FROM PublicDoc
        WHERE title ILIKE $1
    `, searchQuery).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count search results: %w", err)
    }

    // Get documents
    rows, err := r.db.Query(ctx, `
        SELECT pdoc_id, title, file_path, teacher_id, uploaded_at
        FROM PublicDoc
        WHERE title ILIKE $1
        ORDER BY uploaded_at DESC
        LIMIT $2 OFFSET $3
    `, searchQuery, limit, offset)

    if err != nil {
        return nil, 0, fmt.Errorf("failed to search public documents: %w", err)
    }
    defer rows.Close()

    var documents []models.PublicDocWithDetails
    for rows.Next() {
        var doc models.PublicDocWithDetails

        err := rows.Scan(
            &doc.PdocID,
            &doc.Title,
            &doc.FilePath,
            &doc.TeacherID,
            &doc.UploadedAt,
        )
        if err != nil {
            return nil, 0, fmt.Errorf("failed to scan public document: %w", err)
        }

        // Extract file name from path
        doc.FileName = filepath.Base(doc.FilePath)
        doc.FileType = filepath.Ext(doc.FilePath)

        // Get teacher info
        if doc.TeacherID != nil {
            teacher, err := r.getTeacherInfo(ctx, *doc.TeacherID)
            if err == nil {
                doc.Teacher = teacher
            }
        }

        documents = append(documents, doc)
    }

    return documents, total, nil
}

// getTeacherInfo retrieves teacher information
func (r *publicDocRepository) getTeacherInfo(ctx context.Context, teacherID int) (*models.TeacherInfo, error) {
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
        return nil, fmt.Errorf("failed to get teacher info: %w", err)
    }

    return &teacher, nil
}