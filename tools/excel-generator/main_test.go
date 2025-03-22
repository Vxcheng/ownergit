package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func TestGenerateFile(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	t.Run("generateTemplateValue", func(t *testing.T) {
		cases := []struct {
			template Template
		}{
			{
				Template{
					Type:    "int",
					Pattern: "@int{-30,30}",
					Max:     30,
					Min:     -30,
				},
			},
			{
				Template{
					Type:    "float",
					Pattern: "@float.2{25,26}",
					Max:     26,
					Min:     25,
					Decimal: 2,
				},
			},
		}

		for _, cc := range cases {
			rule := Rule{
				Type:           cc.template.Type,
				Decimal:        cc.template.Decimal,
				Max:            cc.template.Max,
				Min:            cc.template.Min,
				RandomFloating: cc.template.RandomFloating,
			}
			v := generateTemplateValue(rule)
			t.Log(v)
		}
	})

	t.Run("FindStringSubmatch", func(t *testing.T) {
		cc := []struct {
			text string
			expr string
		}{
			{expr: "@int", text: "hello@int"},
			{expr: "@int<-30,30>", text: "xx@int<-30,30>,@int<-30,30>"},
			{expr: "@float.1<2.001,2.5>", text: "xx@float.1<2.001,2.5>,@float.1<2.001,2.5>"},
			{expr: "$float.2[25,26.22]", text: "长度：$float.2[25,26.22],$float.2[25,26.22],$float.2[25,26.22]"},
		}
		for _, c := range cc {
			expr := c.expr
			expr = regexp.QuoteMeta(expr)
			expr = fmt.Sprintf(`（%s）`, expr)
			re, err := regexp.Compile(expr)
			if err != nil {
				continue
			}
			ff := re.FindStringSubmatch(c.text)
			text := re.ReplaceAllString(c.text, fmt.Sprintf("#len%d", len(c.expr)))
			t.Log(ff, text)
		}

	})

	t.Run("rand", func(t *testing.T) {
		// 初始化随机数生成器的种子
		rand.Seed(time.Now().UnixNano())

		// 生成 5 个 [-10, 10] 范围内的随机整数和浮点数
		for i := 0; i < 5; i++ {
			randomInt := rand.Intn(21) - 10
			randomFloat := -10 + rand.Float64()*20
			fmt.Printf("Random integer: %d, Random float: %.2f\n", randomInt, randomFloat)

		}
	})

	t.Run("count", func(t *testing.T) {
		v := float64(1) / float64(5)
		t.Log(v, v*100)
	})
}
