# 项目开发日志

## 项目信息

- **项目名称**: 短链接生成服务 (URL Shortener)
- **开发时间**: 2026年1月18日
- **开发者**: 罗文
- **仓库地址**: https://github.com/rowen233/shortLink
- **技术栈**: Golang + Redis + Docker + GitHub Actions

---

## 项目概览

### 核心功能

- ✅ RESTful API 实现（POST 创建短链接 / GET 自动重定向）
- ✅ Redis 高性能存储
- ✅ Docker 容器化部署
- ✅ GitHub Actions 自动化测试
- ✅ 完整的工程化实践

### 技术亮点

1. **分层架构设计** - Handler / Service / Repository / Model 清晰分层
2. **Docker 多阶段构建** - 镜像体积优化，构建效率高
3. **CI/CD 自动化** - GitHub Actions 实现自动构建、测试、部署
4. **高性能存储** - Redis 内存存储，支持原子操作和自动过期

---

## 架构设计

### 系统架构

```
┌─────────────────────────────────────┐
│         Handler Layer               │  HTTP 请求处理
│    (Gin Web Framework)              │
├─────────────────────────────────────┤
│         Service Layer               │  业务逻辑
│    (Short Link Generation)          │
├─────────────────────────────────────┤
│       Repository Layer              │  数据访问
│      (Redis Operations)             │
├─────────────────────────────────────┤
│         Model Layer                 │  数据模型
└─────────────────────────────────────┘
```

### 核心算法

**短链接生成算法**：
```go
MD5(URL) → 取前6字节 → Base64 URL编码 → 截取7字符
理论容量: 64^7 ≈ 4.4万亿
```

**数据存储设计**：
- Redis Hash 结构存储
- Key: `shortlink:{shortCode}`
- 自动过期机制（365天）
- 原子操作支持访问计数

---

## 技术实现

### 1. 后端服务

**技术选型**：
- **Golang 1.21** - 高性能、并发支持
- **Gin Framework** - 轻量级 Web 框架
- **Redis 7** - 内存存储，毫秒级响应

**核心代码结构**：
```
cmd/api/              # 应用入口
internal/
  ├── handler/        # HTTP 路由处理
  ├── service/        # 业务逻辑
  ├── repository/     # 数据访问
  └── model/          # 数据模型
config/               # 配置管理
```

### 2. 容器化部署

**Docker 多阶段构建**：
- 构建阶段：使用 golang:1.21-alpine 编译
- 运行阶段：使用 alpine:latest 最小化镜像
- 镜像体积优化，安全性提升

**Docker Compose 编排**：
- Redis 服务：持久化存储 + 健康检查
- API 服务：依赖管理 + 自动重启
- 网络隔离：独立网络命名空间

### 3. CI/CD 流程

**GitHub Actions 工作流**：
- ✅ 代码构建与测试
- ✅ Docker 镜像构建
- ✅ 集成测试验证
- ✅ 代码质量检查（gofmt, govet）

---

## API 设计

### 端点说明

| 方法 | 路径 | 说明 | 响应 |
|------|------|------|------|
| POST | `/api/shorten` | 创建短链接 | 返回短码和完整URL |
| GET | `/:shortCode` | 重定向 | 302重定向到原始URL |
| GET | `/health` | 健康检查 | 服务状态 |

### 使用示例

```bash
# 创建短链接
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'

# 响应
{
  "short_code": "MJf8qbH",
  "short_url": "http://localhost:8080/MJf8qbH"
}

# 访问短链接（自动重定向）
curl -L http://localhost:8080/MJf8qbH
```

---

## 项目成果

### 代码统计

```
总代码行数: 525 行
├── Go 源代码: 9 个文件
├── 单元测试: 3 个文件
├── 配置文件: 7 个
└── 文档: README.md
```

### 文件结构

```
shortLink/
├── .github/workflows/    # CI/CD 配置
├── cmd/api/             # 应用入口
├── config/              # 配置管理
├── internal/
│   ├── handler/        # HTTP 层
│   ├── service/        # 业务层
│   ├── repository/     # 数据层
│   └── model/          # 模型层
├── docker-compose.yml   # Docker 编排
├── Dockerfile          # 镜像构建
├── Makefile           # 构建脚本
└── README.md          # 项目文档
```

### 测试覆盖

```
✓ config/           - 配置加载测试
✓ internal/model/   - 数据模型测试
✓ internal/service/ - 业务逻辑测试
```

---

## 技术决策

### 为什么选择 Redis？

- **高性能**：内存存储，微秒级延迟
- **原子操作**：HINCRBY 支持并发访问计数
- **过期机制**：自动清理过期数据
- **高可用**：支持主从复制和集群

### 为什么使用三层架构？

- **职责清晰**：每层专注单一职责
- **易于测试**：层与层之间可独立测试
- **便于扩展**：更换实现不影响其他层
- **符合最佳实践**：SOLID 原则

### 为什么采用 Docker 多阶段构建？

- **镜像优化**：运行阶段只包含必要文件
- **安全性**：减少攻击面
- **用户友好**：无需本地安装 Go 环境
- **标准化**：符合云原生最佳实践

---

## 核心特性

### 1. 工程化完整

- ✅ 清晰的分层架构
- ✅ 完整的单元测试
- ✅ CI/CD 自动化
- ✅ 代码质量检查
- ✅ 详细的文档

### 2. 生产就绪

- ✅ 错误处理完善
- ✅ 配置管理规范
- ✅ 日志记录完整
- ✅ 健康检查接口
- ✅ 优雅关闭支持

### 3. 性能优化

- ✅ Redis 内存存储
- ✅ Pipeline 批量操作
- ✅ 连接池管理
- ✅ 并发安全

### 4. 用户友好

- ✅ 一键启动部署
- ✅ 无需安装开发环境
- ✅ 文档清晰详细
- ✅ 快速上手

---

## 部署方式

### 用户使用流程

```bash
# 1. 克隆项目
git clone https://github.com/rowen233/shortLink.git
cd shortLink

# 2. 一键启动（仅需 Docker）
docker-compose up -d

# 3. 验证服务
curl http://localhost:8080/health
```

### 特点

- **零依赖**：无需安装 Go、Redis 等
- **一键启动**：docker-compose up -d
- **自动构建**：Docker 内完成编译
- **数据持久化**：自动配置 Volume

---

## 性能指标

### 容量

- **短码长度**: 7 字符
- **字符集**: Base64 (64个字符)
- **理论容量**: 64^7 ≈ 4.4 万亿

### 性能

- **存储**: Redis 内存，微秒级延迟
- **并发**: Go 协程支持高并发
- **扩展**: 无状态设计，易于水平扩展

---

## 总结

### 项目价值

本项目展示了：
- ✅ 完整的 Go Web 服务开发
- ✅ 标准的 Docker 容器化实践
- ✅ 规范的工程化开发流程
- ✅ 实用的短链接服务实现

### 技术亮点

1. **架构设计** - 清晰的分层，职责明确
2. **工程实践** - CI/CD、测试、文档完整
3. **容器化** - Docker 多阶段构建优化
4. **生产就绪** - 错误处理、日志、监控齐全

### 适用场景

- 学习 Go Web 开发
- 了解 Docker 容器化
- 实践 CI/CD 流程
- 搭建短链接服务

---

**项目状态**: ✅ 已完成并通过所有 CI 检查

**仓库地址**: https://github.com/rowen233/shortLink

**开发者**: 罗文

**完成时间**: 2026-01-18
