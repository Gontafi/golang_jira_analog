package repos

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	User       *UserRepos
	TicketType *TicketTypeRepos
	Comment    *CommentRepos
	Ticket     *TicketRepos
	Project    *ProjectRepos
	Role       *RoleRepos
	Status     *StatusRepos
	Priority   *PriorityRepos
	Stage      *StageRepos

	Auth         *AuthRepository
	UsersProject *UsersProjectRepository
}

func NewRepository(ctx context.Context, db *pgx.Conn, rdb *redis.Client) *Repository {
	return &Repository{
		User:       NewUserRepos(ctx, db),
		TicketType: NewTicketTypeRepos(ctx, db),
		Comment:    NewCommentRepos(ctx, db),
		Ticket:     NewTicketRepos(ctx, db),
		Project:    NewProjectRepos(ctx, db),
		Role:       NewRoleRepos(ctx, db),
		Status:     NewStatusRepos(ctx, db),
		Priority:   NewPriorityRepos(ctx, db),
		Stage:      NewStageRepos(ctx, db),

		UsersProject: NewUserProjectRepos(ctx, db),
		Auth:         NewAuthRepository(ctx, db, rdb),
	}
}
