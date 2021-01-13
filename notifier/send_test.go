package notifier

import (
	"alertmanager-wechatrobot-webhook/model"
	"log"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	notification := model.Notification{
		Version:  "4",
		GroupKey: `{}/{type="waf"}:{domain="topic.eqxiu.com"}`,
		Status:   "firing",
		Receiver: "debug",
		GroupLabels: map[string]string{
			"domain": "topic.eqxiu.com",
		},
		CommonLabels: map[string]string{
			"alertname": "HTTP 5xx 1分钟内次数总和大于 10",
			"domain":    "topic.eqxiu.com",
			"severity":  "warning",
			"type":      "waf",
		},
		CommonAnnotations: map[string]string{
			"description": "$value",
			"summary":     "HTTP 5xx 1分钟内次数总和大于 10",
		},
		ExternalURL: "http://clickhouse01.eqxiu.com:9093",
		Alerts: []model.Alert{
			model.Alert{
				Labels: map[string]string{
					"alertname": "HTTP 5xx 1分钟内次数总和大于 10",
					"domain":    "topic.eqxiu.com",
					"severity":  "warning",
					"type":      "waf",
				},
				Annotations: map[string]string{
					"description": "$value",
					"summary":     "HTTP 5xx 1分钟内次数总和大于 10",
				},
				EndsAt: time.Now(),
			}, model.Alert{
				Labels: map[string]string{
					"alertname": "HTTP 5xx 1分钟内次数总和大于 10",
					"domain":    "topic.eqxiu.com",
					"severity":  "warning",
					"type":      "waf",
				},
				Annotations: map[string]string{
					"description": "$value",
					"summary":     "HTTP 5xx 1分钟内次数总和大于 10",
				},
				StartsAt: time.Now(),
				EndsAt:   time.Now(),
			}, model.Alert{
				Labels: map[string]string{
					"alertname": "HTTP 5xx 1分钟内次数总和大于 10",
					"domain":    "topic.eqxiu.com",
					"severity":  "warning",
					"type":      "waf",
				},
				Annotations: map[string]string{
					"description": "$value",
					"summary":     "HTTP 5xx 1分钟内次数总和大于 10",
				},
				StartsAt: time.Now(),
			},
		},
	}

	log.Println(notification)
	Send(notification, "0717e9e1-cc70-4d1a-b36c-99c1584d77d7")
}

func TestTimzone(t *testing.T) {
	println(time.Now().Format("2006-01-02T15:04:05Z07:00"))
}
