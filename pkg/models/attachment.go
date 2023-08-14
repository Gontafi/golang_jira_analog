package models

import "time"

type Attachment struct {
	Id               int
	IssueId          int
	FileName         string
	FileSize         string
	UploadedByUserId int
	UploadedDate     time.Time
}
