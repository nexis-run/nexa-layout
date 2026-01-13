# kratos-layout

基于DDD领域驱动设计的kratos模板

- https://juejin.cn/post/7298160530292703244
- https://github.com/KendoCross/kendoDDD
- https://github.com/vgocoder/go-ddd
- https://juejin.cn/post/7226556923238203429

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

```
         ┌────────────────────────┐
         │     Infrastructure     │
       / └────────────────────────┘ \
      /                              \
     /                                \
    ▼                                  ▼
┌──────────────┐      ┌─────────────────────┐
│ Presentation │ ────►│ Server / Rest / ... │
└──────────────┘      └─────────────────────┘
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
│   ├── infrastructure                      # 基础设施层，包含数据库、缓存等
│   │   ├── common                          # 基础包
│   │   ├── model                           # 模型目录，包含数据库模型、缓存模型等
│   │   └── vo                              # 定义值对象，包含常量、属性、错误等
│   ├── presentation                        # 展现层，主要负责接收请求和返回响应
│   │   ├── entity                          # 实体目录，用于定义领域对象以及请求、响应数据结构
│   │   ├── repository                      # 仓储目录，实现domain层的repository接口
│   │   └── service                         # 展现服务目录，用户实现domain层的service接口或实现grpc服务，配合repository处理业务逻辑
│   ├── rest                                # Rest接口层
│   │   ├── app                             # 定义rest服务（仅rest使用），例如context、middleware、validator等
│   │   ├── controller                      # 控制器，用于处理接收数据和返回数据，调用service对外提供服务
│   │   └── route                           # 路由
│   └── app                                 # 应用目录，包含grpc、http、websocket等服务
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
