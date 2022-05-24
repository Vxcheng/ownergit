package behavioral

import (
	"fmt"
	"testing"
)

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	sub.Register(&Obsever1{})
	sub.Register(&Obsever2{})
	sub.Notify("hi")

	fmt.Println("now remove observer1...")
	sub.Remove(&Obsever1{})
	sub.Notify("hello")
}
