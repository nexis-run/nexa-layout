.DEFAULT_GOAL := build-local
.PHONY: all clean ent wire doc doc-preview generate build build-local build-image push-image

APP ?= layout
MODULE := $(shell GOWORK=off go list -m)
VERSION ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo dev)
IMAGE ?= $(APP):$(VERSION)
GOOS ?= linux
GOARCH ?= amd64
LDFLAGS := -s -w -X $(MODULE)/internal/config.Version=$(VERSION)

clean:
	rm -rf build/

ent:
	go run nexis.run/nexa/cmd/nexa ent generate

wire:
	go run -mod=mod github.com/google/wire/cmd/wire ./internal/di

doc:
	go run -mod=mod github.com/liasica/godoc/cmd/godoc@v1.0.10 generate

doc-preview:
	go run -mod=mod github.com/liasica/godoc/cmd/godoc@v1.0.10 preview --generate

generate: wire doc

build-local:
	go build -trimpath -ldflags '$(LDFLAGS)' -o build/release/$(APP) ./cmd/$(APP)

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -trimpath -ldflags '$(LDFLAGS)' -o build/release/$(APP) ./cmd/$(APP)

build-image: build
	docker build --platform=$(GOOS)/$(GOARCH) --build-arg APP=$(APP) -t $(IMAGE) .

push-image: build-image
	docker push $(IMAGE)

all: build-local
