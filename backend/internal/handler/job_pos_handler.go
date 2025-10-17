package handler

import (
    "coop/internal/interfaces"
    "coop/internal/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type JobPositionHandler struct {
    jobPositionService interfaces.JobPositionService
}

func NewJobPositionHandler(jobPositionService interfaces.JobPositionService) *JobPositionHandler {
    return &JobPositionHandler{jobPositionService}
}

// GetJobPositionByID retrieves a job position by ID
// @Summary Get job position by ID
// @Tags job-positions
// @Param id path int true "Job Position ID"
// @Success 200 {object} models.JobPositionWithDetails
// @Router /api/v1/job-positions/{id} [get]
func (h *JobPositionHandler) GetJobPositionByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job position ID"})
        return
    }

    jobPosition, err := h.jobPositionService.GetJobPositionByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Job position not found"})
        return
    }

    c.JSON(http.StatusOK, jobPosition)
}

// GetAllJobPositions retrieves all job positions with pagination
// @Summary Get all job positions
// @Tags job-positions
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.JobPositionListResponse
// @Router /api/v1/job-positions [get]
func (h *JobPositionHandler) GetAllJobPositions(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    response, err := h.jobPositionService.GetAllJobPositions(c.Request.Context(), page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get job positions"})
        return
    }

    c.JSON(http.StatusOK, response)
}

// GetJobPositionsByStatus retrieves job positions filtered by status
// @Summary Get job positions by status
// @Tags job-positions
// @Param status path string true "Status (open, closed, pending)"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.JobPositionListResponse
// @Router /api/v1/job-positions/status/{status} [get]
func (h *JobPositionHandler) GetJobPositionsByStatus(c *gin.Context) {
    status := c.Param("status")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    response, err := h.jobPositionService.GetJobPositionsByStatus(c.Request.Context(), status, page, pageSize)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

// GetJobPositionsByCompany retrieves job positions filtered by company
// @Summary Get job positions by company
// @Tags job-positions
// @Param company_id path int true "Company ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.JobPositionListResponse
// @Router /api/v1/job-positions/company/{company_id} [get]
func (h *JobPositionHandler) GetJobPositionsByCompany(c *gin.Context) {
    companyIDStr := c.Param("company_id")
    companyID, err := strconv.Atoi(companyIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
        return
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    response, err := h.jobPositionService.GetJobPositionsByCompany(c.Request.Context(), companyID, page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get job positions"})
        return
    }

    c.JSON(http.StatusOK, response)
}

// CreateJobPosition creates a new job position
// @Summary Create job position
// @Tags job-positions
// @Accept multipart/form-data
// @Param title formData string true "Title"
// @Param description formData string false "Description"
// @Param company_id formData int false "Company ID"
// @Param status formData string false "Status (open, closed, pending)"
// @Param files formData file false "Attachment files"
// @Success 201 {object} models.JobPositionWithDetails
// @Router /api/v1/job-positions [post]
func (h *JobPositionHandler) CreateJobPosition(c *gin.Context) {
    // Parse multipart form
    if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32 MB max
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
        return
    }

    title := c.PostForm("title")
    description := c.PostForm("description")
    companyIDStr := c.PostForm("company_id")
    status := c.PostForm("status")

    if title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
        return
    }

    req := models.CreateJobPositionRequest{
        Title: title,
    }

    if description != "" {
        req.Description = &description
    }

    if companyIDStr != "" {
        companyID, err := strconv.Atoi(companyIDStr)
        if err == nil {
            req.CompanyID = &companyID
        }
    }

    if status != "" {
        req.Status = &status
    }

    // Get uploaded files
    form, _ := c.MultipartForm()
    files := form.File["files"]

    jobPosition, err := h.jobPositionService.CreateJobPosition(c.Request.Context(), req, files)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message":      "Job position created successfully",
        "job_position": jobPosition,
    })
}

// UpdateJobPosition updates a job position
// @Summary Update job position
// @Tags job-positions
// @Param id path int true "Job Position ID"
// @Param request body models.UpdateJobPositionRequest true "Update request"
// @Success 200 {object} models.JobPositionWithDetails
// @Router /api/v1/job-positions/{id} [put]
func (h *JobPositionHandler) UpdateJobPosition(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job position ID"})
        return
    }

    var req models.UpdateJobPositionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    jobPosition, err := h.jobPositionService.UpdateJobPosition(c.Request.Context(), id, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":      "Job position updated successfully",
        "job_position": jobPosition,
    })
}

// DeleteJobPosition deletes a job position
// @Summary Delete job position
// @Tags job-positions
// @Param id path int true "Job Position ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/job-positions/{id} [delete]
func (h *JobPositionHandler) DeleteJobPosition(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job position ID"})
        return
    }

    if err := h.jobPositionService.DeleteJobPosition(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Job position deleted successfully"})
}

// AddFilesToJobPosition adds files to an existing job position
// @Summary Add files to job position
// @Tags job-positions
// @Accept multipart/form-data
// @Param id path int true "Job Position ID"
// @Param files formData file true "Files to upload"
// @Success 200 {object} map[string]string
// @Router /api/v1/job-positions/{id}/files [post]
func (h *JobPositionHandler) AddFilesToJobPosition(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job position ID"})
        return
    }

    // Parse multipart form
    if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
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

    if err := h.jobPositionService.AddFilesToJobPosition(c.Request.Context(), id, files); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Files added successfully"})
}

// DeleteFileFromJobPosition deletes a file from a job position
// @Summary Delete file from job position
// @Tags job-positions
// @Param file_id path int true "File ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/job-positions/files/{file_id} [delete]
func (h *JobPositionHandler) DeleteFileFromJobPosition(c *gin.Context) {
    fileIDStr := c.Param("file_id")
    fileID, err := strconv.Atoi(fileIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
        return
    }

    if err := h.jobPositionService.DeleteFileFromJobPosition(c.Request.Context(), fileID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

// SearchJobPositions searches job positions by keyword
// @Summary Search job positions
// @Tags job-positions
// @Param keyword query string true "Search keyword"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.JobPositionListResponse
// @Router /api/v1/job-positions/search [get]
func (h *JobPositionHandler) SearchJobPositions(c *gin.Context) {
    keyword := c.Query("keyword")
    if keyword == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Keyword is required"})
        return
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    response, err := h.jobPositionService.SearchJobPositions(c.Request.Context(), keyword, page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search job positions"})
        return
    }

    c.JSON(http.StatusOK, response)
}