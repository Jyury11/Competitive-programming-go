package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	table := make([]int, 2)
	stdin.Scan()
	for i, t := range strings.Split(stdin.Text(), " ") {
		table[i] = ParseInt(t)
	}

	end := false

	rows := 0
	cols := make([]bool, table[1])
	bomRows := 0

	for i := 0; i < table[0]; i++ {
		stdin.Scan()
		if end {
			continue
		}

		isBom := false
		for j, bomOrNull := range stdin.Text() {
			if bomOrNull == '#' {
				isBom = true
				cols[j] = true
			}
		}

		if !isBom {
			rows++
		} else {
			bomRows++
		}
	}

	bomCols := 0
	for _, col := range cols {
		if col {
			bomCols++
		}
	}

	fmt.Println(bomRows*table[1] + rows*bomCols)
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
