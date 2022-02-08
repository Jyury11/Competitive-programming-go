package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func getAns() {
	stdin := bufio.NewScanner(os.Stdin)
	n := inputNumber(stdin)
	num := 0
	ans := ""
	fmt.Println(n)
	for i := 0; i < 120; i++ {
		if num < int(n[0]) {
			if num >= 2 && num*2 <= int(n[0]) {
				num *= 2
				ans += "B"
			} else {
				num++
				ans += "A"
			}
		}

		if num == int(n[0]) {
			break
		}
	}
	fmt.Println(ans)
}
