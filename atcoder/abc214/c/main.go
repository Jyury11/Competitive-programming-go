package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := getAns()
	for _, answer := range ans {
		fmt.Println(answer)
	}
}

func input(stdin *bufio.Scanner) []uint64 {
	stdin.Scan()
	s := stdin.Text()
	strSlice := strings.Split(s, " ")
	intSlice := []uint64{}
	for _, i := range strSlice {
		num, _ := strconv.ParseUint(i, 10, 64)
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func getTimes(n, minIndex uint64, t, s []uint64) []uint64 {
	ans := make([]uint64, n)
	var time uint64 = 0
	for i := minIndex; i < n+minIndex; i++ {
		if time == 0 || time > t[i%n] {
			ans[i%n] = t[i%n]
			time = t[i%n]
		} else {
			ans[i%n] = time
		}
		time += s[i%n]
	}
	return ans
}

func getMinIndex(t, s []uint64) uint64 {
	var min uint64 = 0
	var minIndex uint64 = 0
	for index, num := range t {
		if min > num {
			min = num
			minIndex = uint64(index)
		}
	}
	return minIndex
}

func getAns() []uint64 {
	stdin := bufio.NewScanner(os.Stdin)
	n := input(stdin)
	s := input(stdin)
	t := input(stdin)
	minIndex := getMinIndex(t, s)
	return getTimes(n[0], minIndex, t, s)
}
