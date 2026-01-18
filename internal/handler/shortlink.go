package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwy/shortlink/internal/model"
	"github.com/lwy/shortlink/internal/service"
)

// ShortLinkHandler 短链接HTTP处理器
type ShortLinkHandler struct {
	service *service.ShortLinkService
}

// NewShortLinkHandler 创建短链接处理器实例
func NewShortLinkHandler(service *service.ShortLinkService) *ShortLinkHandler {
	return &ShortLinkHandler{
		service: service,
	}
}

// CreateShortLink 创建短链接
// POST /api/shorten
func (h *ShortLinkHandler) CreateShortLink(c *gin.Context) {
	var req model.CreateShortLinkRequest

	// 绑定并验证请求
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Invalid request: " + err.Error(),
		})
		return
	}

	// 调用服务层创建短链接
	resp, err := h.service.CreateShortLink(c.Request.Context(), req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Failed to create short link: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// RedirectToOriginal 重定向到原始URL
// GET /:shortCode
func (h *ShortLinkHandler) RedirectToOriginal(c *gin.Context) {
	shortCode := c.Param("shortCode")

	if shortCode == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Short code is required",
		})
		return
	}

	// 获取原始URL
	originalURL, err := h.service.GetOriginalURL(c.Request.Context(), shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error: "Short link not found",
		})
		return
	}

	// 302重定向到原始URL
	c.Redirect(http.StatusFound, originalURL)
}

// HealthCheck 健康检查
// GET /health
func (h *ShortLinkHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
