package pkg

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type CommentRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewCommentRepos(ctx context.Context, db *pgx.Conn) *CommentRepos {
	return &CommentRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *CommentRepos) GetById(id int) (models.Comment, error) {
	var comment models.Comment
	err := r.db.QueryRow(r.ctx, q.GetCommentById, id).Scan(
		&comment.ID, &comment.UserID, &comment.IssueID, &comment.CommentText, &comment.CreatedAt)
	if err != nil {
		slog.Error("Failed on Comment repository")
		return models.Comment{}, err
	}
	return comment, nil
}

func (r *CommentRepos) Create(comment models.Comment) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateComment,
		comment.UserID, comment.IssueID, comment.CommentText, comment.CreatedAt).Scan(&id)
	if err != nil {
		slog.Error("Failed to create Comment:", err)
		return 0, err
	}
	return id, nil
}

func (r *CommentRepos) GetAll() ([]models.Comment, error) {
	rows, err := r.db.Query(r.ctx, q.GetComments)
	if err != nil {
		slog.Error("Failed on Comment repository")
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.IssueID, &comment.CommentText, &comment.CreatedAt)
		if err != nil {
			slog.Error("Failed on Comment repository")
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed on Comment repository")
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepos) Update(comment models.Comment) error {
	_, err := r.db.Exec(r.ctx, q.UpdateComment,
		comment.ID, comment.UserID, comment.IssueID, comment.CommentText)
	if err != nil {
		slog.Error("failed to update Comment")
		return err
	}
	return nil
}

func (r *CommentRepos) Delete(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteComment, id)
	if err != nil {
		slog.Error("failed to delete user:")
		return err
	}
	return nil
}
