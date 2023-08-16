package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/services/pkg"
)

type Services struct {
	Repos      *repos.Repository
	Attachment *pkg.AttachmentService
	Comment    *pkg.CommentService
	Issue      *pkg.IssueService
	IssueType  *pkg.IssueTypeService
	Priority   *pkg.PriorityService
	Project    *pkg.ProjectService
	Role       *pkg.RoleService
	Status     *pkg.StatusService
	Tag        *pkg.TagService
	User       *pkg.UserService
}

func NewServices(r *repos.Repository) *Services {
	return &Services{
		Repos:      r,
		Attachment: pkg.NewAttachmentService(r.Attachment),
		Comment:    pkg.NewCommentService(r.Comment),
		Issue:      pkg.NewIssueService(r.Issue),
		IssueType:  pkg.NewIssueTypeService(r.IssueType),
		Priority:   pkg.NewPriorityService(r.Priority),
		Project:    pkg.NewProjectService(r.Project),
		Role:       pkg.NewRoleService(r.Role),
		Status:     pkg.NewStatusService(r.Status),
		Tag:        pkg.NewTagService(r.Tag),
		User:       pkg.NewUserService(r.User),
	}
}
