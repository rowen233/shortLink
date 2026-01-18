package model

import "time"

// ShortLink 短链接数据模型
type ShortLink struct {
	ShortCode   string    `json:"short_code"`   // 短链接代码
	OriginalURL string    `json:"original_url"` // 原始长链接
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	VisitCount  int64     `json:"visit_count"`  // 访问次数
}

// CreateShortLinkRequest 创建短链接请求
type CreateShortLinkRequest struct {
	URL string `json:"url" binding:"required,url"` // 原始URL，必填且必须是有效URL
}

// CreateShortLinkResponse 创建短链接响应
type CreateShortLinkResponse struct {
	ShortCode string `json:"short_code"` // 生成的短链接代码
	ShortURL  string `json:"short_url"`  // 完整的短链接URL
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error"` // 错误信息
}
