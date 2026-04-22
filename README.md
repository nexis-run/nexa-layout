# oos

基于DDD领域驱动设计的kratos模板

- https://juejin.cn/post/7298160530292703244
- https://github.com/KendoCross/kendoDDD
- https://github.com/vgocoder/go-ddd
- https://juejin.cn/post/7226556923238203429

### 常用命令

```bash
# 生成 dao
go run -mod=mod nexis.run/nexa/cmd/nexa@master new dao Name

# 生成 ent context
go run -mod=mod nexis.run/nexa/cmd/nexa@master new echoctx Name

# 添加 ent schema
go run -mod=mod nexis.run/nexa/cmd/nexa@master ent new Name

# 生成 ent 代码
go run -mod=mod nexis.run/nexa/cmd/nexa@master ent generate
```

### Commit 格式规范

> 参考文章 [Commit message 和 Change log 编写指南](https://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html)

```
[<type>](<scope>) <subject> (#pr)
docs：                   文档变动
fix：                    bug 修复
feat：                   新增功能
feat-wip：               开发中的功能，比如某功能的部分代码。
improvement：            原有功能的优化和改进
style：                  代码风格调整
typo：                   代码或文档勘误
refactor：               代码重构（不涉及功能变动）
performance/optimize：   性能优化
test：                   单元测试的添加或修复
chore：                  构建工具的修改
revert：                 回滚
deps：                   第三方依赖库的修改
community：              社区相关的修改，如修改 Github Issue 模板等。
```

### 文档工具

[Stoplight Elements](https://stoplight.io/open-source/elements)

### 基本结构

#### 整体架构图

```
┌───────────────────────────────────────────────────────────────────────────────────┐
│                                    项目架构图                                     │
├───────────────────────────────────────────────────────────────────────────────────┤
│                                                                                   │
│  ┌─────────────┐      ┌─────────────┐      ┌─────────────┐      ┌──────────────┐  │
│  │  HTTP/gRPC  │ ───▶ │   Handler   │ ───▶ │   Service   │ ◀─── │ Application  │  │
│  │  (外部请求) │      │   (接口层)  │      │   (业务层)  │      │  (其他服务)  │  │
│  └─────────────┘      └─────────────┘      └─────────────┘      └──────────────┘  │
│                              │                      │                             │
│                              ▼                      ▼                             │
│                       ┌──────────────────────────────────┐                        │
│                       │           DI Container           │                        │
│                       │      (依赖容器 / 单例管理)       │                        │                   
│                       └──────────────────────────────────┘                        │
│                            │                      │                               │
│                            ▼                      ▼                               │
│                    ┌───────────────┐      ┌────────────────┐                      │
│                    │  Integration  │      │ Infrastructure │                      │
│                    │  (服务集成)   │ ───▶ │  (基础设施)    │                      │                   
│                    └───────────────┘      └────────────────┘                      │
│                            │                       │                              │
│                            ▼                       ▼                              │
│                      ┌─────────────────────────────────────┐                      │
│                      │  PostgreSQL / Redis / External API  │                      │
│                      └─────────────────────────────────────┘                      │
│                                                                                   │
└───────────────────────────────────────────────────────────────────────────────────┘
```

#### 请求流程图

```
HTTP Request
    │
    ▼
┌─────────────────┐
│   Middleware    │  ← 认证、日志、限流、CORS 等
│  (http/router)  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Router      │  ← 路由分发 (Manager/Rider)
│  (http/router)  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Handler     │  ← HTTP 处理器
│ (http/*/handler)│  接收请求、参数绑定
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Service     │  ← 业务逻辑
│ (presentation/  │  调用 DAO、Integration
│    service)     │
└────────┬────────┘
         │
         ├─────────────┬──────────────┐
         ▼             ▼              ▼
┌───────────────┐ ┌───────────┐  ┌─────────────┐
│      DAO      │ │Integration│  │    DTO      │
│  (数据访问)   │ │   (集成)  │  │  (数据传输) │
└───────┬───────┘ └─────┬─────┘  └─────────────┘
        │               │
        ▼               ▼
┌─────────────┐ ┌──────────────┐
│  Ent ORM    │ │    Redis     │
│ (database)  │ │  External API│
└─────────────┘ └──────────────┘
```

### 目录结构

```
.
├── assets                              # 资源文件
│   ├── assets.go
│   ├── docs                            # API 文档
│   ├── presets                         # 预设数据
│   └── templates                       # 模板文件
├── pkg                                 # 公共工具包
├── wiki                                # 文档目录
└── internal                            # 内部代码
    ├── bootstrap                       # 启动引导
    │   ├── boot.go                     # 启动入口
    ├── config                          # 配置管理
    │   ├── config.go                   # 配置结构
    ├── di                              # 依赖注入 (DI Container)
    │   ├── di.go                       # 容器定义
    │   ├── wire.go                     # Wire 配置
    ├── infrastructure                  # 基础设施层
    │   ├── dao                         # 数据访问对象
    │   │   └── pagination              # 分页工具
    │   ├── ent                         # Ent ORM
    │   └── model                       # 模型定义
            └── ...                     # 业务模型、错误定义等
    ├── integration                     # 集成服务层
    │   ├── alisms                      # 阿里云短信
    │   ├── auth                        # 认证服务
    │   ├── captcha                     # 验证码服务
    │   ├── hotfilestore                # 热文件存储
    │   ├── wxapp                       # 微信小程序
    │   └── kafka                       # Kafka 客户端封装
    ├── presentation                    # 展现层
    │   ├── dto                         # 数据传输对象
    │   └── service                     # 业务服务
    └── application                     # 应用服务层
        ├── application.go              # 服务启动、停止
        ├── http                        # HTTP 服务
        │   ├── http.go                 # HTTP 初始化
        │   ├── core                    # 核心组件
        │   │   └── ...                 # Context、Middleware 等
        │   ├── handler                 # http处理器
        │   └── router                  # 路由配置
        │       ├── router.go           # 路由入口
        │       └── docs.go             # 文档路由
        └── kafka                       # Kafka 监听
```
