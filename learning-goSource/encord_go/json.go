package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

type User struct {
	Name string
	Addr `mapstructure:",squash"`
}

type Addr struct {
	Dong string
}

func main() {
	convert()
}

func convert() {
	var max int64 = 10
	log.Print(int(max))
}

func unmarshal() {
	str := `
	{
		"name": "ming"
	}
	`
	// var u User
	err := json.Unmarshal([]byte(str), &struct{}{})
	if err != nil {
		log.Fatal(err)
	}
}

func baseDecode() {
	password, _ := base64.StdEncoding.DecodeString("cm9vdDEyMw==")
	log.Println("pwd, ", string(password))
}
