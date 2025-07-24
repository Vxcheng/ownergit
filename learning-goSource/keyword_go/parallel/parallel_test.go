package parallel

import (
	"runtime"
	"testing"
	"time"
)

func TestProduceConsume(t *testing.T) {
	t.Run("", func(t *testing.T) {
		RunProdCons()
	})
}

func TestPublisherSubscriber(t *testing.T) {
	t.Run("", func(t *testing.T) {
		RunPubSub()
	})
}

func TestRunPrimeNumer(t *testing.T) {
	t.Run("", func(t *testing.T) {
		RunPrimeNumer()
	})
	t.Run("", func(t *testing.T) {
		RunClose()
	})
	t.Run("", func(t *testing.T) {
		RunCloseByContext()
	})

	t.Run("", func(t *testing.T) {
		TryLock()
	})
}

func TestChan(t *testing.T) {
	t.Run("chan_stu3", func(t *testing.T) {
		chan_stu3()
	})

	t.Run("chan_stu4", func(t *testing.T) {
		t.Log(runtime.NumGoroutine())
		chan_stu4()
		time.Sleep(time.Second)
		runtime.GC()
		t.Log(runtime.NumGoroutine())
	})
}
