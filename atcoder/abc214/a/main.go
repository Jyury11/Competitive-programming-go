package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := stdin.Text()
	strSlice := strings.Split(s, " ")
	intSlice := []int{}
	for _, i := range strSlice {
		num, _ := strconv.Atoi(i)
		intSlice = append(intSlice, num)
	}

	if intSlice[0] < 126 {
		fmt.Println(4)
	}

	if 126 <= intSlice[0] && intSlice[0] < 212 {
		fmt.Println(6)
	}

	if 212 <= intSlice[0] && intSlice[0] < 215 {
		fmt.Println(8)
	}
}
