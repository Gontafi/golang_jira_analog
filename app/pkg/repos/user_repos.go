package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/jackc/pgx/v5"
)

type UserRepos struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewUserRepos(ctx context.Context, db *pgx.Conn) *UserRepos {
	return &UserRepos{
		db:  db,
		ctx: ctx,
	}
}

func (r *UserRepos) GetByUserID(id int) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, q.GetUserById, id).Scan(
		&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.Password, &user.FullName,
		&user.Email, &user.RoleID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepos) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, q.GetUserByEmail, email).Scan(
		&user.ID, &user.Username, &user.FullName,
		&user.Email, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepos) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, q.GetUserByUsername, username).Scan(
		&user.ID, &user.Username, &user.FullName,
		&user.Email, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepos) CreateUser(user models.User) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, q.CreateUser,
		user.Username, user.Password, user.FullName, user.Email,
		user.RoleID, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepos) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query(r.ctx, q.GetUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.FullName,
			&user.Email, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepos) UpdateUser(user models.User) error {
	_, err := r.db.Exec(r.ctx, q.UpdateUser,
		user.ID, user.UpdatedAt, user.Password, user.FullName, user.Email, user.RoleID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepos) DeleteUser(id int) error {
	_, err := r.db.Exec(r.ctx, q.DeleteUser, id)
	if err != nil {
		return err
	}
	return nil
}
