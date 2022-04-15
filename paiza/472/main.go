package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Member struct {
	ninki int
	id    int
	tomos map[int]bool
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	member := make([]int, 2)
	stdin.Scan()
	for i, t := range strings.Split(stdin.Text(), " ") {
		member[i] = ParseInt(t)
	}

	ninkis := make([]*Member, member[0])
	members := make([]*Member, member[0])
	stdin.Scan()
	for i, ni := range strings.Split(stdin.Text(), " ") {
		mem := &Member{
			ninki: ParseInt(ni),
			id:    i,
			tomos: make(map[int]bool, member[0]),
		}
		ninkis[i] = mem
		members[i] = mem
	}

	sort.Slice(ninkis, func(i, j int) bool { return ninkis[i].ninki > ninkis[j].ninki })

	for i := 0; i < member[1]; i++ {
		stdin.Scan()
		tomo := strings.Split(stdin.Text(), " ")
		t1 := ParseInt(tomo[0]) - 1
		t2 := ParseInt(tomo[1]) - 1
		members[t1].tomos[t2] = true
		members[t2].tomos[t1] = true
	}

	for i := 0; i < member[0]; i++ {
		osusume := -1
		for _, n := range ninkis {
			if n.id == i {
				continue
			}
			if ok, ok2 := n.tomos[i]; ok && ok2 {
				continue
			}

			osusume = n.id + 1
			break
		}
		fmt.Println(osusume)
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
