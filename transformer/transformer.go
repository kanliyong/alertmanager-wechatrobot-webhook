package transformer

import (
	"alertmanager-wechatrobot-webhook/model"
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"time"
)

// TransformToMarkdown transform alertmanager notification to wechat markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	content, err := toContent(notification)
	buffer.WriteString(content)

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content:  buffer.String(),
		},
	}

	return
}


func toContent(notification model.Notification) (content string, err error){

	templateString := templateString()
	out := bytes.NewBuffer([]byte{})

	funcMap := template.FuncMap{"fdate": formDate}
	t ,err := template.New("test").Funcs(funcMap).Parse(templateString)
	if err != nil {
		return
	}
	err = t.Execute(out,  notification)
	return out.String(), nil
}

func formDate(t time.Time) string{
	return t.Format("02 Jan 06 15:04:05")
}

var defaultTemplateString = `# {{ if eq .Status "resolved"}}<font color="info">恢复</font>{{ else if eq .Status "firing"}}<font color="warning">触发</font>{{end}}  {{.CommonLabels.alertname}}  
{{if .CommonLabels.severity          }}## 级别: <font color="warning">{{.CommonLabels.severity}}</font>  {{end}} 
{{ if .CommonAnnotations.description }}## 描述: {{.CommonAnnotations.description }} {{end}}
{{ if .CommonAnnotations.summary     }}## 汇总: {{.CommonAnnotations.summary}}      {{end}}
{{ range .Alerts}}
------  
### {{.Labels.alertname}}
{{if .Annotations.description }}#### 描述: {{.Annotations.description}} {{end}}
{{if .Annotations.summary }}#### 汇总: {{.Annotations.summary}} {{end}}
##### 标签: 
{{ range $key, $value := .Labels}}
1. {{$key}}: {{$value}}  
{{end}}
{{ if not .StartsAt.IsZero }}触发时间 {{.StartsAt | fdate}} {{end}} 
{{ if not .EndsAt.IsZero }}恢复时间 {{.EndsAt | fdate}}   {{end}}
{{end}}
`
func templateString() string{
	filePath := os.Getenv("template_path")
	if filePath == "" {
		return defaultTemplateString
	}
	file, err := os.Open(filePath)
	if err != nil {
		return defaultTemplateString
	}

	b , err := ioutil.ReadAll(file)
	if err != nil {
		return defaultTemplateString
	}
	return string(b)
}