# nexa-layout

基于 Nexa 的 Go 服务模板，提供 HTTP、gRPC、配置、日志、Wire 注入和 OpenAPI 文档。

## 构建与运行

需要 Go 1.25.3 或更高版本。依赖访问配置见 [Nexa 文档](https://github.com/nexis-run/nexa)。

```bash
git clone https://github.com/nexis-run/nexa-layout.git
cd nexa-layout
go mod download
make build-local
./build/release/layout app --config config/config.yaml
```

HTTP 默认监听 `:50050`，gRPC 默认监听 `:50051`。接口文档位于 `/docs`，规范文件位于 `/docs/openapi.yaml`。默认日志输出到控制台。

`wire_gen.go`、protobuf 和 OpenAPI 产物纳入版本控制。普通构建直接使用这些文件；修改对应源文件后执行生成命令。

## 常用命令

| 命令 | 行为 |
| --- | --- |
| `make`、`make build-local` | 构建当前平台的服务 |
| `make build` | 构建 Linux 静态二进制，默认架构为 amd64 |
| `make wire` | 生成依赖注入代码 |
| `make doc` | 生成 OpenAPI 规范 |
| `make generate` | 生成 Wire 与 OpenAPI 代码 |
| `make ent` | 使用当前模块依赖中的 Nexa CLI 生成 Ent 代码 |
| `make doc-preview` | 生成并预览接口文档 |
| `make build-image` | 构建容器镜像 |
| `make push-image` | 构建并推送镜像 |
| `make clean` | 删除构建目录 |

构建变量：

```bash
make build GOARCH=arm64 VERSION=1.0.0
make build-image IMAGE=registry.example.com/team/layout:1.0.0 VERSION=1.0.0
make push-image IMAGE=registry.example.com/team/layout:1.0.0 VERSION=1.0.0
```

`APP` 同时指定命令目录与二进制名称，默认值为 `layout`。`VERSION` 默认取当前提交的短哈希；模块路径从 `go.mod` 读取。镜像内以非 root 用户运行，配置路径为 `/app/config/config.yaml`，可以通过挂载替换。

## 项目结构

| 目录 | 职责 |
| --- | --- |
| `cmd/layout` | CLI 入口、配置参数与退出状态 |
| `internal/bootstrap` | 配置加载、日志初始化与依赖组装 |
| `internal/application` | HTTP、gRPC 服务生命周期 |
| `internal/application/http` | HTTP 适配、上下文、中间件、处理器和路由 |
| `internal/application/micro` | gRPC 服务注册与协议转换 |
| `internal/presentation/service` | 业务逻辑 |
| `internal/presentation/dto` | 请求与响应类型 |
| `internal/infrastructure/dao` | 数据访问与分页工具 |
| `internal/infrastructure/model` | 业务模型与错误 |
| `internal/di` | Wire provider 与依赖容器 |
| `assets` | HTML 模板和 OpenAPI 规范 |
| `pb` | protobuf 定义与生成代码 |

HTTP 与 gRPC 共享 Service。Service 通过构造函数接收 DAO，处理器通过 `di.C.Service` 获取业务依赖。Kratos 应用负责启动两种协议、传播启动错误，并在退出信号后并发关闭服务；服务关闭期限为 30 秒，日志关闭期限为 5 秒。

## 配置

默认配置位于 `config/config.yaml`。`baseurl` 可指定包含反向代理前缀的公开基础地址，例如 `https://api.example.com/service`；未指定时，文档按当前请求地址渲染。

```yaml
app: myapp
environment: development
baseurl: 'https://api.example.com/service'
logger:
  stdout: true
http:
  bind: ':50050'
grpc:
  bind: ':50051'
```

Kafka 日志通过 `logger.kafka.topic` 和 `logger.kafka.brokers` 配置。业务配置定义在 `internal/config/config.go`。

## 添加业务与 Ent

`.nexa.yaml` 定义 Ent、DAO、Context 与 DI 的路径。以 `Order` 为例：

```bash
go run nexis.run/nexa/cmd/nexa ent new Order
go mod tidy
make ent
go run nexis.run/nexa/cmd/nexa new dao Order --di=false
go mod tidy
```

DAO 构造器接收 `*ent.Client`。接入数据库时，在基础设施层创建客户端，将它加入 Wire provider；把 DAO 构造器加入 `daoProviderSet`，将实例加入 `Dao` 结构体，再执行 `make wire`。客户端的连接参数、迁移与关闭逻辑由应用定义。

已配置客户端 provider 的项目可以省略 `--di=false`，由 CLI 维护 DAO 聚合字段。生成器与项目应使用相同的 Ent 版本，参见 [Ent 代码生成文档](https://entgo.io/docs/code-gen/)。

Service、DTO、HTTP Handler 和 gRPC 适配器按业务实现，路由在 `internal/application/http/router` 注册。Echo Context 可以通过以下命令生成：

```bash
go run nexis.run/nexa/cmd/nexa new echoctx Order
```

示例用户服务使用固定演示数据和令牌，业务认证实现在 `internal/presentation/service/user.go`。

## 作为新项目使用

设置新的模块路径后，同步调整 Go import、protobuf 的 `go_package` 与 `buf.gen.yaml` 中的前缀。命令目录按应用名重命名，通过 `APP` 指定对应构建入口。更新 `config/config.yaml` 的应用配置，然后执行 `go mod tidy` 和 `make generate`。

## 检查

```bash
go test -race ./...
/usr/local/bin/gclint run --config .golangci.yml --new-from-rev=HEAD --timeout=10m
```

框架与模板联调时，在两个仓库的上级目录创建本地 Go workspace：

```bash
go work init ./nexa ./nexa-layout
go test ./nexa/... ./nexa-layout/...
```

Go workspace 文件仅用于本地联调。单独检查模块发布状态时设置 `GOWORK=off`。
