package interfaces

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/services"
)

type AttachmentService interface {
	AddAttachment(attachment models.Attachment) (int, error)
	GetByAttachmentID(id int) (models.Attachment, error)
	GetAllAttachments() ([]models.Attachment, error)
	UpdateAttachment(attachment models.Attachment) error
	DeleteAttachment(id int) error
}

type CommentService interface {
	AddComment(comment models.Comment) (int, error)
	GetByCommentID(id int) (models.Comment, error)
	GetAllComments() ([]models.Comment, error)
	UpdateComment(comment models.Comment) error
	DeleteComment(id int) error
}

type IssueService interface {
	AddIssue(issue models.Issue) (int, error)
	GetByIssueID(id int) (models.Issue, error)
	GetAllIssues() ([]models.Issue, error)
	UpdateIssue(issue models.Issue) error
	DeleteIssue(id int) error
}

type IssueTypeService interface {
	AddIssueType(issueType models.IssueType) (int, error)
	GetByIssueTypeID(id int) (models.IssueType, error)
	GetAllIssueTypes() ([]models.IssueType, error)
	UpdateIssueType(issueType models.IssueType) error
	DeleteIssueType(id int) error
}

type ProjectService interface {
	AddProject(project models.Project) (int, error)
	GetByProjectID(id int) (models.Project, error)
	GetAllProjects() ([]models.Project, error)
	UpdateProject(project models.Project) error
	DeleteProject(id int) error
}

type RoleService interface {
	AddRole(role models.Role) (int, error)
	GetByRoleID(id int) (models.Role, error)
	GetAllRoles() ([]models.Role, error)
	UpdateRole(role models.Role) error
	DeleteRole(id int) error
}

type StatusService interface {
	AddStatus(status models.Status) (int, error)
	GetByStatusID(id int) (models.Status, error)
	GetAllStatuses() ([]models.Status, error)
	UpdateStatus(status models.Status) error
	DeleteStatus(id int) error
}

type TagService interface {
	AddTag(tag models.Tag) (int, error)
	GetByTagID(id int) (models.Tag, error)
	GetAllTags() ([]models.Tag, error)
	UpdateTag(tag models.Tag) error
	DeleteTag(id int) error
}

type UserService interface {
	AddUser(user models.User) (int, error)
	GetByUserID(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id int) error
}

type Services struct {
	Repos      *Repository
	Attachment *services.AttachmentService
	Comment    *services.CommentService
	Issue      *services.IssueService
	IssueType  *services.IssueTypeService
	Project    *services.ProjectService
	Role       *services.RoleService
	Status     *services.StatusService
	Tag        *services.TagService
	User       *services.UserService

	Auth *services.AuthService
}

func NewServices(r *Repository) *Services {
	return &Services{
		Repos:      r,
		Attachment: services.NewAttachmentService(r.Attachment),
		Comment:    services.NewCommentService(r.Comment),
		Issue:      services.NewIssueService(r.Issue),
		IssueType:  services.NewIssueTypeService(r.IssueType),
		Project:    services.NewProjectService(r.Project),
		Role:       services.NewRoleService(r.Role),
		Status:     services.NewStatusService(r.Status),
		Tag:        services.NewTagService(r.Tag),
		User:       services.NewUserService(r.User),

		Auth: services.NewAuthService(r.Auth),
	}
}
