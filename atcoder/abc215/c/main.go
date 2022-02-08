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
	num := int(n[0])
	ans := []string{}
	if int(n[0])%2 != 0 {
		ans = append(ans, "A")
		num--
	}

	for i := 0; i < 120; i++ {
		if num%2 == 0 && num/2 >= 0 {
			ans = append(ans, "B")
			num = num / 2
		} else if num > 0 {
			ans = append(ans, "A")
			num--
		}

		if num == 0 {
			break
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Print(ans[i])
	}
	fmt.Print("\n")
}
