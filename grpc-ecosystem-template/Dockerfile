# build stage
FROM golang:1.16-alpine as builder
#ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache git make build-base

WORKDIR /app
COPY . .
RUN make build; du -h ./build

# package stage
FROM alpine
ARG ARCH=amd64
ARG PROJECT_ROOT=/app
ARG BIN=grpc-ecosystem-template

WORKDIR ${PROJECT_ROOT}
COPY --from=builder /app/build/${ARCH}/${BIN} .
COPY ./configs ./configs
COPY ./testdata ./testdata
COPY ./entrypoint.sh ${PROJECT_ROOT}/

ENTRYPOINT ["./entrypoint.sh"]
CMD ["./grpc-ecosystem-template"]