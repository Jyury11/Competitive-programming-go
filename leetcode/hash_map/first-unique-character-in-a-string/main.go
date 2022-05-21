package main

import "math"

func firstUniqChar(s string) int {
	m := make(map[rune]int, len(s))
	for i, r := range s {
		push(r, i, m)
	}

	min := math.MaxInt32
	for _, val := range m {
		if val < min {
			min = val
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

func push(r rune, i int, hashMap map[rune]int) {
	if _, exists := hashMap[r]; exists {
		hashMap[r] = math.MaxInt32
		return
	}
	hashMap[r] = i
}
