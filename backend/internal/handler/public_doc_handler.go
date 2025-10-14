package handler

import (
    "coop/internal/interfaces"
    "coop/internal/models"
    "net/http"
    "path/filepath"
    "strconv"

    "github.com/gin-gonic/gin"
)

type PublicDocHandler struct {
    publicDocService interfaces.PublicDocService
}

func NewPublicDocHandler(publicDocService interfaces.PublicDocService) *PublicDocHandler {
    return &PublicDocHandler{publicDocService}
}

// GetPublicDocByID retrieves a public document by ID
// @Summary Get public document by ID
// @Tags public-documents
// @Param id path int true "Document ID"
// @Success 200 {object} models.PublicDocWithDetails
// @Router /api/v1/public-documents/{id} [get]
func (h *PublicDocHandler) GetPublicDocByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
        return
    }

    doc, err := h.publicDocService.GetPublicDocByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
        return
    }

    c.JSON(http.StatusOK, doc)
}

// GetAllPublicDocs retrieves all public documents with pagination
// @Summary Get all public documents
// @Tags public-documents
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.PublicDocListResponse
// @Router /api/v1/public-documents [get]
func (h *PublicDocHandler) GetAllPublicDocs(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    response, err := h.publicDocService.GetAllPublicDocs(c.Request.Context(), page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get documents"})
        return
    }

    c.JSON(http.StatusOK, response)
}

// CreatePublicDoc creates a new public document
// @Summary Create public document
// @Tags public-documents
// @Accept multipart/form-data
// @Param title formData string true "Title"
// @Param teacher_id formData int false "Teacher ID"
// @Param file formData file true "Document file"
// @Success 201 {object} models.PublicDocWithDetails
// @Router /api/v1/public-documents [post]
func (h *PublicDocHandler) CreatePublicDoc(c *gin.Context) {
    // Parse multipart form
    if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32 MB max
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
        return
    }

    title := c.PostForm("title")
    teacherIDStr := c.PostForm("teacher_id")

    if title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
        return
    }

    req := models.CreatePublicDocRequest{
        Title: title,
    }

    if teacherIDStr != "" {
        teacherID, err := strconv.Atoi(teacherIDStr)
        if err == nil {
            req.TeacherID = &teacherID
        }
    }

    // Get uploaded file
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
        return
    }

    doc, err := h.publicDocService.CreatePublicDoc(c.Request.Context(), req, file)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message":  "Document created successfully",
        "document": doc,
    })
}

// UpdatePublicDoc updates a public document
// @Summary Update public document
// @Tags public-documents
// @Param id path int true "Document ID"
// @Param request body models.UpdatePublicDocRequest true "Update request"
// @Success 200 {object} models.PublicDocWithDetails
// @Router /api/v1/public-documents/{id} [put]
func (h *PublicDocHandler) UpdatePublicDoc(

c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
        return
    }

    var req models.UpdatePublicDocRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    doc, err := h.publicDocService.UpdatePublicDoc(c.Request.Context(), id, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":  "Document updated successfully",
        "document": doc,
    })
}

// DeletePublicDoc deletes a public document
// @Summary Delete public document
// @Tags public-documents
// @Param id path int true "Document ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/public-documents/{id} [delete]
func (h *PublicDocHandler) DeletePublicDoc(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.

StatusBadRequest, gin.H{"error": "Invalid document ID"})
        return
    }

    if err := h.publicDocService.DeletePublicDoc(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}

// SearchPublicDocs searches public documents by title
// @Summary Search public documents
// @Tags public-documents
// @Param q query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} models.PublicDocListResponse
// @Router /api/v1/public-documents/search [get]
func (h *PublicDocHandler) SearchPublicDocs(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
        return
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

    response, err := h.publicDocService.SearchPublicDocs(c.Request.Context(), query, page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search documents"})
        return
    }

    c.JSON(http.StatusOK, response)
}

// DownloadPublicDoc downloads a public document
// @Summary Download public document
// @Tags public-documents
// @Param id path int true "Document ID"
// @Success 200 {file} binary
// @Router /api/v1/public-documents/{id}/download [get]
func (h *PublicDocHandler) DownloadPublicDoc(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
        return
    }

    filePath, err := h.publicDocService.DownloadPublicDoc(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
        return
    }

    // Set headers for file download
    filename := filepath.Base(filePath)
    c.Header("Content-Description", "File Transfer")
    c.Header("Content-Transfer-Encoding", "binary")
    c.Header("Content-Disposition", "attachment; filename="+filename)
    c.Header("Content-Type", "application/octet-stream")

    c.File(filePath)
}