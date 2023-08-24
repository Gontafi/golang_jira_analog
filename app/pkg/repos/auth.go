package repos

import (
	"context"
	"fmt"
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/queries"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"time"
)

type AuthRepository struct {
	db  *pgx.Conn
	rdb *redis.Client
	ctx context.Context
}

func NewAuthRepository(ctx context.Context, db *pgx.Conn, rdb *redis.Client) *AuthRepository {

	return &AuthRepository{
		ctx: ctx,
		db:  db,
		rdb: rdb,
	}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	var id int
	err := r.db.QueryRow(r.ctx, queries.CreateUser,
		user.Username, user.Password, user.FullName, user.Email,
		user.RoleID, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUser(username string, password string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, queries.GetUserByUsernameAndPassword, username, password).Scan(
		&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.FullName,
		&user.Email, &user.RoleID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *AuthRepository) GenerateSessionKey(userID int) (string, error) {
	sessionKey := fmt.Sprintf("session:%d:%s", userID, uuid.New().String())

	err := r.rdb.Set(r.ctx, sessionKey, "active", time.Hour*24*7).Err()
	if err != nil {
		return "", err
	}

	return sessionKey, nil
}

func (r *AuthRepository) StoreResetCode(userID int, code string) error {

	key := fmt.Sprintf("resetcode:%d", userID)
	err := r.rdb.Set(r.ctx, key, code, time.Minute*15).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetResetCode(userId int) (string, error) {
	key := fmt.Sprintf("resetcode:%d", userId)
	code, err := r.rdb.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", err
		}
		return "", err
	}
	return code, nil
}
func (r *AuthRepository) ChangePassword(username string, newPassword string) error {
	_, err := r.db.Exec(r.ctx, queries.UpdateUserPassword, username, newPassword)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, queries.GetUserByUsername, username).Scan(
		&user.ID, &user.Username, &user.FullName,
		&user.Email, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
