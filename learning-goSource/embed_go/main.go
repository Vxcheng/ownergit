package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed static/js/aaa.txt
var content string

//go:embed *
var fileTree embed.FS

func main() {
	fmt.Println("content: ", content)
	http.Handle("/", http.FileServer(http.FS(fileTree)))
	log.Fatal(http.ListenAndServe(":9001", nil))
}
