package handler

import (
	"github.com/Gontafi/golang_jira_analog/pkg/myjira/interfaces"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *interfaces.Services
}

func NewHandler(services *interfaces.Services) *Handler {
	return &Handler{
		service: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		forgotPassword := auth.Group("/forgot-password")
		{
			forgotPassword.POST("/", h.ForgotPassword)
			forgotPassword.POST("verify", h.VerifyCode)
		}
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/", h.GetAllUsers)
			users.GET("/:id", h.GetUserByID)
			users.POST("/", h.CreateUser)
			users.PUT("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}
		attachments := api.Group("/attachments")
		{
			attachments.GET("/", h.GetAllAttachments)
			attachments.GET("/:id", h.GetAttachmentById)
			attachments.POST("/", h.CreateAttachment)
			attachments.PUT("/:id", h.UpdateAttachment)
			attachments.DELETE("/:id", h.DeleteAttachment)
		}
		comments := api.Group("/comments")
		{
			comments.GET("/", h.GetAllComments)
			comments.GET("/:id", h.GetCommentById)
			comments.POST("/", h.CreateComment)
			comments.PUT("/:id", h.UpdateComment)
			comments.DELETE("/:id", h.DeleteComment)
		}
		projects := api.Group("/projects")
		{
			projects.GET("/", h.GetAllProjects)
			projects.GET("/:id", h.GetAllComments)
			projects.POST("/", h.CreateComment)
			projects.PUT("/:id", h.UpdateComment)
			projects.DELETE("/:id", h.DeleteComment)

			issue := projects.Group("/issue")
			{
				issue.GET("/", h.GetAllIssues)
				issue.GET("/:id", h.GetIssueById)
				issue.POST("/", h.CreateIssue)
				issue.PUT(":id/", h.UpdateIssue)
				issue.DELETE("/:id", h.DeleteIssue)
			}
		}
		issueType := api.Group("/issue-type")
		{
			issueType.GET("/", h.GetAllIssueTypes)
			issueType.GET("/:id", h.GetIssueTypeById)
			issueType.POST("/", h.CreateIssueType)
			issueType.PUT("/:id", h.UpdateIssueType)
			issueType.DELETE("/:id", h.DeleteIssueType)
		}
		roles := api.Group("/roles")
		{
			roles.GET("/", h.GetAllRoles)
			roles.GET("/:id", h.GetRoleById)
			roles.POST("/", h.CreateRole)
			roles.PUT("/:id", h.UpdateRole)
			roles.DELETE("/:id", h.DeleteRole)
		}
		statuses := api.Group("/statuses")
		{
			statuses.GET("/", h.GetAllStatuses)
			statuses.GET("/:id", h.GetStatusById)
			statuses.POST("/", h.CreateStatus)
			statuses.PUT("/:id", h.UpdateStatus)
			statuses.DELETE("/:id", h.DeleteStatus)
		}
		tags := api.Group("/tags")
		{
			tags.GET("/", h.GetAllTags)
			tags.GET("/:id", h.GetTagById)
			tags.POST("/", h.CreateTag)
			tags.PUT("/:id", h.UpdateTag)
			tags.DELETE("/:id", h.DeleteTag)
		}
	}
	return router
}
