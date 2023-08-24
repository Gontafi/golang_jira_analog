package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
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

func (r *CommentRepos) GetByCommentID(id int) (models.Comment, error) {
	var comment models.Comment
	err := r.db.QueryRow(r.ctx, q.GetCommentById, id).Scan(
		&comment.ID, &comment.UserID, &comment.TicketID, &comment.CommentText, &comment.CreatedAt)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (r *CommentRepos) CreateComment(comment models.Comment) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateComment,
		comment.UserID, comment.TicketID, comment.CommentText, comment.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CommentRepos) GetAllComments() ([]models.Comment, error) {
	rows, err := r.db.Query(r.ctx, q.GetComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.TicketID, &comment.CommentText, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepos) UpdateComment(comment models.Comment) error {
	_, err := r.db.Exec(r.ctx, q.UpdateComment,
		comment.ID, comment.UserID, comment.TicketID, comment.CommentText)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepos) DeleteComment(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteComment, id)
	if err != nil {
		return err
	}
	return nil
}
