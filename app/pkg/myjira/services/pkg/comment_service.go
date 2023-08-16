package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
	"time"
)

type CommentService struct {
	commentRepo *pkg.CommentRepos
}

func NewCommentService(commentRepo *pkg.CommentRepos) *CommentService {
	return &CommentService{commentRepo}
}

func (s *CommentService) Create(comment models.Comment) (int, error) {
	var id int

	comment.CreatedAt = time.Now()

	id, err := s.commentRepo.Create(comment)
	if err != nil {
		slog.Error("Error in Comment service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *CommentService) GetById(id int) (models.Comment, error) {
	comment, err := s.commentRepo.GetById(id)

	if err != nil {
		slog.Error("Error in Comment service while getting", err)
		return models.Comment{}, err
	}

	return comment, nil
}

func (s *CommentService) GetAll() ([]models.Comment, error) {
	comments, err := s.commentRepo.GetAll()
	if err != nil {
		slog.Error("Error in Comment service while getting", err)
		return []models.Comment{}, err
	}
	return comments, nil
}

func (s *CommentService) Update(comment models.Comment) error {
	err := s.commentRepo.Update(comment)
	if err != nil {
		slog.Error("Error in Comment service while updating", err)
		return err
	}
	return nil
}

func (s *CommentService) Delete(id int) error {
	err := s.commentRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Comment service while deleting", err)
		return err
	}
	return nil
}
