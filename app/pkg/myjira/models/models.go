package models

import "time"

type Attachment struct {
	ID               int
	IssueID          int
	FileName         string
	FileSize         int
	UploadedByUserID int
	UploadedDate     time.Time
}

type Comment struct {
	ID          int
	UserID      int
	IssueID     int
	CommentText string
	CreatedAt   time.Time
}

type Contributors struct {
	ID        int
	UserID    int
	ProjectID int
}

type Issue struct {
	ID               int
	ProjectID        int
	IssueTypeID      int
	IssueSummary     string
	IssueDescription string
	ReporterID       int
	AssigneeID       int
	StageId          int
	StatusID         int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ResolverAt       time.Time
}

type Tag struct {
	ID   int
	Name string
}

type IssueType struct {
	ID          int
	Name        string
	Description string
}

type Priority struct {
	ID          int
	Name        string
	Description string
}

type Status struct {
	ID          int
	Name        string
	Description string
}

type Project struct {
	ID               int
	Name             string
	Resume           string
	Description      string
	Code             string
	ProjectLeadID    int
	ProjectStartDate time.Time
	ProjectEndDate   time.Time
}

type User struct {
	ID        int `json:"id"`
	Username  string
	Password  string
	FullName  string
	Email     string
	RoleID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID          int
	Name        string
	Description string
}
