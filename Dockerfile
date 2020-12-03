FROM golang:1.15.5

ENV GOPROXY="https://goproxy.cn,direct"
RUN apt update && apt-get install -y ca-certificates

WORKDIR /go/src/alertmanager-wechatrobot-webhook
ADD alertmanager-wechatbot-webhook.go .
COPY transformer transformer
COPY notifier notifier
COPY model model
RUN ls -al
RUN go mod init
RUN go build -o alertmanager-wechatbot-webhook alertmanager-wechatbot-webhook.go

FROM ubuntu:18.04
COPY --from=0 /go/src/alertmanager-wechatrobot-webhook/alertmanager-wechatbot-webhook /alertmanager-wechatbot-webhook
ENTRYPOINT ["/alertmanager-wechatbot-webhook"]