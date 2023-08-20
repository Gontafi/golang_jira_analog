package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"log/slog"
)

type RoleService struct {
	roleRepo *repos.RoleRepos
}

func NewRoleService(roleRepo *repos.RoleRepos) *RoleService {
	return &RoleService{roleRepo}
}

func (s *RoleService) AddRole(role models.Role) (int, error) {
	var id int
	id, err := s.roleRepo.CreateRole(role)
	if err != nil {
		slog.Error("Error in Role service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *RoleService) GetByRoleID(id int) (models.Role, error) {
	role, err := s.roleRepo.GetByRoleID(id)

	if err != nil {
		slog.Error("Error in Role service while getting", err)
		return models.Role{}, err
	}

	return role, nil
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	roles, err := s.roleRepo.GetAllRoles()
	if err != nil {
		slog.Error("Error in Role service while getting", err)
		return []models.Role{}, err
	}
	return roles, nil
}

func (s *RoleService) UpdateRole(role models.Role) error {
	err := s.roleRepo.UpdateRole(role)
	if err != nil {
		slog.Error("Error in Role service while updating", err)
		return err
	}
	return nil
}

func (s *RoleService) DeleteRole(id int) error {
	err := s.roleRepo.DeleteRole(id)
	if err != nil {
		slog.Error("Error in Role service while deleting", err)
		return err
	}
	return nil
}
