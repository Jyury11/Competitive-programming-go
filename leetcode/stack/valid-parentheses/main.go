package main

import "fmt"

const (
	circle int = iota
	square
	wave
)

func isValid(s string) bool {
	stack := []int{}
	pop := func() int {
		if len(stack) == 0 {
			return -1
		}
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return c
	}

	for _, char := range s {
		fmt.Println(char)
		switch char {
		case '(':
			stack = append(stack, circle)
		case ')':
			c := pop()
			if c != circle {
				return false
			}
		case '[':
			stack = append(stack, square)
		case ']':
			c := pop()
			if c != square {
				return false
			}
		case '{':
			stack = append(stack, wave)
		case '}':
			c := pop()
			if c != wave {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
