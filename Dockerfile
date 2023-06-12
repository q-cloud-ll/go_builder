# golang:1.19-alpine 这个版本号和自己的go.mod 对应，如果不一致可能会出现镜像构建失败，go和gomod有版本差
FROM golang:1.19-alpine as builder
WORKDIR /app/src/go_builder
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o main .

FROM alpine:latest

LABEL MAINTAINER="cherryqcsj@gmail.com"

WORKDIR /app/src/go_builder

COPY --from=0 /app/src/go_builder/main ./
COPY --from=0 /app/src/go_builder/conf/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./main -c conf/config.docker.yaml
