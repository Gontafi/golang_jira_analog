package models

import "time"

type Issue struct {
	ID               int       `json:"id"`
	ProjectID        int       `json:"project_id"`
	IssueTypeID      int       `json:"issue_type_id"`
	IssueSummary     string    `json:"issue_summary"`
	IssueDescription string    `json:"issue_description"`
	ReporterID       int       `json:"reporter_id"`
	AssigneeID       int       `json:"assignee_id"`
	StageId          int       `json:"stage_id"`
	StatusID         int       `json:"status_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ResolverAt       time.Time `json:"resolver_at"`
}
