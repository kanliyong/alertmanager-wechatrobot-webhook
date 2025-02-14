# alertmanager-wechatrobot-webhook
Convert alertmanger webhook messages to qiye WeChat robots that can receive messages

## build
```bash
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  alertmanager-wechatbot-webhook.go

docker build -t ccr.ccs.tencentyun.com/eqxiu/alertmanager-wechatbot-webhook  .
docker push ccr.ccs.tencentyun.com/eqxiu/alertmanager-wechatbot-webhook


./nginx-kafaka-exporter -servers=hadoop104.eqxiu.com:9092,hadoop105.eqxiu.com:9092,hadoop106.eqxiu.com:9092 \
-group_id=gateway_original_test\
-topic=gateway_original
```

## How use
```
./alertmanager-wechatbot-webhook  --RobotKey="xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
```
或
```
go run alertmanager-wechatbot-webhook.go  --RobotKey="xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
```
## configure

alertmanager.yml 
```
receivers:
- name: webhook-test                                                                                                           
  webhook_configs:                                                                                                             
  - url: 'http://127.0.0.1:8999/webhook?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx'
```
或者在prometheus rules文件中指定微信机器人
```
groups:
- name: ansible managed alert rules
  rules:
  - alert: CriticalCPULoad
    expr: (100 * (1 - avg(irate(node_cpu_seconds_total{mode="idle"}[5m])) BY (instance)))
      > 96
    for: 2m
    labels:
      severity: critical
    annotations:
      description: '{{ $labels.instance }} of mountpoint {{ $labels.mountpoint }} has
        Critical CPU load for more than 2 minutes.'
      summary: Instance {{ $labels.instance }} - Critical CPU load
      wechatRobot: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxx-xxxxxxxx-xxxxx-xxxxxx"

```

## test

### 1、使用默认微信机器人发送消息
```
curl 'http://127.0.0.1:8999/webhook'  -H 'Content-Type: application/json'    -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "test alert1",
          "instance": "example1",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message,ignore",
          "description": "This is test summary,ignore",
          "summary": "This is test message,ignore"
        },
        "startsAt": "2019-07-24T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "test alert1"
    },
    "commonLabels": {
      "alertname": "test alert1",
      "instance": "example1"
    },
    "commonAnnotations": {
      "info": "Test message,ignore",
      "summary": "This is test summary,ignore"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"test alert1\"}"
  }'
```

### 2、使用指定微信机器人发送消息
```
curl 'http://127.0.0.1:8999/webhook?key=0717e9e1-cc70-4d1a-b36c-99c1584d77d7'  -H 'Content-Type: application/json'    -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "test alert1",
          "instance": "example1",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message,ignore",
          "description": "This is test summary,ignore",
          "summary": "This is test message,ignore"
        },
        "startsAt": "2020-01-14T14:38:30+00:00",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "test alert1"
    },
    "commonLabels": {
      "alertname": "test alert1",
      "instance": "example1"
    },
    "commonAnnotations": {
      "info": "Test message,ignore",
      "summary": "This is test summary,ignore"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"test alert1\"}"
  }'
```

### 3、使用alert内容指定微信机器人发送消息
```
  curl 'http://127.0.0.1:8999/webhook'  -H 'Content-Type: application/json'    -d '
  {
    "receiver": "webhook-test",
    "status": "firing",
    "alerts": [
      {
        "status": "firing",
        "labels": {
          "alertname": "test alert1",
          "instance": "example1",
          "severity": "critical"
        },
        "annotations": {
          "info": "Test message,ignore",
          "description": "This is test summary,ignore",
          "summary": "This is test message,ignore"
        },
        "startsAt": "2019-07-24T02:38:30.763785079Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "critical"
      }
    ],
    "groupLabels": {
      "alertname": "test alert1"
    },
    "commonLabels": {
      "alertname": "test alert1",
      "instance": "example1"
    },
    "commonAnnotations": {
      "info": "Test message,ignore",
      "summary": "This is test summary,ignore",
      "wechatRobot": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx-xxxxx-xxxxx-xxxxxx-xxxxxxx"
    },
    "externalURL": "http://localhost:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"^(?:test.*)$\"}:{alertname=\"test alert1\"}"
  }'
```
