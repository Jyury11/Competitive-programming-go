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

func inputNumber(stdin *bufio.Scanner) []int {
	strSlice := input(stdin)
	intSlice := []int{}

	for _, i := range strSlice {
		num, _ := strconv.ParseInt(i, 10, 64)
		intSlice = append(intSlice, int(num))
	}
	return intSlice
}

func getAns() {
	stdin := bufio.NewScanner(os.Stdin)
	n := inputNumber(stdin)
	first := inputNumber(stdin)

	check := inputNumber(stdin)

	search := 0
	if check[0] == first[0] {
		if check[1] != first[1] {
			search = check[0]
		}
	}

	if check[0] == first[1] {
		if check[1] != first[0] {
			search = check[0]
		}
	}

	if check[1] == first[0] {
		if check[0] != first[1] {
			search = check[1]
		}
	}

	if check[1] == first[1] {
		if check[0] != first[0] {
			search = check[1]
		}
	}

	ans := "Yes"

	if search == 0 {
		ans = "No"
	}

	for i := 0; i < int(n[0]-3); i++ {
		in := inputNumber(stdin)
		if in[0] != search && in[1] != search {
			ans = "No"
		}
	}

	fmt.Println(ans)
	return

}
