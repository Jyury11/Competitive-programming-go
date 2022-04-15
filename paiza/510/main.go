package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var moveR = []map[string]bool{
	{
		"R": true,
		"G": false,
		"B": false,
		"Y": true,
		"M": true,
		"C": false,
		"W": true,
	},
	{
		"R": false,
		"G": true,
		"B": false,
		"Y": true,
		"M": false,
		"C": true,
		"W": true,
	},
	{
		"R": false,
		"G": false,
		"B": true,
		"Y": false,
		"M": true,
		"C": true,
		"W": true,
	},
}

var moveLen = map[string]int{
	"L": -1,
	"R": 1,
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := stdin.Text()
	n := ParseInt(s)

	kabuts := make([]int, 3)
	stdin.Scan()
	for i, kabuPos := range strings.Split(stdin.Text(), " ") {
		kabuts[i] = ParseInt(kabuPos)
	}

	end := false
	ret := 0
	for i := 0; i < n; i++ {
		ps := make([]string, 2)
		stdin.Scan()
		if end {
			continue
		}

		for j, p := range strings.Split(stdin.Text(), " ") {
			ps[j] = p
		}

		kabuts = Move(kabuts, ps)

		if kabuts[0] == kabuts[1] && kabuts[1] == kabuts[2] {
			ret = kabuts[0]
			end = true
		}
	}

	if end {
		fmt.Println(ret)
		return
	}

	fmt.Println("no")
}

func Move(kabuts []int, ps []string) []int {
	ret := make([]int, 3)
	for i, kabut := range kabuts {
		if moveR[i][ps[1]] {
			ret[i] = kabut + moveLen[ps[0]]
		} else {
			ret[i] = kabut
		}
	}
	return ret
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
