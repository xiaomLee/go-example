GO=go
GOPATH:=$(shell go env GOPATH)
OS:=$(shell go env GOOS)
ARCH:=$(shell go env GOHOSTARCH)
GCC:=${PWD}/aarch64-linux-musl-cross/bin/aarch64-linux-musl-gcc

ifneq ($(strip $(os)),)
	OS=$(strip $(os))
endif

ifneq ($(strip $(arch)),)
	ARCH=$(strip $(arch))
endif

BUILD_IMAGE := golang:1.16-alpine
BUILD_OUTPUT_DIR := build
PROJECT_NAME := grpc-ecosystem-template

LDFLAGS += -X "$(project)/version.BuildTS=$(shell TZ='Asia/Shanghai' date '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(project)/version.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(project)/version.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "$(project)/version.App=$(PROJECT_NAME)"

# static-link build
LDFLAGS += '-extldflags=-static'

project=grpc-ecosystem-template
SRC = $(shell find . -type f -name '*.go')

.PHONY: clean
clean:
	rm -rf ${BUILD_OUTPUT_DIR}/$(ARCH)

.PHONY: build
code_root = /go/$(project)
DOCKER_MOUNT := ${PWD}:${code_root}
ifdef mount_path
	DOCKER_MOUNT := ${DOCKER_MOUNT} -v ${mount_path}
endif

define docker_env
	docker run --rm \
		--env TZ=Asia/Shanghai \
		--env os=$(OS)\
		--env arch=$(ARCH) \
		--env GOPROXY=https://goproxy.cn,direct \
		-v ${DOCKER_MOUNT} \
		-w ${code_root} \
		${BUILD_IMAGE} sh -c "$(1)"
endef

build: $(SRC) output
	mkdir -p ${BUILD_OUTPUT_DIR}/$(ARCH)
	# sqlite need set cgo_enable=1
	CGO_ENABLED=1 CC=$(GCC) GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -ldflags '$(LDFLAGS)' -a -installsuffix cgo -o ${BUILD_OUTPUT_DIR}/$(ARCH)/$(PROJECT_NAME) $(project)

output:clean
	mkdir -p ${BUILD_OUTPUT_DIR}

docker-build:
	$(call docker_env,make build)

.PHONY: test

.PHONY: check
lint:
	@golangci-lint run ./... --deadline=5m

.PHONY: image
IMAGE_TAG := $(shell git rev-parse --abbrev-ref HEAD | sed 's/\//-/g')-$(shell git rev-parse --short HEAD)

image: clean $(TARGETS)
	docker build --build-arg BIN='$(TARGETS)' -t $(PROJECT_NAME):$(ARCH)-$(IMAGE_TAG) .

.PHONY: proto
proto:
	protoc -I ./pb \
		--proto_path=./pb:./third-party \
		--go_out=paths=source_relative:./api \
		--go-grpc_out=paths=source_relative:./api \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:./api \
        --swagger_out=logtostderr=true:./swagger/ \
		user-service.proto user.proto
