.PHONY: clean check-godoc doc doc-preview build build-image

clean:
	@echo "正在清理构建文件..."
	rm -rf build/
	@echo "正在清理文档..."
	rm -rf docs/

check-godoc:
	@if ! command -v godoc &> /dev/null; then \
		echo "godoc 未安装，正在安装..."; \
		curl -fsSL https://raw.githubusercontent.com/liasica/godoc/master/install.sh | bash; \
	fi

doc: check-godoc
	@echo "正在生成文档..."
	godoc generate

doc-preview:
	@echo "正在生成并预览文档..."
	godoc preview --generate

build: doc
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -tags=sonic,poll_opt -gcflags "all=-N -l" -ldflags "-X nexis.run/nexa-layout/internal/config.Version=$(shell git rev-parse --short HEAD)" -o build/release/layout cmd/layout/main.go
	upx -9 ./build/release/layout

build-image: build
	@echo "正在构建并推送Docker镜像..."
	docker build --platform=linux/amd64 -t harbor.liasica.com/auroraride/layout:prod -f Dockerfile .
	docker push harbor.liasica.com/auroraride/layout:prod
