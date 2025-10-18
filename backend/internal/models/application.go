package models

import "time"

// Application represents the Application table
type Application struct {
	ApplicationID int       `db:"application_id" json:"application_id"`
	JobID         int       `db:"job_id" json:"job_id"`
	StudentID     int       `db:"student_id" json:"student_id"`
	SDocID        *int      `db:"sdoc_id" json:"sdoc_id,omitempty"`
	Status        string    `db:"status" json:"status"` // pending, approved, rejected
	AppliedAt     time.Time `db:"applied_at" json:"applied_at"`
}

// ApplicationWithDetails includes related information
type ApplicationWithDetails struct {
	ApplicationID int          `db:"application_id" json:"application_id"`
	JobID         int          `db:"job_id" json:"job_id"`
	StudentID     int          `db:"student_id" json:"student_id"`
	SDocID        *int         `db:"sdoc_id" json:"sdoc_id,omitempty"`
	Status        string       `db:"status" json:"status"`
	AppliedAt     time.Time    `db:"applied_at" json:"applied_at"`
	Job           *JobPosition `json:"job,omitempty"`
	Student       *StudentInfo `json:"student,omitempty"`
	Documents     []StudentDoc `json:"documents,omitempty"`
}

// StudentInfo contains basic student information for applications
type StudentInfo struct {
	StudentID    int     `json:"student_id"`
	NameTH       string  `json:"name_th"`
	NameEN       string  `json:"name_en"`
	StudentCode  int     `json:"student_code"`
	Year         int     `json:"year"`
	ProfileImage *string `json:"profile_image,omitempty"`
}

// CreateApplicationRequest for creating new applications
type CreateApplicationRequest struct {
	JobID     int  `json:"job_id" binding:"required"`
	StudentID int  `json:"student_id" binding:"required"`
	SDocID    *int `json:"sdoc_id"`
}

// UpdateApplicationRequest for updating applications
type UpdateApplicationRequest struct {
	Status *string `json:"status"`
	SDocID *int    `json:"sdoc_id"`
}

// ApplicationListResponse for paginated list
type ApplicationListResponse struct {
	Applications []ApplicationWithDetails `json:"applications"`
	Total        int                      `json:"total"`
	Page         int                      `json:"page"`
	PageSize     int                      `json:"page_size"`
	TotalPages   int                      `json:"total_pages"`
}

// ApplicationFilterRequest for filtering applications
type ApplicationFilterRequest struct {
	Status    *string `json:"status"`
	JobID     *int    `json:"job_id"`
	StudentID *int    `json:"student_id"`
	Page      int     `json:"page"`
	PageSize  int     `json:"page_size"`
}
