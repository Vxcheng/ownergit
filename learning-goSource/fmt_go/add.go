package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"encoding/json"
)

func Add(a, b uint64) uint64 {
	return uint64(a + b)
}

func Append() {
	slic := make([]string, 0)
	child := signal()
	for _, _ = range child {

	}

	slic = append(slic, child...)
}

func signal() []string {
	return nil
}

func promote() {
	type User struct {
		Name string
		Age   int
	}

	type Team struct {
		User
		Age string
	}

	u := User{
		Name: "xiaoming",
		Age: 10,
	}
	t := Team{
		User: u,
		Age: "20",
	}
	fmt.Printf("t: %+v\n", t)

	buffT, err := json.Marshal(t)
	if err != nil {
		return
	}
	fmt.Printf("t_str: %s\n", string(buffT))
}

func main() {
	promote()
	Append()
	convert1()
	similarTest()

	l := strings.Split("a", ",")
	fmt.Println(l)
	printSN()

	fmt.Println(Add(2, 15))
	fmt.Printf("node: %v\n", Storage)
	a := int64(10)
	fmt.Printf("fl64: %f\n", float64(a))
	b := float64(1)
	if b == 1 {
		fmt.Println("===")
	}

	c := int64(2)
	fmt.Printf("int: %d\n", int(c))
	d := []string{"aaa", "0.0"}
	dF, _ := strconv.ParseFloat(d[1], 64)
	fmt.Println(dF)
	out := Formatted("DD/_DROPPED_0002_OCR")
	fmt.Println("out: ", out)
	convertString()
}

func convert1() {
	baseLid := "0x1f"
	i, err := strconv.ParseInt(baseLid, 0, strconv.IntSize)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(i)
}

type User struct {
	Name string
	Age  int
}

func similarTest() {
	var u1 *User
	u1 = &User{
		Name: "u1",
		Age:  2,
	}
	u2 := *u1
	fmt.Printf("u1: %+v, u2: %+v\n", u1, u2)
}

func Formatted(key string) string {

	aa := strings.Fields(key)
	return strings.ToLower(strings.Join(aa, "_"))
}

type NodeType int64

const (
	Storage NodeType = iota + 1
	Compute
	Fusion
	Quorum
)

func convertString() {
	str := "hel"
	fmt.Println("index: ", str[2])
	for i, v := range str {
		fmt.Println("i: ", i, ", v:", v)
	}

	runs := []rune(str)
	for i, v := range runs {
		fmt.Println("i: ", i, ", s:", string(v))
	}
}
