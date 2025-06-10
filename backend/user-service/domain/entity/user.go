package entity

import (
	"time"

	"github.com/google/uuid"
)

// User 表示用户领域实体
type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser 创建新的用户实体
func NewUser(username, email, password string) *User {
	return &User{
		ID:        uuid.New(),
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Validate 验证用户实体的有效性
func (u *User) Validate() error {
	if len(u.Username) < 3 || len(u.Username) > 50 {
		return ErrInvalidUsername
	}
	if len(u.Email) == 0 {
		return ErrInvalidEmail
	}
	if len(u.Password) < 6 {
		return ErrInvalidPassword
	}
	return nil
}
