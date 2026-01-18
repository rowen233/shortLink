package service

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/lwy/shortlink/internal/model"
	"github.com/lwy/shortlink/internal/repository"
)

// ShortLinkService 短链接服务
type ShortLinkService struct {
	repo    *repository.RedisRepository
	baseURL string // 短链接的基础URL，例如 http://localhost:8080
}

// NewShortLinkService 创建短链接服务实例
func NewShortLinkService(repo *repository.RedisRepository, baseURL string) *ShortLinkService {
	return &ShortLinkService{
		repo:    repo,
		baseURL: baseURL,
	}
}

// CreateShortLink 创建短链接
func (s *ShortLinkService) CreateShortLink(ctx context.Context, originalURL string) (*model.CreateShortLinkResponse, error) {
	// 生成短链接代码
	shortCode := s.generateShortCode(originalURL)

	// 如果短链接已存在，尝试添加时间戳重新生成
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		exists, err := s.repo.ShortCodeExists(ctx, shortCode)
		if err != nil {
			return nil, fmt.Errorf("failed to check short code existence: %w", err)
		}

		if !exists {
			break
		}

		// 添加时间戳重新生成
		shortCode = s.generateShortCode(fmt.Sprintf("%s-%d", originalURL, time.Now().UnixNano()))

		if i == maxRetries-1 {
			return nil, fmt.Errorf("failed to generate unique short code after %d retries", maxRetries)
		}
	}

	// 创建短链接对象
	shortLink := &model.ShortLink{
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
		VisitCount:  0,
	}

	// 保存到存储
	if err := s.repo.SaveShortLink(ctx, shortLink); err != nil {
		return nil, fmt.Errorf("failed to save short link: %w", err)
	}

	return &model.CreateShortLinkResponse{
		ShortCode: shortCode,
		ShortURL:  fmt.Sprintf("%s/%s", s.baseURL, shortCode),
	}, nil
}

// GetOriginalURL 获取原始URL
func (s *ShortLinkService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	url, err := s.repo.GetOriginalURL(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("failed to get original URL: %w", err)
	}
	return url, nil
}

// generateShortCode 生成短链接代码
// 使用MD5哈希后取前8位，然后进行base64编码
func (s *ShortLinkService) generateShortCode(url string) string {
	// 计算MD5哈希
	hash := md5.Sum([]byte(url))

	// 取前6个字节进行base64编码
	encoded := base64.URLEncoding.EncodeToString(hash[:6])

	// 移除base64编码中的特殊字符，只保留字母和数字
	// 取前7个字符作为短链接代码
	if len(encoded) > 7 {
		encoded = encoded[:7]
	}

	return encoded
}
