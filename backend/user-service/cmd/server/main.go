package main

import (
	"fmt"
	"log"

	"github.com/3436lele/go-novel/backend/user-service/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, err := config.InitConfig("config/config.dev.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// // 自动迁移数据库表
	// if err := db.AutoMigrate(&entity.User{}); err != nil {
	// 	log.Fatalf("Failed to migrate database: %v", err)
	// }

	// // 初始化依赖
	// userDAO := persistence.NewUserDAO(db)
	// userDomainService := service.NewUserService(userDAO)
	// userAppService := service.NewUserApplicationService(userDomainService)
	// userHandler := api.NewUserHandler(userAppService)

	// // 设置路由
	// r := gin.Default()

	// // 用户注册路由
	// r.POST("/api/v1/users/register", userHandler.Register)

	// // 启动服务器
	// port := 8080
	// fmt.Printf("Server is running on port %d...\n", port)
	// if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}
