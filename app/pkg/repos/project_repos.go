package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
)

type ProjectRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewProjectRepos(ctx context.Context, db *pgx.Conn) *ProjectRepos {
	return &ProjectRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *ProjectRepos) GetByProjectID(id int) (models.Project, error) {
	var project models.Project
	err := r.db.QueryRow(r.ctx, q.GetProjectById, id).Scan(
		&project.ID, &project.Name, &project.Resume, &project.Description,
		&project.Code, &project.ProjectLeadID, &project.ProjectStartDate, &project.ProjectEndDate)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}

func (r *ProjectRepos) CreateProject(project models.Project) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateProject,
		project.Name, project.Resume, project.Description,
		project.Code, project.ProjectLeadID, project.ProjectStartDate, project.ProjectEndDate).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProjectRepos) GetAllProjects() ([]models.Project, error) {
	rows, err := r.db.Query(r.ctx, q.GetProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.ID, &project.Name, &project.Resume, &project.Description,
			&project.Code, &project.ProjectLeadID, &project.ProjectStartDate, &project.ProjectEndDate)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectRepos) UpdateProject(project models.Project) error {
	_, err := r.db.Exec(r.ctx, q.UpdateProject,
		project.ID, project.Name, project.Resume, project.Description,
		project.Code, project.ProjectLeadID, project.ProjectStartDate, project.ProjectEndDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProjectRepos) DeleteProject(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteProject, id)
	if err != nil {
		return err
	}
	return nil
}
