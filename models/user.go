package models

import "time"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}


type Profile struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type Role struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type Permission struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type RolePermission struct {
	ID            int64  `json:"id"`
	RoleID        int64  `json:"role_id"`
	PermissionID  int64  `json:"permission_id"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
}

type UserRole struct {
	ID            int64  `json:"id"`
	UserID        int64  `json:"user_id"`
	RoleID        int64  `json:"role_id"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
}
