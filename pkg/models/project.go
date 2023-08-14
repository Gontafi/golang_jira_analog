package models

import "time"

type Project struct {
	Id               int
	Name             string
	Resume           string
	Description      string
	Code             string
	ProjectLead      int
	ProjectStartDate time.Time
	ProjectEndDate   time.Time
}
