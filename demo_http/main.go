package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const Hport = "5001"

const (
	a, b = 1, 2
)

var list = []int{1, 2, 3, 4, 5}

var infers = map[interface{}]interface{}{
	1:       1,
	"hello": true,
}

func main() {
	for key, value := range infers {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	log.Printf("a: %d, b: %d \n", a, b)
	log.Println("running a http demo server...")
	http.HandleFunc("/", WebHookHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Hport), nil))
}

func WebHookHandler(w http.ResponseWriter, r *http.Request) {
	msgByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("buf.len is %d, ioutil.ReadAll() err: %s", len(msgByte), err.Error())
		return
	}

	log.Printf("alert msg: %s", string(msgByte))
	return
}
