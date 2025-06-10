package persistence

import (
	"context"

	"github.com/3436lele/go-novel/backend/user-service/domain/entity"
	"github.com/3436lele/go-novel/backend/user-service/domain/repository"
	"gorm.io/gorm"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 创建用户 DAO 实例
func NewUserDAO(db *gorm.DB) repository.UserRepository {
	return &UserDAO{
		db: db,
	}
}

// Create 创建新用户
func (dao *UserDAO) Create(ctx context.Context, user *entity.User) error {
	return dao.db.WithContext(ctx).Create(user).Error
}

// FindByUsername 根据用户名查找用户
func (dao *UserDAO) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := dao.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ExistsByUsername 检查用户名是否存在
func (dao *UserDAO) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&entity.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// ExistsByEmail 检查邮箱是否存在
func (dao *UserDAO) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := dao.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
