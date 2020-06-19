package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	got := Add1(2)
	if !assert.Equal(t, 3, got) {
		t.Fatalf("Add1() got: %d \n", got)
	}

	a := 2
	Sub1(&a)
	if !assert.Equal(t, 1, a) {
		t.Fatalf("Sub1() got: %d \n", a)
	}

	b := 2
	SubNoPoint(b)
	if !assert.Equal(t, 2, b) {
		t.Fatalf("Sub1() got: %d \n", a)
	}
}

var globalSub = 10

func TestUser(t *testing.T) {
	Convey("TestUser", t, func() {
		Convey("TalkName", func() {
			u := &User{Name: "xiaoming", Sex: true, Age: 17}
			fmt.Printf("u variable: %v\n", u)
			u.TalkName()
			u.SaySex()
			(u).Growing()
		})
		Convey("monkey", func() {
			gomonkey.ApplyFunc(Add1, func(num int) int {
				fmt.Printf("monkey add1\n")
				return num
			})
			Add1(1)

			u := &User{Name: "xiaoming", Sex: true, Age: 17}
			gomonkey.ApplyMethod(reflect.TypeOf(*u), "TalkName", func(_ User) {
				fmt.Printf("monkey TalkName\n")
				// return nil
			})
			u.TalkName()
		})

		Convey("globalSub", func() {
			patches := gomonkey.ApplyGlobalVar(&globalSub, 120)
			defer patches.Reset()

			fmt.Printf("globalSub: %d\n", globalSub)
		})

		Convey("default times is 1", func() {
			info1 := 1
			info2 := 2
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{info1}},
				{Values: gomonkey.Params{info2}},
			}
			patches := gomonkey.ApplyFuncSeq(Add1, outputs)
			defer patches.Reset()
			output := Add1(1)
			So(output, ShouldEqual, info1)
			output = Add1(1)
			So(output, ShouldEqual, info2)
		})
	})
	fmt.Printf("out globalSub: %d\n", globalSub)

	t.Run("UnSafePoint", func(t *testing.T) {
		UnSafePoint()
	})
}
