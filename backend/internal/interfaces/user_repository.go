package interfaces

import (
	"context"
	"coop/internal/models"
)

type UserRepository interface {
	// User CRUD operations
	GetUserByID(ctx context.Context, id int) (models.UserWithProfile, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.UserWithProfile, error)
	CreateUser(ctx context.Context, user models.User) (int, error)
	UpdateUser(ctx context.Context, id int, user models.User) error
	DeleteUser(ctx context.Context, id int) error

	// Profile operations
	CreateTeacherProfile(ctx context.Context, teacherID int, teacher models.Teacher) error
	CreateStudentProfile(ctx context.Context, studentID int, student models.Student) error
	CreateCompanyProfile(ctx context.Context, companyID int, company models.Company) error

	GetTeacherProfile(ctx context.Context, teacherID int) (*models.Teacher, error)
	GetStudentProfile(ctx context.Context, studentID int) (*models.Student, error)
	GetCompanyProfile(ctx context.Context, companyID int) (*models.Company, error)

	UpdateTeacherProfile(ctx context.Context, teacherID int, teacher models.Teacher) error
	UpdateStudentProfile(ctx context.Context, studentID int, student models.Student) error
	UpdateCompanyProfile(ctx context.Context, companyID int, company models.Company) error

	// Department operations
	GetAllDepartments(ctx context.Context) ([]models.Department, error)
	GetDepartmentByID(ctx context.Context, id int) (models.Department, error)
}
