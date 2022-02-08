package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := stdin.Text()
	minNum := int(s[0])
	minI := 0
	maxNum := int(s[0])
	maxI := 0
	for i, char := range s {
		if minNum > int(char) {
			minNum = int(char)
			minI = i
		} else if minNum == int(char) {
			for j := 1; j < len(s); j++ {
				next1 := s[(minI+j)%len(s)]
				next2 := s[(i+j)%len(s)]
				if int(next1) > int(next2) {
					minNum = int(char)
					minI = i
					break
				} else if int(next1) == int(next2) {
					continue
				}
				break
			}

		}
		if maxNum < int(char) {
			maxNum = int(char)
			maxI = i
		} else if maxNum == int(char) {
			for j := 1; j < len(s); j++ {
				next1 := s[(maxI+j)%len(s)]
				next2 := s[(i+j)%len(s)]
				if int(next1) < int(next2) {
					maxNum = int(char)
					maxI = i
					break
				} else if int(next1) == int(next2) {
					continue
				}
				break
			}
		}
	}

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[(minI+i)%len(s)]))
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[(maxI+i)%len(s)]))
	}
	fmt.Println()

}
