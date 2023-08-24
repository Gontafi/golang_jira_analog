package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	TicketID    int       `json:"ticket_id"`
	CommentText string    `json:"comment_text"`
	CreatedAt   time.Time `json:"created_at"`
}

type Ticket struct {
	ID                int       `json:"id"`
	ProjectID         int       `json:"project_id"`
	TicketTypeID      int       `json:"ticket_type_id"`
	TicketSummary     string    `json:"ticket_summary"`
	TicketDescription string    `json:"ticket_description"`
	ReporterID        int       `json:"reporter_id"`
	AssigneeID        int       `json:"assignee_id"`
	StageID           int       `json:"stage_id"`
	StatusID          int       `json:"status_id"`
	PriorityID        int       `json:"priority_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ResolvedAt        time.Time `json:"resolved_at"`
}

type TicketType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Project struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Resume           string    `json:"resume"`
	Description      string    `json:"description"`
	Code             string    `json:"code"`
	ProjectLeadID    int       `json:"project_lead_id"`
	ProjectStartDate time.Time `json:"project_start_date"`
	ProjectEndDate   time.Time `json:"project_end_date"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Status struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Stage struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        int       `json:"-"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProjects struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProjectID int `json:"project_id"`
}
