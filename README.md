# kratos-layout

基于DDD领域驱动设计的kratos模板

- https://juejin.cn/post/7298160530292703244
- https://github.com/KendoCross/kendoDDD
- https://github.com/vgocoder/go-ddd
- https://juejin.cn/post/7226556923238203429

### 文档工具

[Stoplight Elements](https://stoplight.io/open-source/elements)

### 基本结构

```
                ┌────────────────────────┐
                │     Infrastructure     │
              / └────────────────────────┘ \
             /               |              \
            /                |               \
           ▼                 ▼                ▼
┌───────────────┐      ┌──────────────┐      ┌─────────────────────┐
│     Domain    │ ────►│ Presentation │ ────►│ Server / Rest / ... │
└───────────────┘      └──────────────┘      └─────────────────────┘
```

### 目录结构

```
.
├── README.md
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── go.sum
│── internal
│   ├── bootstrap                           # 启动目录
│   │   └── boot.go                         # 启动文件
│   ├── domain                              # 领域层，主要负责抽象化业务逻辑，提供领域服务和实体
│   │   ├── entity                          # 实体目录，用于定义领域对象以及请求、响应数据结构
│   │   ├── repository                      # 仓储目录，和数据库进行数据交换，例如：CRUD等
│   │   └── service                         # 领域服务目录，用于抽象业务逻辑
│   ├── infrastructure                      # 基础设施层，包含数据库、缓存等
│   │   ├── common                          # 基础包
│   │   ├── model                           # 模型目录，包含数据库模型、缓存模型等
│   │   └── vo                              # 定义值对象，包含常量、属性、错误等
│   ├── presentation                        # 展现层，主要负责接收请求和返回响应
│   │   ├── repository                      # 仓储目录，实现domain层的repository接口
│   │   └── service                         # 展现服务目录，用户实现domain层的service接口或实现grpc服务，配合repository处理业务逻辑
│   ├── rest                                # Rest接口层
│   │   ├── app                             # 定义rest服务（仅rest使用），例如context、middleware、validator等
│   │   ├── controller                      # 控制器，用于处理接收数据和返回数据，调用service对外提供服务
│   │   └── route                           # 路由
│   └── server                              # 服务目录，包含grpc、http、websocket等服务
│       ├── micro                           # 微服务目录
│       └── rest                            # Rest服务目录
│           ├── app                         # 定义服务（仅 rest 使用），例如context、middleware、validator等
│           │   ├── context.go              # ...
│           │   ├── header.go               # ...
│           │   └── user_middleware.go      # ...
│           ├── controller                  # 控制器目录，用于处理接收数据和返回数据，调用service对外提供服务
│           │   ├── controller.go           # ...
│           │   └── user.go                 # ...
│           ├── route                       # 路由目录
│           │   └── route.go                # 路由入口文件
│           └── rest.go                     # rest服务入口文件
└── pb                              # pb目录，用于存放proto文件，大仓模式全局共享
```
