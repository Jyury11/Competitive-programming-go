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
	in := stdin.Text()
	strSlice := strings.Split(in, " ")
	intSlice := []int{}
	for _, i := range strSlice {
		num, _ := strconv.Atoi(i)
		intSlice = append(intSlice, num)
	}

	s := intSlice[0]
	t := intSlice[1]
	max := s
	cnt := 0

	for a := 0; a <= max; a++ {
		for b := 0; b <= max; b++ {
			for c := 0; c <= max; c++ {
				if a+b+c <= s && a*b*c <= t {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
