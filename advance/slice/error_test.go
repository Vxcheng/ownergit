package slice

import (
	"errors"
	"fmt"
	"testing"
)

type MyError struct {
	s string
}

func (e *MyError) Error() string {
	return e.s
}

func TestError(t *testing.T) {
	e1 := errors.New("error1")
	e2 := errors.New("error2")
	e3 := errors.New("error3")
	e4 := &MyError{
		s: "error4",
	}
	e := fmt.Errorf("%w, %w, %w, %w", e1, e2, e3, e4)

	fmt.Printf("e = %s\n", e.Error()) // error1 error2, error3, error4
	fmt.Println(errors.Is(e, e1))     // true

	var ne *MyError
	fmt.Println(errors.As(e, &ne)) // true
	fmt.Println(ne == e4)          // true
}
