package handler

import (
	"net/http"
	"strconv"

	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStatus(c *gin.Context) {
	var status models.Status
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusID, err := h.service.Status.AddStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create status"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": statusID})
}

func (h *Handler) GetAllStatuses(c *gin.Context) {
	statuses, err := h.service.Status.GetAllStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statuses"})
		return
	}

	c.JSON(http.StatusOK, statuses)
}

func (h *Handler) GetStatusById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	status, err := h.service.Status.GetByStatusID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}

	c.JSON(http.StatusOK, status)
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	var status models.Status
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Status.UpdateStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}

func (h *Handler) DeleteStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	err = h.service.Status.DeleteStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status deleted successfully"})
}
