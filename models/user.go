package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        int `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UUID      string     `gorm:"size:255;not null;" json:"uuid"`
	Username  string     `gorm:"size:15;not null" json:"username"`
	Password  string     `gorm:"size:255;not null" json:"password,omitempty"`
	RoleID    int        `sql:"-" json:"role_id,omitempty"`
	IsDelete  int        `gorm:"not null;size:2;default:1" json:"is_delete"`
}

type ResponseGetAllAdmin struct {
	TotalData   int        `json:"total_data"`
	DataPerPage int        `json:"data_per_page"`
	TotalPage   int        `json:"total_page"`
	FirstPage   int        `json:"first_page"`
	EndPage     int        `json:"end_page"`
	Data        []GetAdmin `json:"data"`
}

type GetAdmin struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Role     string `json:"role,omitempty"`
	IsDelete int    `gorm:"not null;size:2;default:1" json:"is_delete"`
}

// TableName ..
func (s User) TableName() string {
	return "tb_user"
}

// Decrypt ..
type Decrypt struct {
	Encrypt string `json:"encrypt"`
}

// Role ..
type Role struct {
	gorm.Model
	RoleName string `gorm:"not null" json:"role_name"`
	IsDelete int    `gorm:"not null;size:2;default:1" json:"is_delete"`
}

// UserRole ..
type UserRole struct {
	gorm.Model
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

// TableName ..
func (s Role) TableName() string {
	return "tb_role"
}

// TableName ..
func (s UserRole) TableName() string {
	return "tb_user_role"
}

// BeforeCreate will set a UUID rather than numeric UUID.
func (base *User) BeforeCreate(scope *gorm.Scope) error {
	base.UUID = uuid.NewV4().String()
	if base.UUID != "" {
		return nil
	}
	return scope.SetColumn("UUID", base.UUID)
}
