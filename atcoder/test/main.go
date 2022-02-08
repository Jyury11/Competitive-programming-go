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
	fmt.Println(intSlice)
}
