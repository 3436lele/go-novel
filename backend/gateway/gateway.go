package gateway

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/3436lele/go-novel/backend/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Gateway 表示API网关
type Gateway struct {
	engine   *gin.Engine
	server   *http.Server
	config   *config.Config
	logger   *slog.Logger
	services map[string]Service
}

// Service 表示微服务接口
type Service interface {
	RegisterRoutes(router *gin.RouterGroup)
	GetName() string
}

// NewGateway 创建一个新的API网关实例
func NewGateway(cfg *config.Config) *Gateway {
	log := slog.With("domain", "gateway")

	gin.SetMode(cfg.Server.Mode)
	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// CORS配置
	corsConfig := cors.Config{
		AllowMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete,
			http.MethodOptions, http.MethodHead, http.MethodPatch,
		},
		AllowHeaders: []string{
			"Authorization", "Origin", "Content-Length", "Content-Type",
			"X-Requested-With", "Accept", "X-CSRF-Token",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	if cfg.Server.Mode == "debug" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowAllOrigins = false
		corsConfig.AllowOriginFunc = func(origin string) bool {
			// TODO: 实现允许的域名白名单
			return true
		}
	}

	engine.Use(cors.New(corsConfig))

	gateway := &Gateway{
		engine:   engine,
		config:   cfg,
		logger:   log,
		services: make(map[string]Service),
		server: &http.Server{
			Handler:      engine,
			Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}

	// 注册全局中间件
	gateway.registerGlobalMiddleware()

	return gateway
}

// RegisterService 注册微服务
func (g *Gateway) RegisterService(service Service) {
	g.services[service.GetName()] = service
	service.RegisterRoutes(g.engine.Group("/api/v1/" + service.GetName()))
	g.logger.Info("Service registered", "name", service.GetName())
}

// Start 启动API网关
func (g *Gateway) Start() error {
	g.logger.Info("Starting API Gateway", "address", g.server.Addr)
	return g.server.ListenAndServe()
}

// registerGlobalMiddleware 注册全局中间件
func (g *Gateway) registerGlobalMiddleware() {
	// 请求ID中间件
	g.engine.Use(func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	})

	// 认证中间件
	g.engine.Use(func(c *gin.Context) {
		// TODO: 实现JWT认证
		c.Next()
	})

	// 限流中间件
	g.engine.Use(func(c *gin.Context) {
		// TODO: 实现限流
		c.Next()
	})
}

// generateRequestID 生成请求ID
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString 生成随机字符串
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}
