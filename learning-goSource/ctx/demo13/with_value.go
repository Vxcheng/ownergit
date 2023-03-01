package main

import (
	"context"

	"fmt"
)

type akeytype string

type ckeytype string

var keyA akeytype = "keyA"

var keyC ckeytype = "keyA"

//
//var keyA string = "keyA"
//var keyC string = "keyA"

func main() {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, keyA, "valueA")

	fmt.Println("In ctx:")
	fmt.Println("keyA => ", ctx1.Value(keyA))

	ctx2 := context.WithValue(ctx1, keyC, "valueC")
	fmt.Println("In ctx2:")
	fmt.Println("keyA => ", ctx2.Value(keyA))
	fmt.Println("keyC => ", ctx2.Value(keyC))
	//fmt.Println("keyA => ", ctx1.Value(keyA))

	return

}
