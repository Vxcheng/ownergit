package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"ownergit/face/search"
	"strings"
)

func main() {

	http.HandleFunc("/", sayHelloHandler) //	设置访问路由

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "aaa.txt" // multi
	srcDir := "/home"
	files, err := search.SearchFileQuickly(srcDir, fileName)
	if err != nil {
		log.Fatal()
	}

	fmt.Fprintf(w, strings.Join(files, "\n")) //这个写入到w的是输出到客户端的
}
