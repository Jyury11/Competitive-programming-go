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
	searchStatus := make([]bool, n)
	searchStartTime := make([]int, n)
	searchEndTime := make([]int, n)

	for i := 0; i < n; i++ {
		stdin.Scan()
		input[i] = make([]int, 100)
		args := strings.Split(stdin.Text(), " ")
		for j, t := range args {
			input[i][j] = ParseInt(t)
		}
		input[i] = input[i][:len(args)]
		searchStatus[i] = false
		searchStartTime[i] = -1
		searchEndTime[i] = -1
	}

	// fmt.Println(input)
	// fmt.Println(searchStatus)

	time := 0
	for _, node := range input {
		numIndex := 0
		next := node[numIndex]
		if !searchStatus[next-1] {
			DepthFirstSearch(next-1, &input, &searchStatus, &searchStartTime, &searchEndTime, &time)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d %d %d\n", i+1, searchStartTime[i], searchEndTime[i])
	}
}

func DepthFirstSearch(index int, tree *[][]int, searchStatus *[]bool, searchStartTime, searchEndTime *[]int, time *int) {
	(*time)++
	(*searchStartTime)[index] = *time
	(*searchStatus)[index] = true
	relationNumIndex := 1

	node := (*tree)[index]
	n := node[relationNumIndex]
	for i := 0; i < n; i++ {
		next := node[relationNumIndex+1+i]
		if !(*searchStatus)[next-1] {
			DepthFirstSearch(next-1, tree, searchStatus, searchStartTime, searchEndTime, time)
		}
	}
	(*time)++
	(*searchEndTime)[index] = *time
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
