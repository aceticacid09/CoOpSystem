package models

import "time"

// Attachment represents the file attachment structure
type Attachment struct {
	AttachmentID int    `json:"attachment_id" db:"attachment_id"`
	FileCount    int    `json:"file_count" db:"file_count"`
	Files        []File `json:"files,omitempty" db:"-"`
}

// File represents a file in the system
type File struct {
	FileID      int       `json:"file_id" db:"file_id"`
	StoragePath string    `json:"storage_path" db:"storage_path"`
	FileName    string    `json:"file_name" db:"file_name"`
	FileType    string    `json:"file_type" db:"file_type"`
	UploadedAt  time.Time `json:"uploaded_at" db:"uploaded_at"`
}

// AttachmentDetail contains file information
type AttachmentDetail struct {
	FileID      int       `db:"file_id" json:"file_id"`
	FileName    string    `db:"file_name" json:"file_name"`
	FileType    *string   `db:"file_type" json:"file_type,omitempty"`
	StoragePath string    `db:"storage_path" json:"storage_path"`
	UploadedAt  time.Time `db:"uploaded_at" json:"uploaded_at"`
}
