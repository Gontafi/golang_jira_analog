package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
)

type StageRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewStageRepos(ctx context.Context, db *pgx.Conn) *StageRepos {
	return &StageRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *StageRepos) GetByStageID(id int) (models.Stage, error) {
	var stage models.Stage
	err := r.db.QueryRow(r.ctx, queries.GetStageById, id).Scan(
		&stage.ID, &stage.Name)
	if err != nil {
		return models.Stage{}, err
	}
	return stage, nil
}

func (r *StageRepos) CreateStage(stage models.Stage) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, queries.CreateStage,
		stage.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *StageRepos) AllStages() ([]models.Stage, error) {
	rows, err := r.db.Query(r.ctx, queries.GetStages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []models.Stage
	for rows.Next() {
		var stage models.Stage
		err := rows.Scan(&stage.ID, &stage.Name)
		if err != nil {
			return nil, err
		}
		stages = append(stages, stage)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stages, nil
}

func (r *StageRepos) UpdateStage(stage models.Stage) error {
	_, err := r.db.Exec(r.ctx, queries.UpdateStage,
		stage.ID, stage.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *StageRepos) DeleteStage(id int) error {
	_, err := r.db.Exec(r.ctx, queries.DeleteStage, id)
	if err != nil {
		return err
	}
	return nil
}
