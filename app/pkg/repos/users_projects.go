package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
)

type UsersProjectRepository struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewUserProjectRepos(ctx context.Context, db *pgx.Conn) *UsersProjectRepository {
	return &UsersProjectRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *UsersProjectRepository) AddUserToProject(userId int, projectId int) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateUserProject, userId, projectId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UsersProjectRepository) GetUsersFromProject(projectId int) ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query(r.ctx, q.GetUsersFromProject, projectId)
	if err != nil {
		return []models.User{}, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.RoleID)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UsersProjectRepository) GetProjectFromUsers(userId int) ([]models.Project, error) {
	var projects []models.Project
	rows, err := r.db.Query(r.ctx, q.GetProjectFromUsers, userId)
	if err != nil {
		return []models.Project{}, err
	}

	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description,
			&project.Name, &project.Code, &project.ProjectLeadID)
		if err != nil {
			return []models.Project{}, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (r *UsersProjectRepository) DeleteUsersFromProject(userId int, projectId int) error {
	_, err := r.db.Exec(r.ctx, q.RemoveUserFromProject, userId, projectId)
	if err != nil {
		return err
	}

	return nil
}
