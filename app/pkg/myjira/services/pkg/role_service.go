package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
)

type RoleService struct {
	roleRepo *pkg.RoleRepos
}

func NewRoleService(roleRepo *pkg.RoleRepos) *RoleService {
	return &RoleService{roleRepo}
}

func (s *RoleService) Create(role models.Role) (int, error) {
	var id int
	id, err := s.roleRepo.Create(role)
	if err != nil {
		slog.Error("Error in Role service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *RoleService) GetById(id int) (models.Role, error) {
	role, err := s.roleRepo.GetById(id)

	if err != nil {
		slog.Error("Error in Role service while getting", err)
		return models.Role{}, err
	}

	return role, nil
}

func (s *RoleService) GetAll() ([]models.Role, error) {
	roles, err := s.roleRepo.GetAll()
	if err != nil {
		slog.Error("Error in Role service while getting", err)
		return []models.Role{}, err
	}
	return roles, nil
}

func (s *RoleService) Update(role models.Role) error {
	err := s.roleRepo.Update(role)
	if err != nil {
		slog.Error("Error in Role service while updating", err)
		return err
	}
	return nil
}

func (s *RoleService) Delete(id int) error {
	err := s.roleRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Role service while deleting", err)
		return err
	}
	return nil
}
