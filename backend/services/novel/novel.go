package novel

import (
	"github.com/gin-gonic/gin"
)

// NovelService 小说服务
type NovelService struct {
	// TODO: 添加依赖注入
}

// NewNovelService 创建小说服务实例
func NewNovelService() *NovelService {
	return &NovelService{}
}

// GetName 获取服务名称
func (s *NovelService) GetName() string {
	return "novels"
}

// RegisterRoutes 注册路由
func (s *NovelService) RegisterRoutes(router *gin.RouterGroup) {
	// 小说列表
	router.GET("", s.ListNovels)
	// 小说详情
	router.GET("/:id", s.GetNovel)
	// 创建小说
	router.POST("", s.CreateNovel)
	// 更新小说
	router.PUT("/:id", s.UpdateNovel)
	// 删除小说
	router.DELETE("/:id", s.DeleteNovel)
}

// ListNovels 获取小说列表
func (s *NovelService) ListNovels(c *gin.Context) {
	// TODO: 实现获取小说列表逻辑
	c.JSON(200, gin.H{
		"message": "获取小说列表",
	})
}

// GetNovel 获取小说详情
func (s *NovelService) GetNovel(c *gin.Context) {
	id := c.Param("id")
	// TODO: 实现获取小说详情逻辑
	c.JSON(200, gin.H{
		"message": "获取小说详情",
		"id":      id,
	})
}

// CreateNovel 创建小说
func (s *NovelService) CreateNovel(c *gin.Context) {
	// TODO: 实现创建小说逻辑
	c.JSON(201, gin.H{
		"message": "创建小说",
	})
}

// UpdateNovel 更新小说
func (s *NovelService) UpdateNovel(c *gin.Context) {
	id := c.Param("id")
	// TODO: 实现更新小说逻辑
	c.JSON(200, gin.H{
		"message": "更新小说",
		"id":      id,
	})
}

// DeleteNovel 删除小说
func (s *NovelService) DeleteNovel(c *gin.Context) {
	id := c.Param("id")
	// TODO: 实现删除小说逻辑
	c.JSON(200, gin.H{
		"message": "删除小说",
		"id":      id,
	})
}
