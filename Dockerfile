FROM golang:1.23-alpine AS builder
RUN go env -w GO111MODULE=on GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -o ./target/http ./cmd/http && CGO_ENABLED=0 go build -o ./target/cron ./cmd/cron

FROM alpine
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