package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const char = "+"

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := strings.TrimSpace(stdin.Text())

	l := len(s)

	for i := 0; i < 3; i++ {
		leng := l + 1
		if i == 1 {
			leng = 1
		}
		for j := 0; j < leng; j++ {
			fmt.Print(char)
		}
		if i == 1 {
			fmt.Print(s)
		}
		fmt.Println(char)
	}
}
