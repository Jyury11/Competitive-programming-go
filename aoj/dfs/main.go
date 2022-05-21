package main

import (
	"bufio"
	"os"
	"strings"
)

type Answer struct {
	start int
	end   int
}

var (
	graph  = [][]int{}
	seen   = []bool{}
	answer = []Answer{}
)

// 深さ優先探索
func dfs(graph [][]int, v, cnt int) {
	cnt++
	seen[v] = true // v を訪問済にする
	answer[v].start = cnt

	// v から行ける各頂点 next_v について
	for _, next_v := range graph[v] {
		if seen[next_v] {
			continue // next_v が探索済だったらスルー
		}
		dfs(graph, next_v-1, cnt) // 再帰的に探索

		cnt++
		answer[v].end = cnt
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	n := ParseInt(strings.TrimSpace(stdin.Text()))

	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		stdin.Scan()
		s := strings.Split(stdin.Text(), " ")
		graph[i] = make([]int, len(s)-1)
		for j, t := range s {
			if j == 0 {
				continue
			}
			graph[i][j-1] = ParseInt(t)
		}
	}

	cnt := 0
	seen = make([]bool, n)
	answer = make([]Answer, n)
	dfs(graph, 0, cnt)
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
