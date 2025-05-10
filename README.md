# Go-Novel 小说网站

一个基于Vue 3和Go的现代化小说网站项目。

## 项目特点

- 现代化的用户界面设计
- 响应式布局，支持移动端和桌面端
- 基于Vue 3的前端框架
- 使用Go语言构建的后端服务
- Element Plus UI组件库

## 技术栈

### 前端
- Vue 3
- Vue Router
- Element Plus
- ES6+
- Babel
- ESLint

### 后端
- Go
- 标准库
- 日志系统 (slog)

## 项目结构

```
go-novel/
├── frontend/          # Vue 3前端项目
│   ├── src/          # 源代码
│   ├── public/       # 静态资源
│   └── package.json  # 项目依赖配置
└── backend/          # Go后端项目
    ├── main.go       # 主程序入口
    └── go.mod        # Go模块配置
```

## 开发环境要求

- Node.js >= 14.0.0
- Go >= 1.16
- npm >= 6.0.0

## 安装说明

### 前端

1. 进入前端目录：
```bash
cd frontend
```

2. 安装依赖：
```bash
npm install
```

3. 启动开发服务器：
```bash
npm run serve
```

### 后端

1. 进入后端目录：
```bash
cd backend
```

2. 安装Go依赖：
```bash
go mod download
```

3. 运行后端服务：
```bash
go run main.go
```

## 开发指南

### 前端开发
- 使用 `npm run serve` 启动开发服务器
- 使用 `npm run build` 构建生产版本
- 使用 `npm run lint` 运行代码检查

### 后端开发
- 使用 `go run main.go` 运行开发服务器
- 使用 `go build` 构建可执行文件

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建Pull Request

## 许可证

[MIT License](LICENSE)
