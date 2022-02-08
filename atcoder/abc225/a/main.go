package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ans = []int{
	6,
	3,
	1,
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := strings.TrimSpace(stdin.Text())

	sameCnt := 0
	if s[0] == s[1] {
		sameCnt++
	}

	if s[0] == s[2] {
		sameCnt++
	}

	if sameCnt == 0 && s[1] == s[2] {
		sameCnt++
	}

	fmt.Println(ans[sameCnt])

}
