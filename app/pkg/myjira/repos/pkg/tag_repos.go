package pkg

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type TagRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewTagsRepos(ctx context.Context, db *pgx.Conn) *TagRepos {
	return &TagRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *TagRepos) GetById(id int) (models.Tag, error) {
	var tag models.Tag
	err := r.db.QueryRow(r.ctx, q.GetTagById, id).Scan(
		&tag.ID, &tag.Name)
	if err != nil {
		slog.Error("Failed on Tag repository")
		return models.Tag{}, err
	}
	return tag, nil
}

func (r *TagRepos) Create(tag models.Tag) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateTag,
		tag.Name).Scan(&id)
	if err != nil {
		slog.Error("Failed to create Tag:", err)
		return 0, err
	}
	return id, nil
}

func (r *TagRepos) GetAll() ([]models.Tag, error) {
	rows, err := r.db.Query(r.ctx, q.GetTags)
	if err != nil {
		slog.Error("Failed on Tag repository")
		return nil, err
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			slog.Error("Failed on Tag repository")
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed on Tag repository")
		return nil, err
	}

	return tags, nil
}

func (r *TagRepos) Update(tag models.Tag) error {
	_, err := r.db.Exec(r.ctx, q.UpdateTag,
		tag.ID, tag.Name)
	if err != nil {
		slog.Error("failed to update Tag")
		return err
	}
	return nil
}

func (r *TagRepos) Delete(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteTag, id)
	if err != nil {
		slog.Error("failed to delete tag:", err)
		return err
	}
	return nil
}
