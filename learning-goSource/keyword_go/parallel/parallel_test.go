package parallel

import (
	"testing"
)

func Test_Stu_waitGroup(t *testing.T) {
	t.Run("Test_Stu_waitGroup----", func(t *testing.T) {
		Stu_waitGroup(NewUser(), 4)
		return
	})
}
