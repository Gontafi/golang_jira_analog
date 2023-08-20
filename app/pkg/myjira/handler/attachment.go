package handler

import (
	"net/http"
	"strconv"

	"github.com/Gontafi/golang_jira_analog/pkg/myjira/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAttachment(c *gin.Context) {
	var attachment models.Attachment
	if err := c.ShouldBindJSON(&attachment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attachmentID, err := h.service.Attachment.AddAttachment(attachment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attachment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": attachmentID})
}

func (h *Handler) GetAllAttachments(c *gin.Context) {
	attachments, err := h.service.Attachment.GetAllAttachments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attachments"})
		return
	}

	c.JSON(http.StatusOK, attachments)
}

func (h *Handler) GetAttachmentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attachment ID"})
		return
	}

	attachment, err := h.service.Attachment.GetByAttachmentID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attachment not found"})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

func (h *Handler) UpdateAttachment(c *gin.Context) {
	var attachment models.Attachment
	if err := c.ShouldBindJSON(&attachment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Attachment.UpdateAttachment(attachment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attachment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attachment updated successfully"})
}

func (h *Handler) DeleteAttachment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attachment ID"})
		return
	}

	err = h.service.Attachment.DeleteAttachment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attachment deleted successfully"})
}
