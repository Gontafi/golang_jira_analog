package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	IssueID     int       `json:"issue_id"`
	CommentText string    `json:"comment_text"`
	CreatedAt   time.Time `json:"created_at"`
}
