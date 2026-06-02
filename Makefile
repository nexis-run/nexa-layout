.PHONY: clean ent wire doc doc-preview build build-local build-image

clean:
	@echo "正在清理构建文件..."
	rm -rf build/
	@echo "正在清理文档..."
	rm -rf assets/docs/
	@echo "✅ 清理完成"

ent:
	@echo "正在生成ent代码..."
	$(eval NEXA_DIR := $(shell go list -m -f '{{.Dir}}' nexis.run/nexa))
	@# //go:generate sh -c "go run -mod=mod entgo.io/ent/cmd/ent generate --idtype uint64 --feature sql/modifier,sql/upsert,privacy,entql,sql/execquery,intercept,schema/snapshot --template $(go list -m -f '{{.Dir}}' nexis.run/nexa)/kit/entx/template/meta.tmpl,$(go list -m -f '{{.Dir}}' nexis.run/nexa)/kit/entx/template/upsert.tmpl,$(go list -m -f '{{.Dir}}' nexis.run/nexa)/kit/entx/template/soft_delete.tmpl ./schema"
	@# go run -mod=mod entgo.io/ent/cmd/ent generate --idtype uint64 --feature sql/modifier,sql/upsert,privacy,entql,sql/execquery,intercept,schema/snapshot --template $(NEXA_DIR)/cmd/nexa/internal/entgen/template/meta.tmpl,$(NEXA_DIR)/cmd/nexa/internal/entgen/template/upsert.tmpl,$(NEXA_DIR)/cmd/nexa/internal/entgen/template/soft_delete.tmpl ./internal/infrastructure/ent/schema
	go run -mod=mod nexis.run/nexa/cmd/nexa@master ent generate
	@echo "✅ ent代码生成完成"

wire:
	@echo "正在生成 Wire 代码..."
	go run -mod=mod github.com/google/wire/cmd/wire ./internal/di
	@echo "✅ Wire 代码生成完成"

doc:
	@echo "正在生成文档..."
	go run -mod=mod github.com/liasica/godoc/cmd/godoc@v1.0.10 generate
	@echo "✅ 文档生成完成，文档位于 assets/docs/ 目录下"

doc-preview:
	@echo "正在生成并预览文档..."
	go run -mod=mod github.com/liasica/godoc/cmd/godoc@v1.0.10 preview --generate

build-local: doc wire
	go build -trimpath -tags=sonic,poll_opt -gcflags "all=-N -l" -ldflags "-X nexis.run/nexa-layout/internal/config.Version=$(shell git rev-parse --short HEAD)" -o build/release/layout cmd/layout/main.go
	@echo "✅ 本地构建完成，二进制文件位于 build/release/layout"

build: doc wire
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -tags=sonic,poll_opt -gcflags "all=-N -l" -ldflags "-X nexis.run/nexa-layout/internal/config.Version=$(shell git rev-parse --short HEAD)" -o build/release/layout cmd/layout/main.go
	@echo "✅ 构建完成，二进制文件位于 build/release/layout"

build-image: build
	@echo "正在构建并推送Docker镜像..."
	docker build --platform=linux/amd64 -t harbor.liasica.com/auroraride/layout:prod -f Dockerfile .
	docker push harbor.liasica.com/auroraride/layout:prod
	@echo "✅ Docker镜像构建并推送完成"

all: clean build-image
	@echo "✅ 所有任务完成"
