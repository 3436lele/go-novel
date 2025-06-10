package service

import (
	"context"

	"github.com/3436lele/go-novel/backend/user-service/domain/service"
)

// UserApplicationService 用户应用服务
type UserApplicationService struct {
	userDomainService *service.UserService
}

// NewUserApplicationService 创建用户应用服务实例
func NewUserApplicationService(userDomainService *service.UserService) *UserApplicationService {
	return &UserApplicationService{
		userDomainService: userDomainService,
	}
}

// RegisterUserRequest 注册用户请求
type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterUserResponse 注册用户响应
type RegisterUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// RegisterUser 注册新用户
func (s *UserApplicationService) RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
	user, err := s.userDomainService.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &RegisterUserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
