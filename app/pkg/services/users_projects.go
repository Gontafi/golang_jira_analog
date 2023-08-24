package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
)

type UsersProjectsService struct {
	repo *repos.UsersProjectRepository
}

func NewUsersProjectsService(repo *repos.UsersProjectRepository) *UsersProjectsService {
	return &UsersProjectsService{
		repo: repo,
	}
}

func (s *UsersProjectsService) AddUserToProject(userId int, projectId int) (int, error) {
	return s.repo.AddUserToProject(userId, projectId)
}

func (s *UsersProjectsService) GetUsersFromProject(projectId int) ([]models.User, error) {
	return s.repo.GetUsersFromProject(projectId)
}

func (s *UsersProjectsService) GetProjectFromUsers(userId int) ([]models.Project, error) {
	return s.repo.GetProjectFromUsers(userId)
}

func (s *UsersProjectsService) DeleteUsersFromProject(userId int, projectId int) error {
	return s.repo.DeleteUsersFromProject(userId, projectId)
}
