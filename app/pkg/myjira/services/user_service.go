package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"log/slog"
	"time"
)

type UserService struct {
	userRepo *repos.UserRepos
}

func NewUserService(userRepo *repos.UserRepos) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) AddUser(user models.User) (int, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	var id int
	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		slog.Error("failed in UserService:", err)
		return 0, err
	}
	return id, nil
}

func (s *UserService) GetByUserID(id int) (models.User, error) {
	user, err := s.userRepo.GetByUserID(id)
	if err != nil || id < 0 {
		slog.Error("failed in UserService", err)
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		slog.Error("failed in UserService", err)
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (models.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		slog.Error("failed in UserService", err)
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		slog.Error("failed in UserService", err)
		return []models.User{}, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(user models.User) error {
	user.UpdatedAt = time.Now()
	err := s.userRepo.UpdateUser(user)
	if err != nil {
		slog.Error("failed in UserService", err)
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.userRepo.DeleteUser(id)
	if err != nil || id < 0 {
		slog.Error("failed in UserService", err)
		return err
	}
	return nil
}
