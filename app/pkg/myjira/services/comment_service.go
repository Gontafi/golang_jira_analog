package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"log/slog"
	"time"
)

type CommentService struct {
	commentRepo *repos.CommentRepos
}

func NewCommentService(commentRepo *repos.CommentRepos) *CommentService {
	return &CommentService{commentRepo}
}

func (s *CommentService) AddComment(comment models.Comment) (int, error) {
	var id int

	comment.CreatedAt = time.Now()

	id, err := s.commentRepo.CreateComment(comment)
	if err != nil {
		slog.Error("Error in Comment service while creating", err)
		return 0, err
	}
	return id, nil
}

func (s *CommentService) GetByCommentID(id int) (models.Comment, error) {
	comment, err := s.commentRepo.GetByCommentID(id)

	if err != nil {
		slog.Error("Error in Comment service while getting", err)
		return models.Comment{}, err
	}

	return comment, nil
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
	comments, err := s.commentRepo.GetAllComments()
	if err != nil {
		slog.Error("Error in Comment service while getting", err)
		return []models.Comment{}, err
	}
	return comments, nil
}

func (s *CommentService) UpdateComment(comment models.Comment) error {
	err := s.commentRepo.UpdateComment(comment)
	if err != nil {
		slog.Error("Error in Comment service while updating", err)
		return err
	}
	return nil
}

func (s *CommentService) DeleteComment(id int) error {
	err := s.commentRepo.DeleteComment(id)
	if err != nil {
		slog.Error("Error in Comment service while deleting", err)
		return err
	}
	return nil
}
