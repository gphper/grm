#构建npm前端打包环境
FROM node:16 as nbuilder

WORKDIR /build

COPY ./web .

RUN npm install

RUN npm run build

#构建golang打包环境
FROM golang:alpine as gbuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
&& apk update \
&& apk upgrade

RUN apk update
Run apk add upx

WORKDIR /build

ADD . .

#将打包好的前端文件拷贝到gbuilder中
COPY --from=nbuilder /build/dist /build/web/dist

RUN go mod download

RUN go build -ldflags "-s -w" -o grm-build main.go

RUN upx -9 grm-build -o grmsrv

#构建最终的发布镜像
FROM alpine

ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=gbuilder /build/grmsrv /app/grmsrv

#添加管理用户 -u指定用户名 -p指定密码
RUN ./grmsrv user add -u=admin -p=123456
#运行服务
CMD ["./grmsrv","srv","run"]