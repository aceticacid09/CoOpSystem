package service

import (
    "context"
    intInterfaces "coop/internal/interfaces"
    intModels "coop/internal/models"
    "fmt"
    "log"
    "time"

    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key") // TODO: Move to environment variable

type userService struct {
    repo intInterfaces.UserRepository
}

func NewUserService(repo intInterfaces.UserRepository) intInterfaces.UserService {
    return &userService{repo: repo}
}

// GetUserByID retrieves a user with profile by ID
func (s *userService) GetUserByID(ctx context.Context, id int) (intModels.UserWithProfile, error) {
    return s.repo.GetUserByID(ctx, id)
}

// GetAllUsers retrieves all users with their profiles
func (s *userService) GetAllUsers(ctx context.Context) ([]intModels.UserWithProfile, error) {
    return s.repo.GetAllUsers(ctx)
}

// CreateUser creates a new user with their role-specific profile
func (s *userService) CreateUser(ctx context.Context, req intModels.CreateUserRequest) (intModels.UserWithProfile, error) {
    log.Println("Starting CreateUser process...")

    // Validate required fields
    if req.Username == "" || req.Password == "" || req.Role == "" {
        return intModels.UserWithProfile{}, fmt.Errorf("username, password, and role are required")
    }

    // Validate role
    validRoles := map[string]bool{
        "admin":   true,
        "teacher": true,
        "student": true,
        "company": true,
    }

    if !validRoles[req.Role] {
        return intModels.UserWithProfile{}, fmt.Errorf("invalid role: %s", req.Role)
    }

    // Hash password
    log.Println("Hashing password...")
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        return intModels.UserWithProfile{}, fmt.Errorf("failed to hash password: %w", err)
    }

    // Create base user
    user := intModels.User{
        Username: req.Username,
        Password: string(hashedPassword),
        Role:     req.Role,
    }

    userID, err := s.repo.CreateUser(ctx, user)
    if err != nil {
        log.Printf("Error creating user: %v", err)
        return intModels.UserWithProfile{}, fmt.Errorf("failed to create user: %w", err)
    }

    log.Printf("User created successfully with ID: %d", userID)

    // Create role-specific profile
    switch req.Role {
    case "teacher":
        if req.TeacherNameTH == nil && req.TeacherNameEN == nil {
            return intModels.UserWithProfile{}, fmt.Errorf("teacher name is required")
        }
        teacher := intModels.Teacher{
            TeacherID:    userID,
            NameTH:       req.TeacherNameTH,
            NameEN:       req.TeacherNameEN,
            DeptID:       req.TeacherDeptID,
            EmailContact: req.TeacherEmailContact,
            ProfileImage: req.TeacherProfileImage,
        }
        if err := s.repo.CreateTeacherProfile(ctx, userID, teacher); err != nil {
            // Rollback user creation if profile creation fails
            s.repo.DeleteUser(ctx, userID)
            return intModels.UserWithProfile{}, fmt.Errorf("failed to create teacher profile: %w", err)
        }

    case "student":
        if req.StudentCode == nil {
            return intModels.UserWithProfile{}, fmt.Errorf("student code is required")
        }
        student := intModels.Student{
            StudentID:      userID,
            NameTH:         req.StudentNameTH,
            NameEN:         req.StudentNameEN,
            StudentCode:    *req.StudentCode,
            DeptID:         req.StudentDeptID,
            Year:           req.StudentYear,
            EducationLevel: req.StudentEducationLevel,
            ProfileImage:   req.StudentProfileImage,
        }
        if err := s.repo.CreateStudentProfile(ctx, userID, student); err != nil {
            s.repo.DeleteUser(ctx, userID)
            return intModels.UserWithProfile{}, fmt.Errorf("failed to create student profile: %w", err)
        }

    case "company":
        if req.CompanyName == nil {
            return intModels.UserWithProfile{}, fmt.Errorf("company name is required")
        }
        status := "pending"
        if req.CompanyStatus != nil {
            status = *req.CompanyStatus
        }
        company := intModels.Company{
            CompanyID:    userID,
            CompanyName:  *req.CompanyName,
            Status:       status,
            ProfileImage: req.CompanyProfileImage,
        }
        if err := s.repo.CreateCompanyProfile(ctx, userID, company); err != nil {
            s.repo.DeleteUser(ctx, userID)
            return intModels.UserWithProfile{}, fmt.Errorf("failed to create company profile: %w", err)
        }
    }

    // Return the created user with profile
    return s.repo.GetUserByID(ctx, userID)
}

// UpdateUser updates a user and their profile
func (s *userService) UpdateUser(ctx context.Context, id int, req intModels.UpdateUserRequest) (intModels.UserWithProfile, error) {
    // Get existing user
    existingUser, err := s.repo.GetUserByID(ctx, id)
    if err != nil {
        return intModels.UserWithProfile{}, fmt.Errorf("user not found: %w", err)
    }

    // Update base user fields if provided
    if req.Username != nil {
        existingUser.Username = *req.Username
    }
    if req.Password != nil {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
        if err != nil {
            return intModels.UserWithProfile{}, fmt.Errorf("failed to hash password: %w", err)
        }
        existingUser.Password = string(hashedPassword)
    }
    if req.Role != nil {
        existingUser.Role = *req.Role
    }

    // Update base user
    if err := s.repo.UpdateUser(ctx, id, existingUser.User); err != nil {
        return intModels.UserWithProfile{}, fmt.Errorf("failed to update user: %w", err)
    }

    // Update role-specific profile
    switch existingUser.Role {
    case "teacher":
        if existingUser.Teacher != nil {
            teacher := *existingUser.Teacher
            if req.NameTH != nil {
                teacher.NameTH = req.NameTH
            }
            if req.NameEN != nil {
                teacher.NameEN = req.NameEN
            }
            if req.DeptID != nil {
                teacher.DeptID = req.DeptID
            }
            if req.EmailContact != nil {
                teacher.EmailContact = req.EmailContact
            }
            if req.ProfileImage != nil {
                teacher.ProfileImage = req.ProfileImage
            }
            if err := s.repo.UpdateTeacherProfile(ctx, id, teacher); err != nil {
                return intModels.UserWithProfile{}, fmt.Errorf("failed to update teacher profile: %w", err)
            }
        }

    case "student":
        if existingUser.Student != nil {
            student := *existingUser.Student
            if req.NameTH != nil {
                student.NameTH = req.NameTH
            }
            if req.NameEN != nil {
                student.NameEN = req.NameEN
            }
            if req.StudentCode != nil {
                student.StudentCode = *req.StudentCode
            }
            if req.DeptID != nil {
                student.DeptID = req.DeptID
            }
            if req.Year != nil {
                student.Year = req.Year
            }
            if req.EducationLevel != nil {
                student.EducationLevel = req.EducationLevel
            }
            if req.ProfileImage != nil {
                student.ProfileImage = req.ProfileImage
            }
            if err := s.repo.UpdateStudentProfile(ctx, id, student); err != nil {
                return intModels.UserWithProfile{}, fmt.Errorf("failed to update student profile: %w", err)
            }
        }

    case "company":
        if existingUser.Company != nil {
            company := *existingUser.Company
            if req.CompanyName != nil {
                company.CompanyName = *req.CompanyName
            }
            if req.CompanyStatus != nil {
                company.Status = *req.CompanyStatus
            }
            if req.ProfileImage != nil {
                company.ProfileImage = req.ProfileImage
            }
            if err := s.repo.UpdateCompanyProfile(ctx, id, company); err != nil {
                return intModels.UserWithProfile{}, fmt.Errorf("failed to update company profile: %w", err)
            }
        }
    }

    // Return updated user with profile
    return s.repo.GetUserByID(ctx, id)
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(ctx context.Context, id int) error {
    return s.repo.DeleteUser(ctx, id)
}

// GetUserByUsername retrieves a user by username
func (s *userService) GetUserByUsername(ctx context.Context, username string) (intModels.User, error) {
    return s.repo.GetUserByUsername(ctx, username)
}

// Login authenticates a user and returns JWT tokens
func (s *userService) Login(ctx context.Context, req intModels.LoginRequest) (intModels.LoginResponse, error) {
    // Get user by username
    user, err := s.repo.GetUserByUsername(ctx, req.Username)
    if err != nil {
        return intModels.LoginResponse{}, fmt.Errorf("invalid credentials")
    }

    // Verify password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return intModels.LoginResponse{}, fmt.Errorf("invalid credentials")
    }

    // Get user profile
    userWithProfile, err := s.repo.GetUserByID(ctx, user.UserID)
    if err != nil {
        return intModels.LoginResponse{}, fmt.Errorf("failed to get user profile")
    }

    // Create JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  user.UserID,
        "username": user.Username,
        "role":     user.Role,
        "exp":      time.Now().Add(time.Hour * 24).Unix(), // 24 hours
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return intModels.LoginResponse{}, fmt.Errorf("failed to generate token")
    }

    // Create refresh token
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.UserID,
        "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
        "type":    "refresh",
    })

    refreshTokenString, err := refreshToken.SignedString(jwtSecret)
    if err != nil {
        return intModels.LoginResponse{}, fmt.Errorf("failed to generate refresh token")
    }

    // Determine profile based on role
    var profile interface{}
    switch user.Role {
    case "teacher":
        profile = userWithProfile.Teacher
    case "student":
        profile = userWithProfile.Student
    case "company":
        profile = userWithProfile.Company
    }

    return intModels.LoginResponse{
        Token:        tokenString,
        RefreshToken: refreshTokenString,
        UserID:       user.UserID,
        Role:         user.Role,
        Profile:      profile,
        Message:      "Login successful",
    }, nil
}

// GetAllDepartments retrieves all departments
func (s *userService) GetAllDepartments(ctx context.Context) ([]intModels.Department, error) {
    return s.repo.GetAllDepartments(ctx)
}

// GetDepartmentByID retrieves a department by ID
func (s *userService) GetDepartmentByID(ctx context.Context, id int) (intModels.Department, error) {
    return s.repo.GetDepartmentByID(ctx, id)
}