package repos

import (
	"context"
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/repos/pkg"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	User       *pkg.UserRepos
	IssueType  *pkg.IssueTypeRepos
	Attachment *pkg.AttachmentRepos
	Comment    *pkg.CommentRepos
	Issue      *pkg.IssueRepos
	Priority   *pkg.PriorityRepos
	Project    *pkg.ProjectRepos
	Role       *pkg.RoleRepos
	Status     *pkg.StatusRepos
	Tag        *pkg.TagRepos
}

func NewRepository(ctx context.Context, db *pgx.Conn) *Repository {
	return &Repository{
		User:       pkg.NewUserRepos(ctx, db),
		IssueType:  pkg.NewIssueTypeRepos(ctx, db),
		Attachment: pkg.NewAttachmentRepos(ctx, db),
		Comment:    pkg.NewCommentRepos(ctx, db),
		Issue:      pkg.NewIssueRepos(ctx, db),
		Priority:   pkg.NewPriorityRepos(ctx, db),
		Project:    pkg.NewProjectRepos(ctx, db),
		Role:       pkg.NewRoleRepos(ctx, db),
		Status:     pkg.NewStatusRepos(ctx, db),
		Tag:        pkg.NewTagsRepos(ctx, db),
	}
}
