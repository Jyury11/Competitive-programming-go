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
	nodeDistance := make([]int, n)

	for i := 0; i < n; i++ {
		stdin.Scan()
		input[i] = make([]int, 100)
		args := strings.Split(stdin.Text(), " ")
		for j, t := range args {
			input[i][j] = ParseInt(t)
		}
		input[i] = input[i][:len(args)]
		searchStatus[i] = false
		nodeDistance[i] = -1
	}

	// fmt.Println(input)
	// fmt.Println(searchStatus)

	// ループ
	// for i := range input {
	// 	distance := 0
	// 	if !searchStatus[i] {
	// 		BreadthFirstSearch(i, &input, &searchStatus, &nodeDistance, &distance)
	// 	}
	// }

	// ルートから
	distance := 0
	BreadthFirstSearch(0, &input, &searchStatus, &nodeDistance, &distance)

	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n", i+1, nodeDistance[i])
	}
}

func BreadthFirstSearch(index int, tree *[][]int, searchStatus *[]bool, nodeDistance *[]int, distance *int) {
	(*nodeDistance)[index] = *distance
	(*searchStatus)[index] = true
	(*distance)++

	buffer := make([]int, 0, len(*tree))
	relationNumIndex := 1

	node := (*tree)[index]
	for i := 0; i < node[relationNumIndex]; i++ {
		next := node[relationNumIndex+1+i]
		if !(*searchStatus)[next-1] {
			buffer = append(buffer, next-1)
			(*nodeDistance)[next-1] = *distance
			(*searchStatus)[next-1] = true
		}
	}

	(*distance)++
	for {
		if len(buffer) == 0 {
			break
		}

		buf := buffer
		buffer = make([]int, 0, len(*tree))

		for _, childIndex := range buf {
			child := (*tree)[childIndex]
			for i := 0; i < child[relationNumIndex]; i++ {
				next := child[relationNumIndex+1+i]
				if !(*searchStatus)[next-1] {
					buffer = append(buffer, next-1)
					(*nodeDistance)[next-1] = *distance
					(*searchStatus)[next-1] = true
				}
			}
		}
		(*distance)++
	}

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
