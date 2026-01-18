# 短链接生成服务 (URL Shortener)

一个基于 Golang + Redis 实现的短链接生成服务，支持长链接转短链接和自动重定向功能。

## 功能特性

- ✅ 长链接转短链接（POST API）
- ✅ 短链接自动重定向到原始URL（GET）
- ✅ 基于Redis的高性能存储
- ✅ Docker容器化部署
- ✅ 访问次数统计
- ✅ 健康检查接口

## 技术栈

- **语言**: Go 1.21
- **Web框架**: Gin
- **存储**: Redis
- **容器化**: Docker + Docker Compose

## 项目结构

```
shortLink/
├── cmd/
│   └── api/
│       └── main.go              # 应用入口
├── internal/
│   ├── handler/                 # HTTP处理器层
│   │   └── shortlink.go
│   ├── service/                 # 业务逻辑层
│   │   └── shortlink.go
│   ├── repository/              # 数据访问层
│   │   └── redis.go
│   └── model/                   # 数据模型
│       └── shortlink.go
├── config/
│   └── config.go                # 配置管理
├── docker-compose.yml           # Docker编排文件
├── Dockerfile                   # API服务镜像
├── go.mod                       # Go模块依赖
└── README.md                    # 项目文档
```

## 快速开始

### 前置要求

- Docker 20.10+
- Docker Compose 2.0+

### 一键启动

```bash
# 克隆项目
git clone <repository-url>
cd shortLink

# 编译二进制文件（首次启动需要）
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/shortlink ./cmd/api

# 启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f
```

> **注意**: 本项目使用本地编译方案，需要先编译二进制文件再启动 Docker。如果你是 x86 架构，请将 GOARCH 改为 amd64。

服务将在以下端口启动：
- API服务: `http://localhost:8080`
- Redis: `localhost:6379`

### 验证服务

```bash
# 健康检查
curl http://localhost:8080/health
```

## API 使用说明

### 1. 创建短链接

**请求**:
```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.example.com/very/long/url/path"}'
```

**响应**:
```json
{
  "short_code": "aBcDeF1",
  "short_url": "http://localhost:8080/aBcDeF1"
}
```

### 2. 访问短链接（自动重定向）

**请求**:
```bash
curl -L http://localhost:8080/aBcDeF1
```

浏览器访问 `http://localhost:8080/aBcDeF1` 将自动重定向到原始URL。

### 3. 健康检查

**请求**:
```bash
curl http://localhost:8080/health
```

**响应**:
```json
{
  "status": "ok"
}
```

## 环境变量配置

可以通过环境变量自定义配置：

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| `SERVER_PORT` | API服务端口 | `8080` |
| `REDIS_ADDR` | Redis地址 | `redis:6379` |
| `REDIS_PASSWORD` | Redis密码 | `` |
| `BASE_URL` | 短链接基础URL | `http://localhost:8080` |

修改 `docker-compose.yml` 中的 `environment` 部分即可：

```yaml
environment:
  - SERVER_PORT=8080
  - REDIS_ADDR=redis:6379
  - BASE_URL=https://your-domain.com
```

## 本地开发

### 方式 1: 使用 Docker（推荐）

```bash
# 1. 编译二进制文件
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/shortlink ./cmd/api

# 2. 启动服务
docker-compose up -d
```

### 方式 2: 本地直接运行

```bash
# 1. 安装依赖
go mod download

# 2. 启动 Redis
docker run -d -p 6379:6379 redis:7-alpine

# 3. 设置环境变量并运行
export REDIS_ADDR=localhost:6379
export BASE_URL=http://localhost:8080
go run cmd/api/main.go
```

## 停止服务

```bash
# 停止并删除容器
docker-compose down

# 停止并删除容器及数据卷
docker-compose down -v
```

## 短链接生成算法

1. 对原始URL进行MD5哈希
2. 取哈希值的前6个字节
3. 进行Base64 URL编码
4. 截取前7个字符作为短链接代码
5. 如果冲突，添加时间戳重新生成

## 数据持久化

Redis数据通过Docker volume持久化存储：
- 数据卷名称: `shortlink_redis-data`
- 短链接有效期: 1年（可在代码中调整）

## 性能特点

- ✅ Redis内存存储，响应速度快
- ✅ 支持高并发访问
- ✅ 自动记录访问次数
- ✅ 容器化部署，易于扩展

## 测试示例

```bash
# 创建短链接
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'

# 输出: {"short_code":"abc123","short_url":"http://localhost:8080/abc123"}

# 测试重定向
curl -I http://localhost:8080/abc123

# 输出: HTTP/1.1 302 Found
#       Location: https://github.com
```

## 故障排查

### 1. 服务无法启动

```bash
# 查看日志
docker-compose logs api

# 检查Redis连接
docker-compose logs redis
```

### 2. 端口冲突

修改 `docker-compose.yml` 中的端口映射：
```yaml
ports:
  - "8081:8080"  # 将主机端口改为8081
```

### 3. Redis连接失败

确保Redis服务健康：
```bash
docker-compose ps
docker-compose exec redis redis-cli ping
```

## License

MIT License

## 作者

罗文
