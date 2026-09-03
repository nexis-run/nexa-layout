# nexa-layout

基于 DDD 领域驱动设计、`google/wire` 依赖注入与 `nexis.run/nexa` 框架的 Go 服务端项目模板，开箱提供 HTTP（Echo）+ gRPC（Kratos）双协议入口、统一日志、配置加载、优雅退出与接口文档脚手架。

参考资料：

- https://juejin.cn/post/7298160530292703244
- https://github.com/KendoCross/kendoDDD
- https://github.com/vgocoder/go-ddd
- https://juejin.cn/post/7226556923238203429

## 一、从 Clone 到运行

### 1.1 前置依赖

| 工具 | 版本 | 说明 |
| --- | --- | --- |
| Go | ≥ 1.25.3 | `go.mod` 声明的最低版本 |
| Docker | 任意 | 仅 `make build-image` 需要 |
| gclint | 最新 | golangci-lint 增强版，提交前 lint 用，安装见 `nexis.run/nexa` 文档 |

### 1.2 克隆与重命名

模板里所有 import path、命令名、镜像 tag 都包含 `nexa-layout` / `layout` / `layout-demo` 字样，**新项目第一步先做整体替换**。假设新项目名为 `myapp`，模块路径为 `nexis.run/myapp`：

```bash
# 1. 克隆
git clone https://github.com/nexis-run/nexa-layout.git myapp
cd myapp
rm -rf .git && git init

# 2. 重写 go.mod 模块名
go mod edit -module nexis.run/myapp

# 3. 全量替换 import 路径（macOS 用 gsed，Linux 用 sed -i）
grep -rl 'nexis.run/nexa-layout' --include='*.go' --include='*.yaml' --include='Makefile' --include='Dockerfile' . \
  | xargs sed -i '' 's|nexis.run/nexa-layout|nexis.run/myapp|g'

# 4. 重命名 cmd 目录与二进制名
git mv cmd/layout cmd/myapp
sed -i '' 's|cmd/layout|cmd/myapp|g; s|build/release/layout|build/release/myapp|g' Makefile Dockerfile

# 5. 改配置（应用名、Kafka topic、镜像 tag、Cobra Use 字段）
sed -i '' 's/layout-demo/myapp/g' config/config.yaml
sed -i '' 's|"layout"|"myapp"|g; s|nexis.run layout|nexis.run myapp|g' cmd/myapp/main.go
sed -i '' 's|harbor\.liasica\.com/auroraride/layout|harbor.example.com/yourorg/myapp|g' Makefile

# 6. 重置 wire 注释里的 Copyright（可选）
grep -rl 'nexa-layout. 2026-present' --include='*.go' . \
  | xargs sed -i '' 's|nexa-layout. 2026-present|myapp. 2026-present|g'

# 7. 拉依赖、重新生成 wire、本地编译验证
go mod tidy
make wire
go build ./...
```

### 1.3 替换清单（如果不用脚本，手工逐项检查）

| 位置 | 原值 | 改为 |
| --- | --- | --- |
| `go.mod:1` | `module nexis.run/nexa-layout` | 你的模块路径 |
| `cmd/layout/` | 目录名 `layout` | 你的可执行名 |
| `cmd/layout/main.go` | `Use: "layout"`、`Short: "nexis.run layout"` | 你的命令名 |
| `config/config.yaml` | `app: layout-demo`、Kafka topic、Brokers | 你的应用配置 |
| `Makefile` | `build/release/layout`、Harbor 镜像地址 | 你的产物路径与镜像 |
| `Dockerfile` | `COPY ./build/release/layout`、`ENTRYPOINT` | 你的二进制名 |
| `pb/layout/v1/*.proto` | `package layout.v1`、`go_package` | 你的 proto 包名 |
| `buf.gen.yaml` | `go_package_prefix` | 你的 pb gen 路径 |
| `.godoc.yaml` | `path:` 下的扫描目录 | 视新结构调整 |
| 各文件头 Copyright | `oos / micro-layout / nexa-layout` | 你的项目名 |

### 1.4 启动

```bash
# 本地直接 run（不打包）
go run ./cmd/myapp app --config config/config.yaml

# 编译后启动
make build-local
./build/release/myapp app --config config/config.yaml

# 跨平台构建（linux/amd64，给 Docker 用）
make build

# 构建并推送镜像
make build-image
```

默认端口：HTTP `:50050`（OpenAPI 文档 `http://localhost:50050/docs`），gRPC `:50051`。

## 二、常用 Make 目标

| 目标 | 说明 |
| --- | --- |
| `make wire` | 重新生成 `internal/di/wire_gen.go`，修改 di / provider 后必须执行 |
| `make ent` | 生成 Ent 代码（添加 schema 后执行） |
| `make doc` | 生成 OpenAPI/Swagger 静态文档到 `assets/docs/` |
| `make doc-preview` | 本地预览接口文档 |
| `make build-local` | 当前平台构建 |
| `make build` | linux/amd64 构建（CI / 镜像用） |
| `make build-image` | 构建并推送 Docker 镜像 |
| `make clean` | 清理构建产物与生成的文档 |
| `make all` | `clean + build-image` 一键发布 |

## 三、提交前自查

项目根目录有 `.golangci.yml`，启用了 `gclint`、`wsl_v5`、`revive` 等检查。**提交前必跑**：

```bash
/usr/local/bin/gclint run --config .golangci.yml --new-from-rev=HEAD~1 --timeout=10m
```

Commit message 遵循 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/)，常用前缀：

```
feat:         新增功能
fix:          bug 修复
improvement:  原有功能优化
refactor:     代码重构（不涉及功能变动）
performance:  性能优化
chore:        构建工具 / 配置修改
deps:         第三方依赖修改
docs:         文档变动
style:        代码风格调整
test:         单元测试
revert:       回滚
```

## 四、添加新业务的开发流程

以新增 `Order` 业务为例：

```text
1. 模型      internal/infrastructure/model/order.go        定义领域结构体
2. DAO       internal/infrastructure/dao/order.go          定义 OrderDao + NewOrder()
                                                           （可选）使用脚手架：
                                                           go run -mod=mod nexis.run/nexa/cmd/nexa@master new dao Order
3. 接入 DI   internal/di/di.go                             把 dao.NewOrder 加进 daoProviderSet
                                                           Dao 结构体增加 Order *dao.OrderDao 字段
4. Service   internal/presentation/service/order.go        OrderService 构造函数接收 *dao.OrderDao
5. 接入 DI   internal/di/di.go                             把 service.NewOrder 加进 serviceProviderSet
                                                           Service 结构体增加 Order *service.OrderService
6. DTO       internal/presentation/dto/order.go            请求 / 响应结构体（json + validate tag）
7. Handler   internal/application/http/handler/order/order.go
                                                           包级函数：func Create(c echo.Context) error
8. 路由      internal/application/http/router/router.go    注册 orderhandler.Create 等
9. 生成 wire make wire
10. lint     gclint ... --new-from-rev=HEAD~1
11. 提交     git commit -m "feat: 添加 Order 业务"
```

> 重点：所有业务依赖都通过 `di.C.Service.<X>` / `di.C.Dao.<X>` 取，**禁止在 handler / middleware 内部用 `New...()` 临时构造**，否则 Wire 容器会失去意义。

## 五、整体架构

### 5.1 架构图

```
┌──────────────────────────────────────────────────────────────────────────────┐
│                                  项目架构图                                  │
├──────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌──────────────┐   │
│  │  HTTP/gRPC  │ ─▶ │   Handler   │ ─▶ │   Service   │ ◀─ │ Application  │   │
│  │  (外部请求) │    │   (接口层)  │    │   (业务层)  │    │  (其他服务)  │   │
│  └─────────────┘    └─────────────┘    └─────────────┘    └──────────────┘   │
│                            │                  │                              │
│                            ▼                  ▼                              │
│                     ┌────────────────────────────────┐                       │
│                     │      DI Container (di.C)       │                       │
│                     │   Wire 生成, bootstrap 注入    │                       │
│                     └────────────────────────────────┘                       │
│                            │                  │                              │
│                            ▼                  ▼                              │
│                   ┌───────────────┐   ┌────────────────┐                     │
│                   │  Integration  │   │ Infrastructure │                     │
│                   │  (服务集成)   │ ─▶│  (基础设施)    │                     │
│                   └───────────────┘   └────────────────┘                     │
│                            │                  │                              │
│                            ▼                  ▼                              │
│                     ┌────────────────────────────────────┐                   │
│                     │ PostgreSQL / Redis / External API  │                   │
│                     └────────────────────────────────────┘                   │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘
```

### 5.2 请求流程

```
HTTP Request
    │
    ▼
┌────────────────────┐
│    Middleware      │  ← CORS、Context、Dump、Auth
│  (http/router)     │     全局：CORS / Context / Dump
└────────┬───────────┘     路由组：Auth
         ▼
┌────────────────────┐
│      Router        │  ← e.POST / e.Group(...).GET
│  (http/router)     │
└────────┬───────────┘
         ▼
┌────────────────────┐
│      Handler       │  ← 参数绑定、调用 Service
│  (http/handler/X)  │     从 di.C.Service.X 取依赖
└────────┬───────────┘
         ▼
┌────────────────────┐
│      Service       │  ← 业务逻辑
│ (presentation/     │     从构造函数接收 DAO / Integration
│  service)          │
└────────┬───────────┘
         ├──────────────┬──────────────┐
         ▼              ▼              ▼
┌──────────────┐ ┌────────────┐ ┌───────────────┐
│      DAO     │ │Integration │ │      DTO      │
│  (数据访问)  │ │   (集成)   │ │   (数据传输)  │
└──────┬───────┘ └─────┬──────┘ └───────────────┘
       ▼               ▼
┌──────────────┐ ┌──────────────┐
│   Ent ORM    │ │    Redis     │
│  (database)  │ │ External API │
└──────────────┘ └──────────────┘
```

## 六、目录结构

```
.
├── assets/                              # 静态资源（embed）
│   ├── assets.go                        # TemplateFS / OpenAPIData
│   ├── docs/                            # OpenAPI / Swagger 输出
│   └── templates/                       # HTML 模板
├── cmd/
│   └── layout/                          # 可执行入口（按业务改名）
│       ├── main.go                      # Cobra 根命令
│       └── internal/app.go              # app 子命令：启动服务端
├── config/
│   └── config.yaml                      # 默认配置
├── pb/                                  # protobuf 源 + 生成产物
│   ├── layout/v1/*.proto
│   └── gen/                             # buf 生成的 *.pb.go
├── internal/
│   ├── bootstrap/boot.go                # 启动引导：加载 config → init logger → wire DI
│   ├── config/config.go                 # 配置 Load() / Config 结构
│   ├── di/                              # 依赖注入容器
│   │   ├── di.go                        # Container / Dao / Service / Integration 聚合
│   │   ├── wire.go                      # Wire build tag 入口
│   │   └── wire_gen.go                  # 生成产物
│   ├── infrastructure/                  # 基础设施层
│   │   ├── dao/
│   │   │   ├── user.go
│   │   │   └── pagination/              # 通用分页工具
│   │   ├── ent/                         # Ent ORM（添加 schema 后生成）
│   │   └── model/                       # 领域模型与领域错误
│   ├── integration/                     # 第三方集成（按需新增子包）
│   ├── presentation/                    # 展现层
│   │   ├── dto/                         # 请求 / 响应结构体
│   │   └── service/                     # 业务逻辑实现
│   └── application/                     # 应用服务层
│       ├── server.go                    # graceful Start / Stop（并发关闭）
│       ├── http/
│       │   ├── http.go                  # Echo Setup
│       │   ├── core/                    # Context、Header、AuthMiddleware
│       │   ├── handler/                 # 每业务一个子包
│       │   │   └── user/user.go
│       │   └── router/
│       │       ├── router.go            # 路由 + 全局中间件
│       │       └── docs.go              # OpenAPI / Redoc 路由
│       └── micro/                       # gRPC 服务端
│           ├── micro.go                 # Kratos gRPC Setup
│           └── user/                    # 每业务一个子包
│               ├── transport.go         # pb ↔ dto 适配 + 调 di.C.Service
│               └── user.go              # RegisterUserServiceServer
```

## 七、文档工具

接口文档使用 [Stoplight Elements](https://stoplight.io/open-source/elements) 渲染，访问 `http://localhost:50050/docs`。OpenAPI Spec 由 [godoc](https://github.com/liasica/godoc) 从代码注释生成：

- 配置：`.godoc.yaml`
- 生成：`make doc`
- 预览：`make doc-preview`

## 八、常用脚手架命令

```bash
# 生成 DAO
go run -mod=mod nexis.run/nexa/cmd/nexa@master new dao Name

# 生成 Echo Context
go run -mod=mod nexis.run/nexa/cmd/nexa@master new echoctx Name

# 添加 Ent schema
go run -mod=mod nexis.run/nexa/cmd/nexa@master ent new Name

# 生成 Ent 代码
go run -mod=mod nexis.run/nexa/cmd/nexa@master ent generate
```
