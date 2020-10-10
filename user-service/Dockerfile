FROM golang:1.14-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache git

WORKDIR /app/laracom-user-service

COPY . .

RUN go mod download

# 构建二进制文件，添加来一些额外参数以便可以在 Alpine 中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-user-service

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash supervisor ca-certificates

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/laracom-user-service/laracom-user-service .

CMD ["./laracom-user-service"]