package pkg

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/jackc/pgx/v5"
	"log/slog"
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

func (r *RoleRepos) GetById(id int) (models.Role, error) {
	var role models.Role
	err := r.db.QueryRow(r.ctx, q.GetRoleById, id).Scan(
		&role.ID, &role.Name, &role.Description)
	if err != nil {
		slog.Error("Failed on Role repository")
		return models.Role{}, err
	}
	return role, nil
}

func (r *RoleRepos) Create(role models.Role) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateRole,
		role.Name, role.Description).Scan(&id)
	if err != nil {
		slog.Error("Failed to create Role:", err)
		return 0, err
	}
	return id, nil
}

func (r *RoleRepos) GetAll() ([]models.Role, error) {
	rows, err := r.db.Query(r.ctx, q.GetRoles)
	if err != nil {
		slog.Error("Failed on Role repository")
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.Description)
		if err != nil {
			slog.Error("Failed on Role repository")
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed on Role repository")
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepos) Update(role models.Role) error {
	_, err := r.db.Exec(r.ctx, q.UpdateRole,
		role.ID, role.Name, role.Description)
	if err != nil {
		slog.Error("failed to update Role")
		return err
	}
	return nil
}

func (r *RoleRepos) Delete(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteRole, id)
	if err != nil {
		slog.Error("failed to delete role:", err)
		return err
	}
	return nil
}
