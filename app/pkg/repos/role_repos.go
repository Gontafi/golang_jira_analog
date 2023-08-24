package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
)

type RoleRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewRoleRepos(ctx context.Context, db *pgx.Conn) *RoleRepos {
	return &RoleRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *RoleRepos) GetByRoleID(id int) (models.Role, error) {
	var role models.Role
	err := r.db.QueryRow(r.ctx, q.GetRoleById, id).Scan(
		&role.ID, &role.Name, &role.Description)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (r *RoleRepos) CreateRole(role models.Role) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateRole,
		role.Name, role.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RoleRepos) GetAllRoles() ([]models.Role, error) {
	rows, err := r.db.Query(r.ctx, q.GetRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.Description)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepos) UpdateRole(role models.Role) error {
	_, err := r.db.Exec(r.ctx, q.UpdateRole,
		role.ID, role.Name, role.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleRepos) DeleteRole(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteRole, id)
	if err != nil {
		return err
	}
	return nil
}
