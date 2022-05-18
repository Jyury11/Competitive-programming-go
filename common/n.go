package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	n := ParseInt(strings.TrimSpace(stdin.Text()))
	input := make([][]int, n)

	for i := 0; i < n; i++ {
		stdin.Scan()
		input[i] = make([]int, 100)
		args := strings.Split(stdin.Text(), " ")
		for j, t := range args {
			input[i][j] = ParseInt(t)
		}
		input[i] = input[i][:len(args)]
	}

	fmt.Println(input)
}

func ParseInt(s string) int {
	n := 0
	isStart := false
	sign := 1
	for _, r := range s {
		if !isStart {
			if '-' == r {
				sign = -1
				continue
			} else {
				sign = 1
			}
		}

		if '0' <= r && r <= '9' {
			n = n*10 + int(r-'0')
			isStart = true
		} else if isStart {
			break
		}
	}
	return n * sign
}
