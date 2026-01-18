package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lwy/shortlink/internal/model"
)

// RedisRepository Redis存储实现
type RedisRepository struct {
	client *redis.Client
}

// NewRedisRepository 创建Redis存储实例
func NewRedisRepository(addr, password string, db int) (*RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisRepository{
		client: client,
	}, nil
}

// SaveShortLink 保存短链接
func (r *RedisRepository) SaveShortLink(ctx context.Context, shortLink *model.ShortLink) error {
	// 使用Hash存储短链接信息
	key := fmt.Sprintf("shortlink:%s", shortLink.ShortCode)

	pipe := r.client.Pipeline()
	pipe.HSet(ctx, key, map[string]interface{}{
		"original_url": shortLink.OriginalURL,
		"created_at":   shortLink.CreatedAt.Unix(),
		"visit_count":  shortLink.VisitCount,
	})

	// 设置过期时间为1年
	pipe.Expire(ctx, key, 365*24*time.Hour)

	_, err := pipe.Exec(ctx)
	return err
}

// GetOriginalURL 根据短链接代码获取原始URL
func (r *RedisRepository) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	key := fmt.Sprintf("shortlink:%s", shortCode)

	url, err := r.client.HGet(ctx, key, "original_url").Result()
	if err == redis.Nil {
		return "", fmt.Errorf("short link not found")
	}
	if err != nil {
		return "", err
	}

	// 增加访问次数
	r.client.HIncrBy(ctx, key, "visit_count", 1)

	return url, nil
}

// ShortCodeExists 检查短链接代码是否已存在
func (r *RedisRepository) ShortCodeExists(ctx context.Context, shortCode string) (bool, error) {
	key := fmt.Sprintf("shortlink:%s", shortCode)
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// Close 关闭Redis连接
func (r *RedisRepository) Close() error {
	return r.client.Close()
}
