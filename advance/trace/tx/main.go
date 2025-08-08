package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/trace"
	"strings"
)

type document struct {
	Channel struct {
		Items []struct {
			Title string `xml:"title"`
		} `xml:"item"`
	} `xml:"channel"`
}

func freq(docs []string) int {
	var count int
	for _, doc := range docs {
		f, err := os.OpenFile(doc, os.O_RDONLY, 0)
		if err != nil {
			return 0
		}
		data, err := io.ReadAll(f)
		if err != nil {
			return 0
		}
		var d document
		if err := xml.Unmarshal(data, &d); err != nil {
			log.Printf("Decoding Document [Ns] : ERROR :%+v", err)
			return 0
		}
		for _, item := range d.Channel.Items {
			if strings.Contains(strings.ToLower(item.Title), "go") {
				count++
			}
		}
	}
	return count
}

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()
	files := make([]string, 0)
	for i := 0; i < 100; i++ {
		files = append(files, "index.xml")
	}
	count := freq(files)
	log.Println(fmt.Sprintf("find key word go %d count", count))
}
