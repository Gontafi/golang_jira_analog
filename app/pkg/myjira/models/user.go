package models

import "time"

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
