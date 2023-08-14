package models

import "time"

type User struct {
	Id        int `JSON:"id"`
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Password  string
	FullName  string
	Email     string
	RoleId    int
}

type Role struct {
	Id          int
	Name        string
	Description string
}
