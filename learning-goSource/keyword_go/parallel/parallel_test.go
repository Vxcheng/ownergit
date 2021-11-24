package parallel

import "testing"

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
}
