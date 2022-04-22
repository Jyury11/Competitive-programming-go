package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	input := make([]int, 5)
	stdin.Scan()
	for i, t := range strings.Split(stdin.Text(), " ") {
		input[i] = ParseInt(t)
	}

	ans := maxInt
	max := input[3] * 2
	if max < input[4]*2 {
		max = input[4] * 2
	}
	for i := 0; i <= max; i = i + 2 {
		a := 0
		if 0 < input[3]-i/2 {
			a = (input[3] - i/2) * input[0]
		}

		b := 0
		if 0 < input[4]-i/2 {
			b = (input[4] - i/2) * input[1]
		}
		price := a + b + input[2]*i
		if price < ans {
			ans = price
		}
	}

	fmt.Println(ans)
}

func ParseInt(s string) int {
	n := 0
	isStart := false
	sign := 1
	for _, r := range s {
		if !isStart {
			if r == '-' {
				sign = -1
				continue
			}
			sign = 1
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
