package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type kInfo struct {
	index int
	a     int
}

func main() {
	getAns()
}

func input(stdin *bufio.Scanner) []string {
	stdin.Scan()
	s := stdin.Text()
	return strings.Split(s, " ")
}

func inputNumber(stdin *bufio.Scanner) []uint64 {
	strSlice := input(stdin)
	intSlice := []uint64{}
	for _, i := range strSlice {
		num, _ := strconv.ParseUint(i, 10, 64)
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func makeKInfo(ka [][]uint64) []kInfo {
	ks := []kInfo{}
	for i, k := range ka {
		if len(k) > 0 {
			ks = append(ks, kInfo{
				index: i + 1,
				a:     int(k[0]),
			})
		}
	}
	return ks
}

func checkKsStart(ks []kInfo) ([]int, bool) {
	same := make(map[int]int)
	for _, k := range ks {
		if same[k.a] == 0 {
			same[k.a] = k.index
		} else {
			return []int{same[k.a], k.index}, true
		}
	}
	return []int{}, false
}

func getAns() {
	stdin := bufio.NewScanner(os.Stdin)
	nm := inputNumber(stdin)
	m := int(nm[1])
	ka := [][]uint64{}
	for i := 0; i < m; i++ {
		_ = int(inputNumber(stdin)[0])
		a := inputNumber(stdin)
		ka = append(ka, a)
	}

	ans := "No"
	ks := makeKInfo(ka)

	for {
		if len(ks) == 0 {
			ans = "Yes"
			break
		}
		kSames, ok := checkKsStart(ks)
		if ok {
			for _, kSame := range kSames {
				ka[kSame-1] = ka[kSame-1][1:len(ka[kSame-1])]
			}
			ks = makeKInfo(ka)
		} else {
			break
		}
	}
	fmt.Println(ans)
}
