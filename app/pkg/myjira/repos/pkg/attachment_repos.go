package pkg

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type AttachmentRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewAttachmentRepos(ctx context.Context, db *pgx.Conn) *AttachmentRepos {
	return &AttachmentRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *AttachmentRepos) GetByID(id int) (models.Attachment, error) {
	var attachment models.Attachment
	err := r.db.QueryRow(r.ctx, q.GetAttachmentById, id).Scan(
		&attachment.ID, &attachment.IssueID, &attachment.FileSize,
		&attachment.FileSize, &attachment.UploadedByUserID, &attachment.UploadedDate)
	if err != nil {
		slog.Error("Failed to get Attachment from db:", err)
		return models.Attachment{}, err
	}
	return attachment, nil
}

func (r *AttachmentRepos) Create(attachment models.Attachment) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateAttachment,
		attachment.ID, attachment.IssueID, attachment.FileSize,
		attachment.FileSize, attachment.UploadedByUserID, attachment.UploadedDate).Scan(&id)
	if err != nil {
		slog.Error("Failed to create Attachment:", err)
		return 0, err
	}
	return id, nil
}

func (r *AttachmentRepos) GetAll() ([]models.Attachment, error) {
	rows, err := r.db.Query(r.ctx, q.GetAttachments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []models.Attachment
	for rows.Next() {
		var attachment models.Attachment
		err := rows.Scan(&attachment.ID, &attachment.IssueID, &attachment.FileSize,
			&attachment.FileSize, &attachment.UploadedByUserID, &attachment.UploadedDate)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, attachment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attachments, nil
}

func (r *AttachmentRepos) Update(attachment models.Attachment) error {
	_, err := r.db.Exec(r.ctx, q.UpdateAttachment,
		attachment.ID, attachment.IssueID, attachment.FileSize,
		attachment.FileSize, attachment.UploadedByUserID, attachment.UploadedDate)
	if err != nil {
		slog.Error("failed to update Attachment:", err)
		return err
	}
	return nil
}

func (r *AttachmentRepos) Delete(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteAttachment, id)
	if err != nil {
		slog.Error("failed to delete Attachment:", err)
		return err
	}
	return nil
}
