package repository

import (
    "context"
    "coop/internal/models"
    "fmt"

    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
    db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
    return &userRepository{db: db}
}

// GetUserByID retrieves a user with their profile by ID
func (r *userRepository) GetUserByID(ctx context.Context, id int) (models.UserWithProfile, error) {
    var result models.UserWithProfile

    // Get base user
    err := r.db.QueryRow(ctx, `
        SELECT user_id, username, password, role
        FROM Users WHERE user_id = $1
    `, id).Scan(
        &result.UserID,
        &result.Username,
        &result.Password,
        &result.Role,
    )
    if err != nil {
        return result, fmt.Errorf("user not found: %w", err)
    }

    // Get role-specific profile
    switch result.Role {
    case "teacher":
        teacher, err := r.GetTeacherProfile(ctx, id)
        if err == nil {
            result.Teacher = teacher
        }
    case "student":
        student, err := r.GetStudentProfile(ctx, id)
        if err == nil {
            result.Student = student
        }
    case "company":
        company, err := r.GetCompanyProfile(ctx, id)
        if err == nil {
            result.Company = company
        }
    }

    return result, nil
}

// GetUserByUsername retrieves a user by username (for authentication)
func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
    var user models.User
    err := r.db.QueryRow(ctx, `
        SELECT user_id, username, password, role
        FROM Users WHERE username = $1
    `, username).Scan(
        &user.UserID,
        &user.Username,
        &user.Password,
        &user.Role,
    )
    if err != nil {
        return user, fmt.Errorf("user not found: %w", err)
    }
    return user, nil
}

// GetAllUsers retrieves all users with their profiles
func (r *userRepository) GetAllUsers(ctx context.Context) ([]models.UserWithProfile, error) {
    rows, err := r.db.Query(ctx, `
        SELECT user_id, username, password, role FROM Users
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.UserWithProfile
    for rows.Next() {
        var u models.UserWithProfile
        if err := rows.Scan(&u.UserID, &u.Username, &u.Password, &u.Role); err != nil {
            return nil, err
        }

        // Get role-specific profile
        switch u.Role {
        case "teacher":
            teacher, err := r.GetTeacherProfile(ctx, u.UserID)
            if err == nil {
                u.Teacher = teacher
            }
        case "student":
            student, err := r.GetStudentProfile(ctx, u.UserID)
            if err == nil {
                u.Student = student
            }
        case "company":
            company, err := r.GetCompanyProfile(ctx, u.UserID)
            if err == nil {
                u.Company = company
            }
        }

        users = append(users, u)
    }

    return users, nil
}

// CreateUser inserts a new user into the Users table
func (r *userRepository) CreateUser(ctx context.Context, user models.User) (int, error) {
    var userID int
    err := r.db.QueryRow(ctx, `
        INSERT INTO Users (username, password, role)
        VALUES ($1, $2, $3)
        RETURNING user_id
    `, user.Username, user.Password, user.Role).Scan(&userID)

    if err != nil {
        return 0, fmt.Errorf("failed to create user: %w", err)
    }

    return userID, nil
}

// UpdateUser updates a user's base information
func (r *userRepository) UpdateUser(ctx context.Context, id int, user models.User) error {
    _, err := r.db.Exec(ctx, `
        UPDATE Users
        SET username = $1, password = $2, role = $3
        WHERE user_id = $4
    `, user.Username, user.Password, user.Role, id)

    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }

    return nil
}

// DeleteUser deletes a user (cascade will delete profile)
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
    _, err := r.db.Exec(ctx, `DELETE FROM Users WHERE user_id = $1`, id)
    if err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }
    return nil
}

// CreateTeacherProfile creates a teacher profile
func (r *userRepository) CreateTeacherProfile(ctx context.Context, teacherID int, teacher models.Teacher) error {
    _, err := r.db.Exec(ctx, `
        INSERT INTO Teacher (teacher_id, name_th, name_en, dept_id, email_contact, profile_image)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, teacherID, teacher.NameTH, teacher.NameEN, teacher.DeptID, teacher.EmailContact, teacher.ProfileImage)

    if err != nil {
        return fmt.Errorf("failed to create teacher profile: %w", err)
    }
    return nil
}

// CreateStudentProfile creates a student profile
func (r *userRepository) CreateStudentProfile(ctx context.Context, studentID int, student models.Student) error {
    _, err := r.db.Exec(ctx, `
        INSERT INTO Student (student_id, name_th, name_en, student_code, dept_id, year, education_level, profile_image)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `, studentID, student.NameTH, student.NameEN, student.StudentCode, student.DeptID, student.Year, student.EducationLevel, student.ProfileImage)

    if err != nil {
        return fmt.Errorf("failed to create student profile: %w", err)
    }
    return nil
}

// CreateCompanyProfile creates a company profile
func (r *userRepository) CreateCompanyProfile(ctx context.Context, companyID int, company models.Company) error {
    _, err := r.db.Exec(ctx, `
        INSERT INTO Company (company_id, company_name, status, profile_image)
        VALUES ($1, $2, $3, $4)
    `, companyID, company.CompanyName, company.Status, company.ProfileImage)

    if err != nil {
        return fmt.Errorf("failed to create company profile: %w", err)
    }
    return nil
}

// GetTeacherProfile retrieves a teacher profile
func (r *userRepository) GetTeacherProfile(ctx context.Context, teacherID int) (*models.Teacher, error) {
    var teacher models.Teacher
    err := r.db.QueryRow(ctx, `
        SELECT teacher_id, name_th, name_en, dept_id, email_contact, profile_image
        FROM Teacher WHERE teacher_id = $1
    `, teacherID).Scan(
        &teacher.TeacherID,
        &teacher.NameTH,
        &teacher.NameEN,
        &teacher.DeptID,
        &teacher.EmailContact,
        &teacher.ProfileImage,
    )

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get teacher profile: %w", err)
    }

    return &teacher, nil
}

// GetStudentProfile retrieves a student profile
func (r *userRepository) GetStudentProfile(ctx context.Context, studentID int) (*models.Student, error) {
    var student models.Student
    err := r.db.QueryRow(ctx, `
        SELECT student_id, name_th, name_en, student_code, dept_id, year, education_level, profile_image
        FROM Student WHERE student_id = $1
    `, studentID).Scan(
        &student.StudentID,
        &student.NameTH,
        &student.NameEN,
        &student.StudentCode,
        &student.DeptID,
        &student.Year,
        &student.EducationLevel,
        &student.ProfileImage,
    )

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get student profile: %w", err)
    }

    return &student, nil
}

// GetCompanyProfile retrieves a company profile
func (r *userRepository) GetCompanyProfile(ctx context.Context, companyID int) (*models.Company, error) {
    var company models.Company
    err := r.db.QueryRow(ctx, `
        SELECT company_id, company_name, status, profile_image
        FROM Company WHERE company_id = $1
    `, companyID).Scan(
        &company.CompanyID,
        &company.CompanyName,
        &company.Status,
        &company.ProfileImage,
    )

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("failed to get company profile: %w", err)
    }

    return &company, nil
}

// UpdateTeacherProfile updates a teacher profile
func (r *userRepository) UpdateTeacherProfile(ctx context.Context, teacherID int, teacher models.Teacher) error {
    _, err := r.db.Exec(ctx, `
        UPDATE Teacher
        SET name_th = $1, name_en = $2, dept_id = $3, email_contact = $4, profile_image = $5
        WHERE teacher_id = $6
    `, teacher.NameTH, teacher.NameEN, teacher.DeptID, teacher.EmailContact, teacher.ProfileImage, teacherID)

    if err != nil {
        return fmt.Errorf("failed to update teacher profile: %w", err)
    }
    return nil
}

// UpdateStudentProfile updates a student profile
func (r *userRepository) UpdateStudentProfile(ctx context.Context, studentID int, student models.Student) error {
    _, err := r.db.Exec(ctx, `
        UPDATE Student
        SET name_th = $1, name_en = $2, student_code = $3, dept_id = $4, year = $5, education_level = $6, profile_image = $7
        WHERE student_id = $8
    `, student.NameTH, student.NameEN, student.StudentCode, student.DeptID, student.Year, student.EducationLevel, student.ProfileImage, studentID)

    if err != nil {
        return fmt.Errorf("failed to update student profile: %w", err)
    }
    return nil
}

// UpdateCompanyProfile updates a company profile
func (r *userRepository) UpdateCompanyProfile(ctx context.Context, companyID int, company models.Company) error {
    _, err := r.db.Exec(ctx, `
        UPDATE Company
        SET company_name = $1, status = $2, profile_image = $3
        WHERE company_id = $4
    `, company.CompanyName, company.Status, company.ProfileImage, companyID)

    if err != nil {
        return fmt.Errorf("failed to update company profile: %w", err)
    }
    return nil
}

// GetAllDepartments retrieves all departments
func (r *userRepository) GetAllDepartments(ctx context.Context) ([]models.Department, error) {
    rows, err := r.db.Query(ctx, `SELECT dept_id, dept_name FROM Department`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var departments []models.Department
    for rows.Next() {
        var dept models.Department
        if err := rows.Scan(&dept.DeptID, &dept.DeptName); err != nil {
            return nil, err
        }
        departments = append(departments, dept)
    }

    return departments, nil
}

// GetDepartmentByID retrieves a department by ID
func (r *userRepository) GetDepartmentByID(ctx context.Context, id int) (models.Department, error) {
    var dept models.Department
    err := r.db.QueryRow(ctx, `
        SELECT dept_id, dept_name FROM Department WHERE dept_id = $1
    `, id).Scan(&dept.DeptID, &dept.DeptName)

    if err != nil {
        return dept, fmt.Errorf("department not found: %w", err)
    }
    return dept, nil
}