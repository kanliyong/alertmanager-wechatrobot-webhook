package notifier

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"alertmanager-wechatrobot-webhook/model"
	"alertmanager-wechatrobot-webhook/transformer"
)

// Send send markdown message to dingtalk
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}

	SendMarkDown(markdown, robotURL, defaultRobot)
	return
}

func SendMarkDown(markdown *model.WeChatMarkdown,robotURL string, robot string){
	data, err := json.Marshal(markdown)

	println(data)
	if err != nil {
		return
	}

	var wechatRobotURL string
	if robotURL != ""{
		wechatRobotURL = robotURL
	}else{
		wechatRobotURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + robot
	}

	req, err := http.NewRequest(
		"POST",
		wechatRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		//TLSClientConfig:        &tls.Config{InsecureSkipVerify:true},
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	return
}