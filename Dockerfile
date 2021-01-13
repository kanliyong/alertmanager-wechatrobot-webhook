FROM golang:1.15.5

ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /go/src/alertmanager-wechatrobot-webhook
ADD alertmanager-wechatbot-webhook.go .
COPY transformer transformer
COPY notifier notifier
COPY model model
RUN ls -al
RUN go mod init
RUN go build -o alertmanager-wechatbot-webhook alertmanager-wechatbot-webhook.go

FROM ubuntu:18.04

RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN  apt-get clean
## 时区设置
ENV TZ=Asia/Shanghai
RUN apt-get update \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && apt-get -y install tzdata \
    && apt-get -y clean \
    && apt-get -y autoclean 
    
RUN apt-get -y install ca-certificates

COPY --from=0 /go/src/alertmanager-wechatrobot-webhook/alertmanager-wechatbot-webhook /alertmanager-wechatbot-webhook
ENTRYPOINT ["/alertmanager-wechatbot-webhook"]