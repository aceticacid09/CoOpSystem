package models

import "time"

// PublicDoc represents the PublicDoc table
type PublicDoc struct {
    PdocID     int       `db:"pdoc_id" json:"pdoc_id"`
    Title      string    `db:"title" json:"title"`
    FilePath   string    `db:"file_path" json:"file_path"`
    TeacherID  *int      `db:"teacher_id" json:"teacher_id,omitempty"`
    UploadedAt time.Time `db:"uploaded_at" json:"uploaded_at"`
}

// PublicDocWithDetails includes teacher information
type PublicDocWithDetails struct {
    PdocID     int          `db:"pdoc_id" json:"pdoc_id"`
    Title      string       `db:"title" json:"title"`
    FilePath   string       `db:"file_path" json:"file_path"`
    FileName   string       `json:"file_name"`
    FileType   string       `json:"file_type,omitempty"`
    FileSize   int64        `json:"file_size,omitempty"`
    TeacherID  *int         `db:"teacher_id" json:"teacher_id,omitempty"`
    UploadedAt time.Time    `db:"uploaded_at" json:"uploaded_at"`
    Teacher    *TeacherInfo `json:"teacher,omitempty"`
}

// CreatePublicDocRequest for creating new public documents
type CreatePublicDocRequest struct {
    Title     string `json:"title" binding:"required"`
    TeacherID *int   `json:"teacher_id"`
}

// UpdatePublicDocRequest for updating public documents
type UpdatePublicDocRequest struct {
    Title     *string `json:"title"`
    TeacherID *int    `json:"teacher_id"`
}

// PublicDocListResponse for paginated list
type PublicDocListResponse struct {
    Documents  []PublicDocWithDetails `json:"documents"`
    Total      int                    `json:"total"`
    Page       int                    `json:"page"`
    PageSize   int                    `json:"page_size"`
    TotalPages int                    `json:"total_pages"`
}