package models

type UsersProjects struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProjectID int `json:"project_id"`
}
