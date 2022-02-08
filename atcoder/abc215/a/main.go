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
	strSlice := strings.Split(s, ".")
	float, _ := strconv.Atoi(strSlice[1])
	var ret string
	if 0 <= float && float <= 2 {
		ret = "-"
	}
	if 3 <= float && float <= 6 {
		ret = ""
	}
	if 7 <= float && float <= 9 {
		ret = "+"
	}
	fmt.Printf("%s%s\n", strSlice[0], ret)
}
