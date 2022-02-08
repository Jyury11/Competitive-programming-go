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
	in := make([][]uint64, int(n[0]))
	for i := 0; i < int(n[0]); i++ {
		in[i] = inputNumber(stdin)
	}

	var length float64 = 0
	var heb float64 = 0
	var tmpHeb float64 = 0
	for _, val := range in {
		heb += float64(val[0]) / float64(val[1])
	}
	time := heb / float64(2)
	for _, val := range in {
		tmp := float64(val[0]) / float64(val[1])
		if (tmpHeb + tmp) > time {
			t := time - tmpHeb
			if t == 0 {
				break
			}
			// l := float64(val[1]) / t
			fmt.Println(tmpHeb)
			fmt.Println(tmp)
			fmt.Println(t)
			length += t
			break
		}
		tmpHeb += tmp
		length += float64(val[0])

	}

	fmt.Println(length)

}
