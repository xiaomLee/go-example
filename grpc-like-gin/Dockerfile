FROM alpine

ARG ARCH=amd64
ARG PROJECT_ROOT=/app
ARG BIN=grpc-like-gin

WORKDIR ${PROJECT_ROOT}
COPY build/${ARCH} ${PROJECT_ROOT}

CMD ["/bin/sh", "-c", "${PROJECT_ROOT}/${BIN}"]