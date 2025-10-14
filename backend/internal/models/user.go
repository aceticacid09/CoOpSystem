package models

// Users table
type User struct {
	UserID   int    `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"` // Password hashed
	Role     string `db:"role" json:"role"`  // role_type: teacher, student, company
}

// Teacher table
type Teacher struct {
	TeacherID    int     `db:"teacher_id" json:"teacher_id"`
	NameTH       *string `db:"name_th" json:"name_th,omitempty"`
	NameEN       *string `db:"name_en" json:"name_en,omitempty"`
	DeptID       *int    `db:"dept_id" json:"dept_id,omitempty"`
	EmailContact *string `db:"email_contact" json:"email_contact,omitempty"`
	ProfileImage *string `db:"profile_image" json:"profile_image,omitempty"`
}

// Student table
type Student struct {
	StudentID      int     `db:"student_id" json:"student_id"`
	NameTH         *string `db:"name_th" json:"name_th,omitempty"`
	NameEN         *string `db:"name_en" json:"name_en,omitempty"`
	StudentCode    int     `db:"student_code" json:"student_code"`
	DeptID         *int    `db:"dept_id" json:"dept_id,omitempty"`
	Year           *int    `db:"year" json:"year,omitempty"`
	EducationLevel *string `db:"education_level" json:"education_level,omitempty"`
	ProfileImage   *string `db:"profile_image" json:"profile_image,omitempty"`
}

// Company table
type Company struct {
	CompanyID    int     `db:"company_id" json:"company_id"`
	CompanyName  string  `db:"company_name" json:"company_name"`
	Status       string  `db:"status" json:"status"` // company_status: active, inactive, pending
	ProfileImage *string `db:"profile_image" json:"profile_image,omitempty"`
}

// Department table
type Department struct {
	DeptID   int    `db:"dept_id" json:"dept_id"`
	DeptName string `db:"dept_name" json:"dept_name"`
}

// UserWithProfile combines User with their role-specific profile
type UserWithProfile struct {
	User
	Teacher *Teacher `json:"teacher,omitempty"`
	Student *Student `json:"student,omitempty"`
	Company *Company `json:"company,omitempty"`
}

// CreateUserRequest for user registration
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"` // admin, teacher, student, company

	// Teacher-specific fields
	TeacherNameTH       *string `json:"teacher_name_th,omitempty"`
	TeacherNameEN       *string `json:"teacher_name_en,omitempty"`
	TeacherDeptID       *int    `json:"teacher_dept_id,omitempty"`
	TeacherEmailContact *string `json:"teacher_email_contact,omitempty"`
	TeacherProfileImage *string `json:"teacher_profile_image,omitempty"`

	// Student-specific fields
	StudentNameTH         *string `json:"student_name_th,omitempty"`
	StudentNameEN         *string `json:"student_name_en,omitempty"`
	StudentCode           *int    `json:"student_code,omitempty"`
	StudentDeptID         *int    `json:"student_dept_id,omitempty"`
	StudentYear           *int    `json:"student_year,omitempty"`
	StudentEducationLevel *string `json:"student_education_level,omitempty"`
	StudentProfileImage   *string `json:"student_profile_image,omitempty"`

	// Company-specific fields
	CompanyName         *string `json:"company_name,omitempty"`
	CompanyStatus       *string `json:"company_status,omitempty"`
	CompanyProfileImage *string `json:"company_profile_image,omitempty"`
}

// UpdateUserRequest for updating user information
type UpdateUserRequest struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Role     *string `json:"role,omitempty"`

	// Profile-specific fields (based on role)
	NameTH         *string `json:"name_th,omitempty"`
	NameEN         *string `json:"name_en,omitempty"`
	DeptID         *int    `json:"dept_id,omitempty"`
	EmailContact   *string `json:"email_contact,omitempty"`
	ProfileImage   *string `json:"profile_image,omitempty"`
	StudentCode    *int    `json:"student_code,omitempty"`
	Year           *int    `json:"year,omitempty"`
	EducationLevel *string `json:"education_level,omitempty"`
	CompanyName    *string `json:"company_name,omitempty"`
	CompanyStatus  *string `json:"company_status,omitempty"`
}

// LoginRequest for authentication
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse for authentication response
type LoginResponse struct {
	Token        string      `json:"token"`
	RefreshToken string      `json:"refresh_token"`
	UserID       int         `json:"user_id"`
	Role         string      `json:"role"`
	Profile      interface{} `json:"profile,omitempty"` // Teacher, Student, or Company
	Message      string      `json:"message"`
}
