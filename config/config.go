package config

import (
	"os"
)

// Config 应用配置
type Config struct {
	ServerPort  string // 服务器端口
	RedisAddr   string // Redis地址
	RedisPasswd string // Redis密码
	RedisDB     int    // Redis数据库编号
	BaseURL     string // 短链接基础URL
}

// LoadConfig 加载配置
func LoadConfig() *Config {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		RedisAddr:   getEnv("REDIS_ADDR", "redis:6379"),
		RedisPasswd: getEnv("REDIS_PASSWORD", ""),
		RedisDB:     0,
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
