package service

import (
	"context"
	"errors"

	"github.com/3436lele/go-novel/backend/user-service/domain/entity"
	"github.com/3436lele/go-novel/backend/user-service/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService 用户领域服务
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户领域服务实例
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register 注册新用户
func (s *UserService) Register(ctx context.Context, username, email, password string) (*entity.User, error) {
	// 检查用户名是否已存在
	exists, err := s.userRepo.ExistsByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	exists, err = s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("邮箱已被注册")
	}

	// 创建用户实体
	user := entity.NewUser(username, email, password)

	// 验证用户实体
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// 保存用户
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
