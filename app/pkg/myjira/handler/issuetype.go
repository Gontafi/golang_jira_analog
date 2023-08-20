package handler

import (
	"net/http"
	"strconv"

	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateIssueType(c *gin.Context) {
	var issueType models.IssueType
	if err := c.ShouldBindJSON(&issueType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	issueTypeID, err := h.service.IssueType.AddIssueType(issueType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create issue type"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": issueTypeID})
}

func (h *Handler) GetAllIssueTypes(c *gin.Context) {
	issueTypes, err := h.service.IssueType.GetAllIssueTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch issue types"})
		return
	}

	c.JSON(http.StatusOK, issueTypes)
}

func (h *Handler) GetIssueTypeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue type ID"})
		return
	}

	issueType, err := h.service.IssueType.GetByIssueTypeID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue type not found"})
		return
	}

	c.JSON(http.StatusOK, issueType)
}

func (h *Handler) UpdateIssueType(c *gin.Context) {
	var issueType models.IssueType
	if err := c.ShouldBindJSON(&issueType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.IssueType.UpdateIssueType(issueType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update issue type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue type updated successfully"})
}

func (h *Handler) DeleteIssueType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue type ID"})
		return
	}

	err = h.service.IssueType.DeleteIssueType(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete issue type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue type deleted successfully"})
}
