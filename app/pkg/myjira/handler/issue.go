package handler

import (
	"net/http"
	"strconv"

	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateIssue(c *gin.Context) {
	var issue models.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	issueID, err := h.service.Issue.AddIssue(issue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create issue"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": issueID})
}

func (h *Handler) GetAllIssues(c *gin.Context) {
	issues, err := h.service.Issue.GetAllIssues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch issues"})
		return
	}

	c.JSON(http.StatusOK, issues)
}

func (h *Handler) GetIssueById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue ID"})
		return
	}

	issue, err := h.service.Issue.GetByIssueID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	c.JSON(http.StatusOK, issue)
}

func (h *Handler) UpdateIssue(c *gin.Context) {
	var issue models.Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Issue.UpdateIssue(issue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update issue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue updated successfully"})
}

func (h *Handler) DeleteIssue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue ID"})
		return
	}

	err = h.service.Issue.DeleteIssue(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete issue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue deleted successfully"})
}
