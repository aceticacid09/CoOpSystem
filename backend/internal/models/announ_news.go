package models

import "time"

// Announcement represents the Announ_News table
type Announcement struct {
    PostID       int       `db:"post_id" json:"post_id"`
    Title        string    `db:"title" json:"title"`
    Description  *string   `db:"description" json:"description,omitempty"`
    AttachmentID *int      `db:"attachment_id" json:"attachment_id,omitempty"`
    TeacherID    *int      `db:"teacher_id" json:"teacher_id,omitempty"`
    CreatedAt    time.Time `db:"created_at" json:"created_at"`
    UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// AnnouncementWithDetails includes related information
type AnnouncementWithDetails struct {
    PostID       int                `db:"post_id" json:"post_id"`
    Title        string             `db:"title" json:"title"`
    Description  *string            `db:"description" json:"description,omitempty"`
    AttachmentID *int               `db:"attachment_id" json:"attachment_id,omitempty"`
    TeacherID    *int               `db:"teacher_id" json:"teacher_id,omitempty"`
    CreatedAt    time.Time          `db:"created_at" json:"created_at"`
    UpdatedAt    time.Time          `db:"updated_at" json:"updated_at"`
    Teacher      *TeacherInfo       `json:"teacher,omitempty"`
    Attachments  []AttachmentDetail `json:"attachments,omitempty"`
}

// TeacherInfo contains basic teacher information for announcements
type TeacherInfo struct {
    TeacherID    int     `json:"teacher_id"`
    NameTH       *string `json:"name_th,omitempty"`
    NameEN       *string `json:"name_en,omitempty"`
    EmailContact *string `json:"email_contact,omitempty"`
    ProfileImage *string `json:"profile_image,omitempty"`
}

// CreateAnnouncementRequest for creating new announcements
type CreateAnnouncementRequest struct {
    Title       string  `json:"title" binding:"required"`
    Description *string `json:"description"`
    TeacherID   *int    `json:"teacher_id"`
}

// UpdateAnnouncementRequest for updating announcements
type UpdateAnnouncementRequest struct {
    Title       *string `json:"title"`
    Description *string `json:"description"`
    TeacherID   *int    `json:"teacher_id"`
}

// AnnouncementListResponse for paginated list
type AnnouncementListResponse struct {
    Announcements []AnnouncementWithDetails `json:"announcements"`
    Total         int                       `json:"total"`
    Page          int                       `json:"page"`
    PageSize      int                		`json:"page_size"`
    TotalPages    int                       `json:"total_pages"`
}