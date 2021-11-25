FROM alpine

ARG ARCH=amd64
ARG PROJECT_ROOT=/app
ARG BIN=grpc-ecosystem-template

COPY build/${ARCH}/${BIN} ${PROJECT_ROOT}/
WORKDIR ${PROJECT_ROOT}

CMD ["./grpc-ecosystem-template"]