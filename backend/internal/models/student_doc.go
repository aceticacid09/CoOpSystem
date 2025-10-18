package models

import "time"

// DocType represents the type of student document
type DocType string

// Constants for document types
const (
	DocTypeResume     DocType = "resume"
	DocTypeCV         DocType = "cv"
	DocTypePortfolio  DocType = "portfolio"
	DocTypeTranscript DocType = "transcript"
	DocTypeOther      DocType = "other"
)

// StudentDoc represents a student document in the system
type StudentDoc struct {
	SDocID       int       `json:"sdoc_id" db:"sdoc_id"`
	StudentID    int       `json:"student_id" db:"student_id"`
	DocType      DocType   `json:"doc_type" db:"doc_type"`
	AttachmentID *int      `json:"attachment_id,omitempty" db:"attachment_id"`
	UploadedAt   time.Time `json:"uploaded_at" db:"uploaded_at"`

	// Related data (populated when needed)
	Attachment *Attachment `json:"attachment,omitempty" db:"-"`
}