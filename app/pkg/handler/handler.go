package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		service: services,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	auth := app.Group("/auth")
	{
		auth.Post("/sign-up", h.SignUp)
		auth.Post("/sign-in", h.SignIn)
		forgotPassword := auth.Group("/forgot-password")
		{
			forgotPassword.Post("/", h.ForgotPassword)
			forgotPassword.Post("verify", h.VerifyCode)
		}
	}

	api := app.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.Get("/", h.GetAllUsers)
			users.Get("/:id", h.GetUserByID)
			users.Post("/", h.CreateUser)
			users.Put("/:id", h.UpdateUser)
			users.Delete("/:id", h.DeleteUser)

			projects := users.Group("/:id")
			{
				projects.Get("/projects", h.GetProjectsFromUser)
			}
		}
		comments := api.Group("/comments")
		{
			comments.Get("/", h.GetAllComments)
			comments.Get("/:id", h.GetCommentById)
			comments.Post("/", h.CreateComment)
			comments.Put("/:id", h.UpdateComment)
			comments.Delete("/:id", h.DeleteComment)
		}
		projects := api.Group("/projects")
		{
			projects.Get("/", h.GetAllProjects)
			projects.Get("/:id", h.GetProjectById)
			projects.Post("/", h.CreateProject)
			projects.Put("/:id", h.UpdateProject)
			projects.Delete("/:id", h.DeleteProject)

			users := projects.Group("/:id/users")
			{
				users.Get("/", h.GetUsersFromProject)
				users.Post("/:id", h.AddUserToProject)
				users.Delete("/:id", h.DeleteUsersFromProject)
			}
		}
		ticket := api.Group("/ticket")
		{
			ticket.Get("/", h.GetAllTickets)
			ticket.Get("/:id", h.GetTicketById)
			ticket.Post("/", h.CreateTicket)
			ticket.Put("/:id", h.UpdateTicket)
			ticket.Delete("/:id", h.DeleteTicket)
		}
		ticketType := api.Group("/ticket-type")
		{
			ticketType.Get("/", h.GetAllTicketTypes)
			ticketType.Get("/:id", h.GetTicketTypeById)
			ticketType.Post("/", h.CreateTicketType)
			ticketType.Put("/:id", h.UpdateTicketType)
			ticketType.Delete("/:id", h.DeleteTicketType)
		}
		roles := api.Group("/roles")
		{
			roles.Get("/", h.GetAllRoles)
			roles.Get("/:id", h.GetRoleById)
			roles.Post("/", h.CreateRole)
			roles.Put("/:id", h.UpdateRole)
			roles.Delete("/:id", h.DeleteRole)
		}
		statuses := api.Group("/statuses")
		{
			statuses.Get("/", h.GetAllStatuses)
			statuses.Get("/:id", h.GetStatusById)
			statuses.Post("/", h.CreateStatus)
			statuses.Put("/:id", h.UpdateStatus)
			statuses.Delete("/:id", h.DeleteStatus)
		}
		priorities := api.Group("/priorities")
		{
			priorities.Get("/", h.GetAllPriorities)
			priorities.Get("/:id", h.GetPriorityById)
			priorities.Post("/", h.CreatePriority)
			priorities.Put("/:id", h.UpdatePriority)
			priorities.Delete("/:id", h.DeletePriority)
		}
	}
	return app
}
