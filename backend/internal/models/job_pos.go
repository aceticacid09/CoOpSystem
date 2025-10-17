package models

import "time"

// JobPosition represents the Job_Position table
type JobPosition struct {
    JobID        int       `db:"job_id" json:"job_id"`
    Title        string    `db:"title" json:"title"`
    Description  *string   `db:"description" json:"description,omitempty"`
    CompanyID    *int      `db:"company_id" json:"company_id,omitempty"`
    AttachmentID *int      `db:"attachment_id" json:"attachment_id,omitempty"`
    Status       string    `db:"status" json:"status"` // open, closed, pending
    CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

// JobPositionWithDetails includes related information
type JobPositionWithDetails struct {
    JobID        int                `db:"job_id" json:"job_id"`
    Title        string             `db:"title" json:"title"`
    Description  *string            `db:"description" json:"description,omitempty"`
    CompanyID    *int               `db:"company_id" json:"company_id,omitempty"`
    AttachmentID *int               `db:"attachment_id" json:"attachment_id,omitempty"`
    Status       string             `db:"status" json:"status"`
    CreatedAt    time.Time          `db:"created_at" json:"created_at"`
    Company      *CompanyInfo       `json:"company,omitempty"`
    Attachments  []AttachmentDetail `json:"attachments,omitempty"`
}

// CompanyInfo contains basic company information for job positions
type CompanyInfo struct {
    CompanyID    int     `json:"company_id"`
    CompanyName  string  `json:"company_name"`
    Status       string  `json:"status"`
    ProfileImage *string `json:"profile_image,omitempty"`
}

// CreateJobPositionRequest for creating new job positions
type CreateJobPositionRequest struct {
    Title       string  `json:"title" binding:"required"`
    Description *string `json:"description"`
    CompanyID   *int    `json:"company_id"`
    Status      *string `json:"status"` // open, closed, pending
}

// UpdateJobPositionRequest for updating job positions
type UpdateJobPositionRequest struct {
    Title       *string `json:"title"`
    Description *string `json:"description"`
    CompanyID   *int    `json:"company_id"`
    Status      *string `json:"status"`
}

// JobPositionListResponse for paginated list
type JobPositionListResponse struct {
    JobPositions []JobPositionWithDetails `json:"job_positions"`
    Total        int                      `json:"total"`
    Page         int                      `json:"page"`
    PageSize     int                      `json:"page_size"`
    TotalPages   int                      `json:"total_pages"`
}

// JobPositionFilterRequest for filtering job positions
type JobPositionFilterRequest struct {
    Status    *string `json:"status"`
    CompanyID *int    `json:"company_id"`
    Page      int     `json:"page"`
    PageSize  int     `json:"page_size"`
}