package api

import (
	"net/http"

	service "github.com/3436lele/go-novel/backend/user-service/application"
	"github.com/gin-gonic/gin"
)

// UserHandler 用户 HTTP 处理器
type UserHandler struct {
	userAppService *service.UserApplicationService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(userAppService *service.UserApplicationService) *UserHandler {
	return &UserHandler{
		userAppService: userAppService,
	}
}

// Register 处理用户注册请求
func (h *UserHandler) Register(c *gin.Context) {
	var req service.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的请求参数",
		})
		return
	}

	resp, err := h.userAppService.RegisterUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
