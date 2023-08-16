package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
	"time"
)

type UserService struct {
	userRepo *pkg.UserRepos
}

func NewUserService(userRepo *pkg.UserRepos) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(user models.User) (int, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	var id int
	id, err := s.userRepo.Create(user)
	if err != nil {
		slog.Error("failed in UserService:", err)
		return 0, err
	}
	return id, nil
}

func (s *UserService) GetByID(id int) (models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil || id < 0 {
		slog.Error("failed in UserService", err)
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		slog.Error("failed in UserService", err)
		return []models.User{}, err
	}
	return users, nil
}

func (s *UserService) Update(user models.User) error {
	user.UpdatedAt = time.Now()
	err := s.userRepo.Update(user)
	if err != nil {
		slog.Error("failed in UserService", err)
		return err
	}
	return nil
}

func (s *UserService) Delete(id int) error {
	err := s.userRepo.Delete(id)
	if err != nil || id < 0 {
		slog.Error("failed in UserService", err)
		return err
	}
	return nil
}
