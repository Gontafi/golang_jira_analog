package pkg

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"log/slog"
	"time"
)

type AttachmentService struct {
	attRepo *pkg.AttachmentRepos
}

func NewAttachmentService(attRepo *pkg.AttachmentRepos) *AttachmentService {
	return &AttachmentService{attRepo}
}

func (s *AttachmentService) Create(attachment models.Attachment) (int, error) {
	var id int

	attachment.UploadedDate = time.Now()

	id, err := s.attRepo.Create(attachment)
	if err != nil {
		slog.Error("Error in Attachment service while adding", err)
		return 0, err
	}
	return id, nil
}

func (s *AttachmentService) GetById(id int) (models.Attachment, error) {
	attachment, err := s.attRepo.GetByID(id)

	if err != nil {
		slog.Error("Error in Attachment service while getting")
		return models.Attachment{}, err
	}

	return attachment, nil
}

func (s *AttachmentService) GetAll() ([]models.Attachment, error) {
	attachments, err := s.attRepo.GetAll()
	if err != nil {
		slog.Error("Error in Attachment service while getting", err)
		return []models.Attachment{}, err
	}
	return attachments, nil
}

func (s *AttachmentService) Update(attachment models.Attachment) error {
	err := s.attRepo.Update(attachment)
	if err != nil {
		slog.Error("Error in Attachment service while updating", err)
		return err
	}
	return nil
}

func (s *AttachmentService) Delete(id int) error {
	err := s.attRepo.Delete(id)
	if err != nil {
		slog.Error("Error in Attachment service while deleting", err)
		return err
	}
	return nil
}
