package parallel

import "testing"

func TestRace(t *testing.T) {
	t.Run("race_stu1", func(t *testing.T) {
		race_stu1()
	})

	t.Run("race_stu2", func(t *testing.T) {
		race_stu2()
	})

	t.Run("race_stu3", func(t *testing.T) {
		race_stu3()
	})
}
