package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	typStart int = iota
	typEnd
	typKabe
	typMiti
)

type Block struct {
	me       []int
	nexts    []int
	typ      int
	isSearch bool
	min      int
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	field := make([]int, 2)
	stdin.Scan()
	for i, t := range strings.Split(stdin.Text(), " ") {
		field[i] = ParseInt(t)
	}

	start := []int{0, 0}
	end := []int{0, 0}
	fields := make([][]*Block, field[1])
	for i := 0; i < field[1]; i++ {
		fields[i] = make([]*Block, field[0])
		stdin.Scan()
		row := strings.Split(stdin.Text(), " ")
		for j, rec := range row {
			if rec == "0" {
				fields[i][j] = &Block{
					me:  []int{i, j},
					typ: typMiti,
					min: 10000000,
				}
				continue
			}

			if rec == "1" {
				fields[i][j] = &Block{
					me:  []int{i, j},
					typ: typKabe,
					min: 10000000,
				}
				continue
			}

			if rec == "s" {
				start = []int{i, j}
				fields[i][j] = &Block{
					me:  []int{i, j},
					typ: typStart,
					min: 10000000,
				}
				continue
			}

			if rec == "g" {
				end = []int{i, j}
				fields[i][j] = &Block{
					me:  []int{i, j},
					typ: typEnd,
					min: 10000000,
				}
			}
		}
	}

	defaultX := 1
	if end[1] <= start[1] {
		defaultX = -1
	}

	defaultY := 1
	if end[0] <= start[0] {
		defaultY = -1
	}

	ans := search(&defaultX, &defaultY, &start, &end, &fields)
	if 1000000 <= ans {
		fmt.Println("Fail")
	} else {
		fmt.Println(ans)
	}
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

func search(defaultX, defaultY *int, start, end *[]int, fields *[][]*Block) int {
	cersolX := (*start)[1]
	cersolY := (*start)[0]
	min := 10000000

	cersolX += (*defaultX)
	if 0 <= cersolX && cersolX != len((*fields)[0]) {
		switch (*fields)[cersolY][cersolX].typ {
		case typEnd:
			(*fields)[cersolY][cersolX].isSearch = true
			(*fields)[cersolY][cersolX].min = 1
			return 1
		case typMiti:
			if !(*fields)[cersolY][cersolX].isSearch {
				(*fields)[cersolY][cersolX].isSearch = true
				ret := search(defaultX, defaultY, &[]int{cersolY, cersolX}, end, fields)
				if ret <= min {
					min = ret + 1
				}
				(*fields)[cersolY][cersolX].min = ret + 1
			} else {
				if (*fields)[cersolY][cersolX].min <= min {
					min = (*fields)[cersolY][cersolX].min + 1
				}
			}
		}
	}

	cersolX -= (*defaultX)
	cersolY += (*defaultY)
	if 0 <= cersolY && cersolY != len((*fields)) {
		switch (*fields)[cersolY][cersolX].typ {
		case typEnd:
			(*fields)[cersolY][cersolX].isSearch = true
			(*fields)[cersolY][cersolX].min = 1
			return 1
		case typMiti:
			if !(*fields)[cersolY][cersolX].isSearch {
				(*fields)[cersolY][cersolX].isSearch = true
				ret := search(defaultX, defaultY, &[]int{cersolY, cersolX}, end, fields)
				if ret <= min {
					min = ret + 1
				}
				(*fields)[cersolY][cersolX].min = ret + 1
			} else {
				if (*fields)[cersolY][cersolX].min <= min {
					min = (*fields)[cersolY][cersolX].min + 1
				}
			}
		}
	}

	cersolX -= (*defaultX)
	cersolY -= (*defaultY)
	if 0 <= cersolX && cersolX != len((*fields)[0]) {
		switch (*fields)[cersolY][cersolX].typ {
		case typEnd:
			(*fields)[cersolY][cersolX].isSearch = true
			(*fields)[cersolY][cersolX].min = 1
			return 1
		case typMiti:
			if !(*fields)[cersolY][cersolX].isSearch {
				(*fields)[cersolY][cersolX].isSearch = true
				ret := search(defaultX, defaultY, &[]int{cersolY, cersolX}, end, fields)
				if ret <= min {
					min = ret + 1
				}
				(*fields)[cersolY][cersolX].min = ret + 1
			} else {
				if (*fields)[cersolY][cersolX].min <= min {
					min = (*fields)[cersolY][cersolX].min + 1
				}
			}
		}
	}

	cersolX += (*defaultX)
	cersolY -= (*defaultY)
	if 0 <= cersolY && cersolY != len((*fields)) {
		switch (*fields)[cersolY][cersolX].typ {
		case typEnd:
			(*fields)[cersolY][cersolX].isSearch = true
			(*fields)[cersolY][cersolX].min = 1
			return 1
		case typMiti:
			if !(*fields)[cersolY][cersolX].isSearch {
				(*fields)[cersolY][cersolX].isSearch = true
				ret := search(defaultX, defaultY, &[]int{cersolY, cersolX}, end, fields)
				if ret <= min {
					min = ret + 1
				}
				(*fields)[cersolY][cersolX].min = ret + 1
			} else {
				if (*fields)[cersolY][cersolX].min <= min {
					min = (*fields)[cersolY][cersolX].min + 1
				}
			}
		}
	}
	return min
}
