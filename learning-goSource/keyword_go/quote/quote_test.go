package quote

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestList(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = [...]int{1, 2, 3} // a 是一个数组
		var b = &a                // b 是指向数组的指针
		fmt.Println(a[0], a[1])   // 打印数组的前2个元素
		fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似
		for i, v := range b {     // 通过数组指针迭代数组的元素
			fmt.Println(i, v)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("", func(t *testing.T) {
		str := "hello, world"
		fmt.Printf("%#v\n", str2bytes(str))

		buff := []byte("世界abc")
		fmt.Printf("%#v\n", bytes2str(buff))

		r := str2runes(buff)
		fmt.Printf("%#v\n", string(r))

		fmt.Printf("%#v\n", runes2string(r))
	})

	t.Run("", func(t *testing.T) {
		fmt.Printf("%#v\n", []rune("世界"))
		fmt.Printf("%#v\n", string([]rune{'世', '界'}))
	})

	t.Run("a", func(t *testing.T) {
		str := "hello, world"
		strBottom := (*reflect.StringHeader)(unsafe.Pointer(&str))
		println(strBottom, strBottom.Len, strBottom.Data)
	})

	t.Run("b", func(t *testing.T) {
		fmt.Printf("%#v\n", []byte("Hello, 世界"))
		fmt.Println(0xe7, 0x95, 0x8c)
		fmt.Println("\xe7\x95\x8c")

		bad := "\xe4\x00\x00\xe7\x95\x8cabc"
		fmt.Println(bad)
		for i, c := range bad {
			fmt.Println(i, c)
		}

		for i, c := range []byte("世界abc") {
			fmt.Println(i, c)
		}

		const s = "\xe4\x00\x00\xe7\x95\x8cabc"
		for i := 0; i < len(s); i++ {
			fmt.Printf("%d %x\n", i, s[i])
		}
	})
}

func TestSlice(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
		SortFloat64FastV1(a)

		var a1 = []float64{4, 2, 5, 7, 2, 1, 88, 1}
		SortFloat64FastV2(a1)
	})

	t.Run("", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		len := copy(a, a[1:])
		a = a[:len] // 删除开头1个元素 a = a[:copy(a, a[N:])]
		fmt.Printf("%#v\n", a)

		b := []int{1, 2, 3, 4}
		i := 1
		b = b[:i+copy(b[i:], b[i+1:])] // 删除中间1个元素
		// b = b[:i+copy(a[i:], b[i+N:])] // 删除中间N个元素
		fmt.Printf("%#v\n", b)
	})
}
