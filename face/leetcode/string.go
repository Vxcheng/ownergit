package leetcode

import (
	"strings"
	"sync"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	var result [][]int
	var preRight int
	for _, area := range intervals {
		if preRight == 0 || area[0] > preRight {
			preRight = area[1]
			result = append(result, []int{area[0], area[1]})
			continue
		}
		if area[1] > preRight {
			preRight = area[1]
			result[len(result)-1][1] = area[1]
		}
	}
	return result
}

func LongestDupSubstring(S string) string {
	low, high := 0, len(S)-1
	v := ""
	for low <= high {
		mid := low + (high-low)/2
		v1 := RabinKarp(S, mid)
		if v1 == "" {
			high = mid - 1
		} else {
			low = mid + 1
			v = v1
		}
	}
	return v
}
func RabinKarp(s string, length int) string {
	m := make(map[int]bool)
	now := 0
	r, mod := 256, 6*(1<<20)+1
	for i := 0; i < length; i++ {
		now = ((now * r) + int(s[i])) % mod
	}
	rm := 1
	for i := 1; i < length; i++ {
		rm = (rm * r) % mod
	}
	m[now] = true
	for i := length; i < len(s); i++ {
		now = (now - rm*int(s[i-length])%mod + mod) % mod
		now = (now*r + int(s[i])) % mod
		if m[now] && strings.Contains(s[:i], s[i-length+1:i+1]) {
			return s[i-length+1 : i+1]
		}
		m[now] = true
	}
	return ""
}

func para() {

	count := 5
	vec := make(map[int]int)
	wg := &sync.WaitGroup{}
	// mu := &sync.Mutex{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		// tmp := i
		func() {
			wg.Done()

			// mu.Lock()

			// vec[tmp] = tmp
			vec[i] = i
			// mu.Unlock()
		}()
	}

	wg.Wait()
	print("len: ", len(vec))
	return
}
