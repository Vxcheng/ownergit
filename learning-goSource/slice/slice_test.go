package slice

import "testing"

func TestDemo(t *testing.T) {
	t.Run("lastNumsBySlice", func(t *testing.T) {
		origin := []int{1, 2, 3, 4, 5}
		result := lastNumsBySlice(origin)
		t.Logf("result: %v", result)
	})

	t.Run("lastNumsByCopy", func(t *testing.T) {
		origin := []int{1, 2, 3, 4, 5}
		result := lastNumsByCopy(origin)
		t.Logf("result: %v", result)
	})

	t.Run("DeleteSlice3", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		elem := 3
		result := DeleteSlice3(s, elem)
		t.Logf("result: %v", result)
	})

}

func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}

// DeleteSlice3 删除指定元素。
func DeleteSlice3(s []int, elem int) []int {
	j := 0
	for _, v := range s {
		if v != elem {
			s[j] = v
			j++
		}
	}
	return s[:j]
}
