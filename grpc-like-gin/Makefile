GO=go
GOPATH:=$(shell go env GOPATH)

ifeq ($(strip $(goos)),)
	GOOS=linux
else
	GOOS=$(strip $(goos))
endif

ifeq ($(strip $(platforms)),)
	PLATFORMS=amd64
else
	PLATFORMS=$(strip $(platforms))
endif

BUILD := `git rev-parse --short HEAD`
BUILD_IMAGE := golang:1.16.1
BUILD_OUTPUT_DIR := build
PROJECT_NAME := grpc-like-gin

LDFLAGS += -X "$(project)/version.BuildTS=$(shell TZ='Asia/Shanghai' date '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "$(project)/version.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "$(project)/version.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "$(project)/version.App=$(PROJECT_NAME)"

project=github.com/xiaomLee/go-example/grpc-like-gin
SRC = $(shell find . -type f -name '*.go')

.PHONY: clean
clean:
	rm -rf ${BUILD_OUTPUT_DIR}

.PHONY: build
code_root = /go/$(project)
DOCKER_MOUNT := ${PWD}:${code_root}
ifdef mount_path
	DOCKER_MOUNT := ${DOCKER_MOUNT} -v ${mount_path}
endif

define docker_env
	docker run --rm \
		--env TZ=Asia/Shanghai \
		--env goos=$(goos)\
		--env platforms="$(platforms)" \
		--env GOPROXY=https://goproxy.cn,direct \
		-v ${DOCKER_MOUNT} \
		-w ${code_root} \
		${BUILD_IMAGE} sh -c "$(1)"
endef

$(PLATFORMS):$(SRC) output
	mkdir -p ${BUILD_OUTPUT_DIR}/$@
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$@ $(GO) build -ldflags '$(LDFLAGS)' -a -installsuffix cgo -o ${BUILD_OUTPUT_DIR}/$@ $(project)

output:clean
	mkdir -p ${BUILD_OUTPUT_DIR}

build: $(PLATFORMS)

docker-build:
	$(call docker_env,make build)

.PHONY: test

.PHONY: check
lint:
	@golangci-lint run ./... --deadline=5m

.PHONY: docker
image: clean $(TARGETS)
	docker build --build-arg BIN='$(TARGETS)' -t $(PROJECT_NAME):$(VERSION)-$(GOARCH)-$(BUILD) .

.PHONY: proto
proto:
	protoc -I ./pb --go_out=paths=source_relative:./api --go-grpc_out=paths=source_relative:./api grpc_end.proto
