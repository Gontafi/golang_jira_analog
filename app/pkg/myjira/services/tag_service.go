package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"log/slog"
)

type TagService struct {
	tagRepo *repos.TagRepos
}

func NewTagService(tagRepo *repos.TagRepos) *TagService {
	return &TagService{tagRepo}
}

func (s *TagService) AddTag(tag models.Tag) (int, error) {
	var id int
	id, err := s.tagRepo.CreateTag(tag)
	if err != nil {
		slog.Error("Error in Tag service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *TagService) GetByTagID(id int) (models.Tag, error) {
	tag, err := s.tagRepo.GetByTagID(id)

	if err != nil {
		slog.Error("Error in Tag service while getting", err)
		return models.Tag{}, err
	}

	return tag, nil
}

func (s *TagService) GetAllTags() ([]models.Tag, error) {
	tags, err := s.tagRepo.GetAllTags()
	if err != nil {
		slog.Error("Error in Tag service while getting", err)
		return []models.Tag{}, err
	}
	return tags, nil
}

func (s *TagService) UpdateTag(tag models.Tag) error {
	err := s.tagRepo.UpdateTag(tag)
	if err != nil {
		slog.Error("Error in Tag service while updating", err)
		return err
	}
	return nil
}

func (s *TagService) DeleteTag(id int) error {
	err := s.tagRepo.DeleteTag(id)
	if err != nil {
		slog.Error("Error in Tag service while deleting", err)
		return err
	}
	return nil
}
