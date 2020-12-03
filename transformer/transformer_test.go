package transformer

import (
	"alertmanager-wechatrobot-webhook/model"
	"testing"
	"time"
)

func TestToContent(t *testing.T){
	notification := model.Notification{
		Version:           "4",
		GroupKey:          `{}/{type="waf"}:{domain="topic.eqxiu.com"}`,
		Status:            "resolved",
		Receiver:          "debug",
		GroupLabels: map[string]string{
			"domain": "topic.eqxiu.com",
		},
		CommonLabels: map[string]string{
			"alertname": "HTTP 5xx 1分钟内次数总和大于 10",
			"domain": "topic.eqxiu.com",
			"severity": "warning",
			"type": "waf",
		},
		CommonAnnotations: map[string]string{
			"description": "$value",
			"summary": "HTTP 5xx 1分钟内次数总和大于 10",
		},
		ExternalURL:       "http://clickhouse01.eqxiu.com:9093",
		Alerts:            []model.Alert{
			model.Alert{
				Labels: map[string]string{
					"alertname": "HTTP 5xx 1分钟内次数总和大于 10",
					"domain": "topic.eqxiu.com",
					"severity": "warning",
					"type": "waf",
				},
				Annotations: map[string]string{
					"description": "$value",
					"summary": "HTTP 5xx 1分钟内次数总和大于 10",
				},
				StartsAt:    time.Now(),
				EndsAt:      time.Now(),
			},
			model.Alert{
				Labels: map[string]string{
					"alertname": "HTTP 5xx 1分钟内次数总和大于 10",
					"domain": "topic.eqxiu.com",
					"severity": "warning",
					"type": "waf",
				},
				Annotations: map[string]string{
					"description": "$value",
					"summary": "HTTP 5xx 1分钟内次数总和大于 10",
				},
				StartsAt:    time.Now(),
				EndsAt:      time.Now(),
			},
		},
	}

	content, _  := toContent(notification)

	println(content)

}
