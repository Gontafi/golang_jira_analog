package models

import "time"

type Issue struct {
	Id               int
	ProjectId        int
	IssueTypeID      int
	IssueSummary     string
	IssueDescription string
	ReporterId       int
	AssigneeId       int
	PriorityId       int
	StatusId         int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ResolverAt       time.Time
}

type IssueType struct {
	Id          int
	name        string
	description string
}

type Priority struct {
	Id          int
	name        string
	description string
}

type Status struct {
	Id          int
	name        string
	description string
}
