package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lwy/shortlink/config"
	"github.com/lwy/shortlink/internal/handler"
	"github.com/lwy/shortlink/internal/repository"
	"github.com/lwy/shortlink/internal/service"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化Redis存储
	repo, err := repository.NewRedisRepository(cfg.RedisAddr, cfg.RedisPasswd, cfg.RedisDB)
	if err != nil {
		log.Fatalf("Failed to initialize Redis repository: %v", err)
	}
	defer func() {
		if err := repo.Close(); err != nil {
			log.Printf("Failed to close Redis connection: %v", err)
		}
	}()

	log.Println("Successfully connected to Redis")

	// 初始化服务层
	shortLinkService := service.NewShortLinkService(repo, cfg.BaseURL)

	// 初始化处理器
	shortLinkHandler := handler.NewShortLinkHandler(shortLinkService)

	// 设置Gin路由
	router := gin.Default()

	// 健康检查
	router.GET("/health", shortLinkHandler.HealthCheck)

	// API路由组
	api := router.Group("/api")
	{
		api.POST("/shorten", shortLinkHandler.CreateShortLink)
	}

	// 短链接重定向路由
	router.GET("/:shortCode", shortLinkHandler.RedirectToOriginal)

	// 启动服务器
	addr := "0.0.0.0:" + cfg.ServerPort
	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
