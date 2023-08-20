package interfaces

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type AttachmentRepos interface {
	CreateAttachment(attachment models.Attachment) (int, error)
	GetByAttachmentID(id int) (models.Attachment, error)
	GetAllAttachments() ([]models.Attachment, error)
	UpdateAttachment(attachment models.Attachment) error
	DeleteAttachment(id int) error
}

type CommentRepos interface {
	CreateComment(comment models.Comment) (int, error)
	GetByCommentID(id int) (models.Comment, error)
	GetAllComments() ([]models.Comment, error)
	UpdateComment(comment models.Comment) error
	DeleteComment(id int) error
}

type IssueRepos interface {
	GetByIssueID(id int) (models.Issue, error)
	CreateIssue(issue models.Issue) (int, error)
	GetAllIssues() ([]models.Issue, error)
	UpdateIssue(issue models.Issue) error
	DeleteIssue(id int) error
}

type TagRepos interface {
	GetByTagID(id int) (models.Tag, error)
	CreateTag(tag models.Tag) (int, error)
	GetAllTags() ([]models.Tag, error)
	UpdateTag(tag models.Tag) error
	DeleteTag(id int) error
}

type IssueTypeRepos interface {
	GetByIssueTypeID(id int) (models.IssueType, error)
	CreateIssueType(issueType models.IssueType) (int, error)
	GetAllIssueTypes() ([]models.IssueType, error)
	UpdateIssueType(issueType models.IssueType) error
	DeleteIssueType(id int) error
}

type StatusRepos interface {
	GetByStatusID(id int) (models.Status, error)
	CreateStatus(issue models.Status) (int, error)
	GetAllStatuses() ([]models.Status, error)
	UpdateStatus(issue models.Status) error
	DeleteStatus(id int) error
}

type ProjectRepos interface {
	GetByProjectID(id int) (models.Project, error)
	CreateProject(project models.Project) (int, error)
	GetAllProjects() ([]models.Project, error)
	UpdateProject(project models.Project) error
	DeleteProject(id int) error
}

type UserRepos interface {
	GetByUserID(id int) (models.User, error)
	CreateUser(user models.User) (int, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id int) error
}

type RoleRepos interface {
	GetByRoleID(id int) (models.Role, error)
	CreateRole(role models.Role) (int, error)
	GetAllRoles() ([]models.Role, error)
	UpdateRole(role models.Role) error
	DeleteRole(id int) error
}

type UsersProjectRepository interface {
	AddUserToProject(userId int, projectId int) (int, error)
	GetUsersFromProject(projectId int) ([]models.User, error)
	GetProjectFromUsers(userId int) ([]models.Project, error)
	DeleteUsersFromProject(userId int, projectId int) error
}

type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string, password string) (models.User, error)
	GenerateSessionKey(userID int) (string, error)
	StoreResetCode(userID int, code string) error
	GetResetCode(userId int) (string, error)
	ChangePassword(username string, newPassword string) error
	GetUserByUsername(username string) (models.User, error)
}

type Repository struct {
	User       *repos.UserRepos
	IssueType  *repos.IssueTypeRepos
	Attachment *repos.AttachmentRepos
	Comment    *repos.CommentRepos
	Issue      *repos.IssueRepos
	Project    *repos.ProjectRepos
	Role       *repos.RoleRepos
	Status     *repos.StatusRepos
	Tag        *repos.TagRepos

	Auth         *repos.AuthRepository
	UsersProject *repos.UsersProjectRepository
}

func NewRepository(ctx context.Context, db *pgx.Conn, rdb *redis.Client) *Repository {
	return &Repository{
		User:       repos.NewUserRepos(ctx, db),
		IssueType:  repos.NewIssueTypeRepos(ctx, db),
		Attachment: repos.NewAttachmentRepos(ctx, db),
		Comment:    repos.NewCommentRepos(ctx, db),
		Issue:      repos.NewIssueRepos(ctx, db),
		Project:    repos.NewProjectRepos(ctx, db),
		Role:       repos.NewRoleRepos(ctx, db),
		Status:     repos.NewStatusRepos(ctx, db),
		Tag:        repos.NewTagsRepos(ctx, db),

		UsersProject: repos.NewUserProjectRepos(ctx, db),
		Auth:         repos.NewAuthRepository(ctx, db, rdb),
	}
}
