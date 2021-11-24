package workpool

import (
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	s := NewScheduler()
	s.Run()

	for i := 0; i < 6; i++ {
		s.AddJob(&Payload{id: i + 1})
	}

	time.Sleep(time.Second * 5)
	s.Close()
	time.Sleep(time.Second * 10)
}
