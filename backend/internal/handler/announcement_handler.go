package handler

import (
	"coop/internal/interfaces"
	"coop/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnnouncementHandler struct {
	announcementService interfaces.AnnouncementService
}

func NewAnnouncementHandler(announcementService interfaces.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{announcementService}
}

// GetAnnouncementByID retrieves an announcement by ID
// @Summary Get announcement by ID
// @Tags announcements
// @Param id path int true "Announcement ID"
// @Success 200 {object} models.AnnouncementWithDetails
// @Router /api/v1/announcements/{id} [get]
func (h *AnnouncementHandler) GetAnnouncementByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid announcement ID"})
		return
	}

	announcement, err := h.announcementService.GetAnnouncementByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Announcement not found"})
		return
	}

	c.JSON(http.StatusOK, announcement)
}

// GetAllAnnouncements retrieves all announcements with pagination
// @Summary Get all announcements
// @Tags announcements
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.AnnouncementListResponse
// @Router /api/v1/announcements [get]
func (h *AnnouncementHandler) GetAllAnnouncements(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	response, err := h.announcementService.GetAllAnnouncements(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get announcements"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateAnnouncement creates a new announcement
// @Summary Create announcement
// @Tags announcements
// @Accept multipart/form-data
// @Param title formData string true "Title"
// @Param description formData string false "Description"
// @Param teacher_id formData int false "Teacher ID"
// @Param files formData file false "Attachment files"
// @Success 201 {object} models.AnnouncementWithDetails
// @Router /api/v1/announcements [post]
func (h *AnnouncementHandler) CreateAnnouncement(c *gin.Context) {
	// Parse multipart form
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32 MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	teacherIDStr := c.PostForm("teacher_id")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	req := models.CreateAnnouncementRequest{
		Title: title,
	}

	if description != "" {
		req.Description = &description
	}

	if teacherIDStr != "" {
		teacherID, err := strconv.Atoi(teacherIDStr)
		if err == nil {
			req.TeacherID = &teacherID
		}
	}

	// Get uploaded files
	form, _ := c.MultipartForm()
	files := form.File["files"]

	announcement, err := h.announcementService.CreateAnnouncement(c.Request.Context(), req, files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Announcement created successfully",
		"announcement": announcement,
	})
}

// UpdateAnnouncement updates an announcement
// @Summary Update announcement
// @Tags announcements
// @Param id path int true "Announcement ID"
// @Param request body models.UpdateAnnouncementRequest true "Update request"
// @Success 200 {object} models.AnnouncementWithDetails
// @Router /api/v1/announcements/{id} [put]
func (h *AnnouncementHandler) UpdateAnnouncement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid announcement ID"})
		return
	}

	var req models.UpdateAnnouncementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	announcement, err := h.announcementService.UpdateAnnouncement(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Announcement updated successfully",
		"announcement": announcement,
	})
}

// DeleteAnnouncement deletes an announcement
// @Summary Delete announcement
// @Tags announcements
// @Param id path int true "Announcement ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/announcements/{id} [delete]
func (h *AnnouncementHandler) DeleteAnnouncement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid announcement ID"})
		return
	}

	if err := h.announcementService.DeleteAnnouncement(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete announcement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Announcement deleted successfully"})
}

// AddFilesToAnnouncement adds files to an existing announcement
// @Summary Add files to announcement
// @Tags announcements
// @Accept multipart/form-data
// @Param id path int true "Announcement ID"
// @Param files formData file true "Files to upload"
// @Success 200 {object} map[string]string
// @Router /api/v1/announcements/{id}/files [post]
func (h *AnnouncementHandler) AddFilesToAnnouncement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid announcement ID"})
		return
	}

	// Parse multipart form
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32 MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	// Get uploaded files
	form, _ := c.MultipartForm()
	files := form.File["files"]

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files provided"})
		return
	}

	if err := h.announcementService.AddFilesToAnnouncement(c.Request.Context(), id, files); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files added successfully"})
}

// DeleteFileFromAnnouncement deletes a specific file from an announcement
// @Summary Delete file from announcement
// @Tags announcements
// @Param file_id path int true "File ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/announcements/files/{file_id} [delete]
func (h *AnnouncementHandler) DeleteFileFromAnnouncement(c *gin.Context) {
    fileIDStr := c.Param("file_id")
    fileID, err := strconv.Atoi(fileIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
        return
    }

    // TODO: Check if the authenticated teacher owns the announcement containing this file
    // teacherID := c.GetInt("user_id") // from JWT
    // announcementID, err := h.announcementService.GetAnnouncementIDByFileID(c, fileID)
    // if err != nil {
    //     c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
    //     return
    // }
    // announcement, _ := h.announcementService.GetAnnouncementByID(c, announcementID)
    // if announcement.TeacherID != nil && *announcement.TeacherID != teacherID {
    //     c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete files from your own announcements"})
    //     return
    // }

    if err := h.announcementService.DeleteFileFromAnnouncement(c.Request.Context(), fileID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

// SearchAnnouncements searches announcements by keyword
// @Summary Search announcements
// @Tags announcements
// @Param q query string true "Search keyword"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.AnnouncementListResponse
// @Router /api/v1/announcements/search [get]
func (h *AnnouncementHandler) SearchAnnouncements(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	response, err := h.announcementService.SearchAnnouncements(c.Request.Context(), keyword, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search announcements"})
		return
	}

	c.JSON(http.StatusOK, response)
}

