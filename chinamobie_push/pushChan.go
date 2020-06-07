package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Pusher is push chan.
type Pusher struct {
	pushChan chan string
}

const (
	userName          = "ZDATA"
	password          = "abcd1234"
	isChinaMobilePush = true
	pchCap            = 10
	loginURL          = "http://192.168.195.97:8100/pushdata/pushdataLogin/login"
	pushURL           = "http://192.168.195.97:8100/pushdata/pushdataController/pushAlarmData"
	contentType       = "json/application"
)

func main() {
	if !isChinaMobilePush {
		fmt.Println("不是移动客户不需要推送")
		return
	}

	fmt.Println("移动API接口开始测试")
	p := &Pusher{
		pushChan: make(chan string, pchCap),
	}

	pushmsg := `[{
 "plat":"OneLink1",
 "product":"jizhongyewujiankongpingtai1",
 "originalMetric":"",
 "alarmType":0,
 "message":"enmo test...",
 "time":1578305062,
 "priority":1,
 "eventId":"",
 "state":0
},
{
 "plat":"OneLink2",
 "product":"jizhongyewujiankongpingtai2",
 "originalMetric":"",
 "alarmType":0,
 "message":"enmo test...",
 "time":1578305066,
 "priority":1,
 "eventId":"",
 "state":0
}]`
	go func(data string) {
		for {
			go p.writeToChan(data) // 测试平台定时产生告警消息
			time.Sleep(time.Second * 3)
		}
	}(pushmsg)

	var count int
	for data := range p.pushChan { // 读取告警消息请求API
		count++
		fmt.Printf("第%d个元素，值：'%s'\n", count, data)
		PushAlerm(data)
	}

}

func (p *Pusher) writeToChan(data string) {
	p.pushChan <- data
}

var accessToken string

type ResResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func PushAlerm(requestJsonStr string) {
	var result string
	resultMap := ResResult{}
	if accessToken == "" {
		result = exeLogin(userName, password, loginURL, pushURL, requestJsonStr)
		log.Printf("exeLogin result: %s \n", result)
		return
	}
	result, err := push(pushURL, requestJsonStr)
	if err != nil {
		log.Printf("push err: %v \n", err.Error())
		return
	}
	json.Unmarshal([]byte(result), &resultMap)
	if resultMap.Code == -1 {
		result = exeLogin(userName, password, loginURL, pushURL, requestJsonStr)
		log.Printf("push code=-1, exeLogin result: %s \n", result)
	} else {
		log.Printf("push result: %s \n", result)
	}
	return

}

func exeLogin(uName, pword, lUrl, pUrl, requestJsonStr string) string {
	var result string
	result, err := login(uName, pword, lUrl)
	if err != nil {
		log.Printf("login err: %v \n", err.Error())
		return result
	}
	resultMap := ResResult{}
	json.Unmarshal([]byte(result), &resultMap)
	log.Println("resultMap:", resultMap)
	if resultMap.Code == 0 {
		accessToken = resultMap.Data.(string)
		result, err = push(pushURL, requestJsonStr)
		if err != nil {
			log.Printf("push err: %v \n", err.Error())
			return result
		}
	} else {
		log.Printf("login 失败: %s \n", result)
	}
	return result
}

func push(pUrl, requestJsonStr string) (string, error) {
	pushMap := make(map[string]string)
	pushMap["accessToken"] = accessToken
	pushMap["requestJsonStr"] = requestJsonStr
	jStr, _ := marshalMap(pushMap)
	return PostJson(pUrl, jStr)
}

func login(uName, pword, lUrl string) (string, error) {
	loginMap := make(map[string]string)
	loginMap["userName"] = uName
	loginMap["password"] = pword
	jStr, _ := marshalMap(loginMap)
	return PostJson(lUrl, jStr)
}

func marshalMap(jMap map[string]string) (string, error) {
	jByte, err := json.Marshal(jMap)
	if err != nil {
		return "", err
	}
	return string(jByte), nil
}

// HTTPCLI provides http indivate
func PostJson(url string, reqBody string) (result string, err error) {
	method := "POST"
	bBody := strings.NewReader(reqBody)
	switch method {
	case http.MethodGet, "Get":
		bBody = nil
	case http.MethodPost, "Post":
	case http.MethodPut, "Put":
	case http.MethodDelete, "Delete":
	default:
		fmt.Printf("can't handle '%s' method\n", method)
	}

	r, err := http.NewRequest(method, url, bBody)
	if err != nil {
		return
	}
	r.Header.Set("Content-Type", contentType)

	c := &http.Client{}
	res, err := c.Do(r)
	if err != nil {
		return
	}
	defer c.CloseIdleConnections()

	if res.StatusCode == http.StatusOK {
		var resByte []byte
		resByte, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return
		}
		result = string(resByte)
	} else {
		// code = res.StatusCode
		err = errors.New("res.StatusCode is not 200")
	}

	return
}
