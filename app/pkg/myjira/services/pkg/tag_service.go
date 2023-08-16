package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
)

type TagService struct {
	tagRepo *pkg.TagRepos
}

func NewTagService(tagRepo *pkg.TagRepos) *TagService {
	return &TagService{tagRepo}
}

func (s *TagService) Create(tag models.Tag) (int, error) {
	var id int
	id, err := s.tagRepo.Create(tag)
	if err != nil {
		slog.Error("Error in Tag service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *TagService) GetById(id int) (models.Tag, error) {
	tag, err := s.tagRepo.GetById(id)

	if err != nil {
		slog.Error("Error in Tag service while getting", err)
		return models.Tag{}, err
	}

	return tag, nil
}

func (s *TagService) GetAll() ([]models.Tag, error) {
	tags, err := s.tagRepo.GetAll()
	if err != nil {
		slog.Error("Error in Tag service while getting", err)
		return []models.Tag{}, err
	}
	return tags, nil
}

func (s *TagService) Update(tag models.Tag) error {
	err := s.tagRepo.Update(tag)
	if err != nil {
		slog.Error("Error in Tag service while updating", err)
		return err
	}
	return nil
}

func (s *TagService) Delete(id int) error {
	err := s.tagRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Tag service while deleting", err)
		return err
	}
	return nil
}
