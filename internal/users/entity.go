package users

import "time"

type Role string

const (
	RoleSuperAdmin Role = "superadmin"
	RoleAdmin      Role = "admin"
	RoleUser       Role = "user"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Exclude from JSON response
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
	Role      Role      `json:"role"`
}
