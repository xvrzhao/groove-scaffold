FROM golang:1.23-alpine3.20@sha256:9dd2625a1ff2859b8d8b01d8f7822c0f528942fe56cfe7a1e7c38d3b8d72d679 AS builder
RUN go env -w GO111MODULE=on GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -o ./target/http ./cmd/http && CGO_ENABLED=0 go build -o ./target/cron ./cmd/cron

FROM alpine:3.20@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /app/target/ /app/bin/exec ./
ARG PUBLISH_MODE=production
COPY .env.${PUBLISH_MODE} ./.env
CMD ["./exec", "./http"]
# CMD ["./exec", "./cron"]