package models

import "time"

type Attachment struct {
	ID               int       `json:"id"`
	IssueID          int       `json:"issue_id"`
	FileName         string    `json:"file_name"`
	FileSize         int       `json:"file_size"`
	UploadedByUserID int       `json:"uploaded_by_user_id"`
	UploadedDate     time.Time `json:"uploaded_date"`
}
