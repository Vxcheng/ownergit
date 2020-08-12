package main

import (
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	fmt.Println("buf")
	buf, err := ioutil.ReadFile("./a.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}