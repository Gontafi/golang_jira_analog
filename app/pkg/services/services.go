package services

import (
	"github.com/Gontafi/golang_jira_analog/pkg/repos"
)

type Services struct {
	Repos      *repos.Repository
	Comment    *CommentService
	Ticket     *TicketService
	TicketType *TicketTypeService
	Project    *ProjectService
	Role       *RoleService
	Status     *StatusService
	Priority   *PriorityService
	User       *UserService
	Stage      *StageService

	UsersProjects *UsersProjectsService
	Auth          *AuthService
}

func NewServices(r *repos.Repository) *Services {
	return &Services{
		Repos:         r,
		Comment:       NewCommentService(r.Comment),
		Ticket:        NewTicketService(r.Ticket),
		TicketType:    NewTicketTypeService(r.TicketType),
		Project:       NewProjectService(r.Project),
		Role:          NewRoleService(r.Role),
		Status:        NewStatusService(r.Status),
		Priority:      NewPriorityService(r.Priority),
		User:          NewUserService(r.User),
		Stage:         NewStageService(r.Stage),
		UsersProjects: NewUsersProjectsService(r.UsersProject),
		Auth:          NewAuthService(r.Auth),
	}
}
