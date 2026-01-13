#!/usr/bin/env bash

.PHONY: doc, doc-clean, doc-preview, build, build-image

doc:
	@echo "正在生成文档..."
	go run -mod=mod github.com/liasica/godoc/cmd/godoc generate

doc-clean:
	@echo "清理文档..."
	rm -rf docs

doc-preview:
	@echo "正在生成并预览文档..."
	go run -mod=mod github.com/liasica/godoc/cmd/godoc preview --generate

build:
	make doc
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -tags=sonic,poll_opt -gcflags "all=-N -l" -ldflags "-X nexis.run/nexa-layout/internal/config.Version=${git rev-parse --short HEAD}" -o build/release/layout cmd/layout/main.go

build-image:
	make build
	@echo "正在构建并推送Docker镜像..."
	docker build --platform=linux/amd64 -t harbor.liasica.com/auroraride/layout:prod -f Dockerfile .
	docker push harbor.liasica.com/auroraride/layout:prod
