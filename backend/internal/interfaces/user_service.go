package interfaces

import (
	"context"
	"coop/internal/models"
)

type UserService interface {
	// User services
	GetUserByID(ctx context.Context, id int) (models.UserWithProfile, error)
	GetAllUsers(ctx context.Context) ([]models.UserWithProfile, error)
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserWithProfile, error)
	UpdateUser(ctx context.Context, id int, req models.UpdateUserRequest) (models.UserWithProfile, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)

	// Authentication
	Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error)

	// Department services
	GetAllDepartments(ctx context.Context) ([]models.Department, error)
	GetDepartmentByID(ctx context.Context, id int) (models.Department, error)
}
