.PHONY: build run test docker-build docker-up docker-down clean help

# 变量定义
APP_NAME=shortlink
DOCKER_IMAGE=$(APP_NAME):latest

# 默认目标
.DEFAULT_GOAL := help

## build: 编译应用
build:
	@echo "Building application..."
	go build -o bin/$(APP_NAME) ./cmd/api

## run: 本地运行应用（需要Redis）
run:
	@echo "Running application..."
	go run ./cmd/api/main.go

## test: 运行测试
test:
	@echo "Running tests..."
	go test -v ./...

## docker-build: 构建Docker镜像
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

## docker-up: 启动Docker Compose服务
docker-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d
	@echo "Services started! API available at http://localhost:8080"

## docker-down: 停止Docker Compose服务
docker-down:
	@echo "Stopping services..."
	docker-compose down

## docker-logs: 查看服务日志
docker-logs:
	docker-compose logs -f

## clean: 清理构建产物
clean:
	@echo "Cleaning..."
	rm -rf bin/
	docker-compose down -v

## lint: 运行代码检查
lint:
	@echo "Running linter..."
	golangci-lint run

## help: 显示帮助信息
help: Makefile
	@echo "Available targets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
