package repository

import (
	"context"

	"github.com/3436lele/go-novel/backend/user-service/domain/entity"
)

// UserRepository 定义用户仓储接口
type UserRepository interface {
	// Create 创建新用户
	Create(ctx context.Context, user *entity.User) error
	// FindByUsername 根据用户名查找用户
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	// FindByEmail 根据邮箱查找用户
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	// ExistsByUsername 检查用户名是否存在
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	// ExistsByEmail 检查邮箱是否存在
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
