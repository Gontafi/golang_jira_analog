package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"log/slog"
	"time"
)

type AttachmentService struct {
	attRepo *repos.AttachmentRepos
}

func NewAttachmentService(attRepo *repos.AttachmentRepos) *AttachmentService {
	return &AttachmentService{attRepo}
}

func (s *AttachmentService) AddAttachment(attachment models.Attachment) (int, error) {
	var id int

	attachment.UploadedDate = time.Now()

	id, err := s.attRepo.CreateAttachment(attachment)
	if err != nil {
		slog.Error("Error in Attachment service while adding", err)
		return 0, err
	}
	return id, nil
}

func (s *AttachmentService) GetByAttachmentID(id int) (models.Attachment, error) {
	attachment, err := s.attRepo.GetByAttachmentID(id)

	if err != nil {
		slog.Error("Error in Attachment service while getting")
		return models.Attachment{}, err
	}

	return attachment, nil
}

func (s *AttachmentService) GetAllAttachments() ([]models.Attachment, error) {
	attachments, err := s.attRepo.GetAllAttachments()
	if err != nil {
		slog.Error("Error in Attachment service while getting", err)
		return []models.Attachment{}, err
	}
	return attachments, nil
}

func (s *AttachmentService) UpdateAttachment(attachment models.Attachment) error {
	err := s.attRepo.UpdateAttachment(attachment)
	if err != nil {
		slog.Error("Error in Attachment service while updating", err)
		return err
	}
	return nil
}

func (s *AttachmentService) DeleteAttachment(id int) error {
	err := s.attRepo.DeleteAttachment(id)
	if err != nil {
		slog.Error("Error in Attachment service while deleting", err)
		return err
	}
	return nil
}
