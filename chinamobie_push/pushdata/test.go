package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	userName    = "ZDATA"
	password    = "abcd1234"
	loginURL    = "http://192.168.195.97:8100/pushdata/pushdataLogin/login"
	contentType = "json/application"
)

func main() {
	fmt.Printf("测试开始...\n")
	result, err := login(userName, password, loginURL)
	if err != nil {
		log.Printf("login err:, %s \n", err.Error())
		return
	}
	log.Printf("success, result：%s", result)
}

func marshalMap(jMap map[string]string) (string, error) {
	jByte, err := json.Marshal(jMap)
	if err != nil {
		return "", err
	}
	return string(jByte), nil
}

func login(uName, pword, lUrl string) (string, error) {
	loginMap := make(map[string]string)
	loginMap["userName"] = uName
	loginMap["password"] = pword
	jStr, _ := marshalMap(loginMap)
	return PostJson(lUrl, jStr)
}

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
