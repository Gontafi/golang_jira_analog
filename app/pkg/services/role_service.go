package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/models"
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
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
		return 0, err
	}
	return id, nil
}

func (s *RoleService) GetByRoleID(id int) (models.Role, error) {
	role, err := s.roleRepo.GetByRoleID(id)

	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	roles, err := s.roleRepo.GetAllRoles()
	if err != nil {
		return []models.Role{}, err
	}
	return roles, nil
}

func (s *RoleService) UpdateRole(role models.Role) error {
	err := s.roleRepo.UpdateRole(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) DeleteRole(id int) error {
	err := s.roleRepo.DeleteRole(id)
	if err != nil {
		return err
	}
	return nil
}
