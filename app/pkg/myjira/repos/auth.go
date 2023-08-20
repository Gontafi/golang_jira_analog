package repos

import (
	"context"
	"fmt"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	q "github.com/Gontafi/golang_jira_analog/pkg/myjira/queries"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"log/slog"
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
	err := r.db.QueryRow(r.ctx, q.CreateUser,
		user.Username, user.Password, user.FullName, user.Email,
		user.RoleID, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		slog.Error("Failed to create User:", err)
		return 0, err
	}
	return id, nil
}

func (r *AuthRepository) GetUser(username string, password string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, q.GetUserByUsernameAndPassword, username, password).Scan(
		&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.FullName,
		&user.Email, &user.RoleID)
	if err != nil {
		return models.User{}, err
	}
	fmt.Println("get user id:", user.ID)
	return user, nil
}

func (r *AuthRepository) GenerateSessionKey(userID int) (string, error) {
	sessionKey := fmt.Sprintf("session:%d:%s", userID, uuid.New().String())

	err := r.rdb.Set(r.ctx, sessionKey, "active", time.Hour*24*7).Err()
	if err != nil {
		slog.Error("Failed to store session key:", err)
		return "", err
	}

	return sessionKey, nil
}

func (r *AuthRepository) StoreResetCode(userID int, code string) error {

	key := fmt.Sprintf("resetcode:%d", userID)
	err := r.rdb.Set(r.ctx, key, code, time.Minute*15).Err()
	if err != nil {
		slog.Error("Failed to store verification code:", err)
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
		slog.Error("Failed to get reset code:", err)
		return "", err
	}
	return code, nil
}
func (r *AuthRepository) ChangePassword(username string, newPassword string) error {
	_, err := r.db.Exec(r.ctx, q.UpdateUserPassword, username, newPassword)
	if err != nil {
		slog.Error("Failed to update password:", err)
		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, q.GetUserByUsername, username).Scan(
		&user.ID, &user.Username, &user.FullName,
		&user.Email, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		slog.Error("Failed to get user by username from db:", err)
		return models.User{}, err
	}
	return user, nil
}
