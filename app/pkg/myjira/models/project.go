package models

import (
	"time"
)

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
