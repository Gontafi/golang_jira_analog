package models

import "time"

type Comment struct {
	Id          int
	IssueId     int
	UserId      int
	CommentText string
	CreatedAt   time.Time
}
