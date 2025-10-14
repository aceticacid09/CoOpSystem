package handler

import (
    "coop/internal/interfaces"
    "coop/internal/models"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
    return &UserHandler{userService}
}

// GetUserByID retrieves a user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := h.userService.GetUserByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Don't return password
    user.Password = ""
    c.JSON(http.StatusOK, user)
}

// GetAllUsers retrieves all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
    users, err := h.userService.GetAllUsers(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
        return
    }

    // Don't return passwords
    for i := range users {
        users[i].Password = ""
    }

    c.JSON(http.StatusOK, users)
}

// CreateUser creates a new user with profile
func (h *UserHandler) CreateUser(c *gin.Context) {
    // Parse multipart form
    if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
        return
    }

    username := c.PostForm("username")
    password := c.PostForm("password")
    role := c.PostForm("role")

    if username == "" || password == "" || role == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username, password, and role are required"})
        return
    }

    req := models.CreateUserRequest{
        Username: username,
        Password: password,
        Role:     role,
    }

    // Handle profile image upload
    var profileImagePath *string
    file, header, err := c.Request.FormFile("profile_image")
    if err == nil && file != nil {
        defer file.Close()

        ext := filepath.Ext(header.Filename)
        if ext == "" {
            ext = ".jpg"
        }

        sanitizedUsername := strings.ReplaceAll(username, " ", "_")
        sanitizedUsername = regexp.MustCompile(`[^a-zA-Z0-9_-]`).ReplaceAllString(sanitizedUsername, "")
        filename := fmt.Sprintf("%s_profile%s", sanitizedUsername, ext)

        uploadDir := "/app/uploads/profile"
        if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
            return
        }

        dst := filepath.Join(uploadDir, filename)
        out, err := os.Create(dst)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
            return
        }
        defer out.Close()

        if _, err = io.Copy(out, file); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
            return
        }

        profileImagePath = &filename
    }

    // Parse role-specific fields
    switch role {
    case "teacher":
        nameTH := c.PostForm("name_th")
        nameEN := c.PostForm("name_en")
        emailContact := c.PostForm("email_contact")

        if nameTH != "" {
            req.TeacherNameTH = &nameTH
        }
        if nameEN != "" {
            req.TeacherNameEN = &nameEN
        }
        if emailContact != "" {
            req.TeacherEmailContact = &emailContact
        }
        if deptIDStr := c.PostForm("dept_id"); deptIDStr != "" {
            if deptID, err := strconv.Atoi(deptIDStr); err == nil {
                req.TeacherDeptID = &deptID
            }
        }
        req.TeacherProfileImage = profileImagePath

    case "student":
        nameTH := c.PostForm("name_th")
        nameEN := c.PostForm("name_en")
        studentCodeStr := c.PostForm("student_code")
        educationLevel := c.PostForm("education_level")

        if nameTH != "" {
            req.StudentNameTH = &nameTH
        }
        if nameEN != "" {
            req.StudentNameEN = &nameEN
        }
        if studentCodeStr != "" {
            if studentCode, err := strconv.Atoi(studentCodeStr); err == nil {
                req.StudentCode = &studentCode
            }
        }
        if educationLevel != "" {
            req.StudentEducationLevel = &educationLevel
        }
        if deptIDStr := c.PostForm("dept_id"); deptIDStr != "" {
            if deptID, err := strconv.Atoi(deptIDStr); err == nil {
                req.StudentDeptID = &deptID
            }
        }
        if yearStr := c.PostForm("year"); yearStr != "" {
            if year, err := strconv.Atoi(yearStr); err == nil {
                req.StudentYear = &year
            }
        }
        req.StudentProfileImage = profileImagePath

    case "company":
        companyName := c.PostForm("company_name")
        companyStatus := c.PostForm("status")

        if companyName != "" {
            req.CompanyName = &companyName
        }
        if companyStatus != "" {
            req.CompanyStatus = &companyStatus
        }
        req.CompanyProfileImage = profileImagePath
    }

    user, err := h.userService.CreateUser(c.Request.Context(), req)
    if err != nil {
        if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "already exists") {
            c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Don't return password
    user.Password = ""
    c.JSON(http.StatusCreated, gin.H{
        "message": "User created successfully",
        "user":    user,
    })
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var req models.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    updatedUser, err := h.userService.UpdateUser(c.Request.Context(), id, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Don't return password
    updatedUser.Password = ""
    c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUserByUsername retrieves a user by username
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
    username := c.Param("username")

    user, err := h.userService.GetUserByUsername(c.Request.Context(), username)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Don't return password
    user.Password = ""
    c.JSON(http.StatusOK, user)
}

// GetAllDepartments retrieves all departments
func (h *UserHandler) GetAllDepartments(c *gin.Context) {
    departments, err := h.userService.GetAllDepartments(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get departments"})
        return
    }

    c.JSON(http.StatusOK, departments)
}

// GetDepartmentByID retrieves a department by ID
func (h *UserHandler) GetDepartmentByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
        return
    }

    department, err := h.userService.GetDepartmentByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
        return
    }

    c.JSON(http.StatusOK, department)
}