package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	startChar    = '('
	endChar      = ')'
	maxStrLength = 60
	alphaLength  = 26
)

var (
	ret = map[rune][]int{
		'a': make([]int, maxStrLength),
		'b': make([]int, maxStrLength),
		'c': make([]int, maxStrLength),
		'd': make([]int, maxStrLength),
		'e': make([]int, maxStrLength),
		'f': make([]int, maxStrLength),
		'g': make([]int, maxStrLength),
		'h': make([]int, maxStrLength),
		'i': make([]int, maxStrLength),
		'j': make([]int, maxStrLength),
		'k': make([]int, maxStrLength),
		'l': make([]int, maxStrLength),
		'm': make([]int, maxStrLength),
		'n': make([]int, maxStrLength),
		'o': make([]int, maxStrLength),
		'p': make([]int, maxStrLength),
		'q': make([]int, maxStrLength),
		'r': make([]int, maxStrLength),
		's': make([]int, maxStrLength),
		't': make([]int, maxStrLength),
		'u': make([]int, maxStrLength),
		'v': make([]int, maxStrLength),
		'w': make([]int, maxStrLength),
		'x': make([]int, maxStrLength),
		'y': make([]int, maxStrLength),
		'z': make([]int, maxStrLength),
	}

	retCnt = map[rune][]int{
		'a': make([]int, 0, maxStrLength),
		'b': make([]int, 0, maxStrLength),
		'c': make([]int, 0, maxStrLength),
		'd': make([]int, 0, maxStrLength),
		'e': make([]int, 0, maxStrLength),
		'f': make([]int, 0, maxStrLength),
		'g': make([]int, 0, maxStrLength),
		'h': make([]int, 0, maxStrLength),
		'i': make([]int, 0, maxStrLength),
		'j': make([]int, 0, maxStrLength),
		'k': make([]int, 0, maxStrLength),
		'l': make([]int, 0, maxStrLength),
		'm': make([]int, 0, maxStrLength),
		'n': make([]int, 0, maxStrLength),
		'o': make([]int, 0, maxStrLength),
		'p': make([]int, 0, maxStrLength),
		'q': make([]int, 0, maxStrLength),
		'r': make([]int, 0, maxStrLength),
		's': make([]int, 0, maxStrLength),
		't': make([]int, 0, maxStrLength),
		'u': make([]int, 0, maxStrLength),
		'v': make([]int, 0, maxStrLength),
		'w': make([]int, 0, maxStrLength),
		'x': make([]int, 0, maxStrLength),
		'y': make([]int, 0, maxStrLength),
		'z': make([]int, 0, maxStrLength),
	}

	alphas = [alphaLength]rune{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
	}

	buffer       = make([][]rune, maxStrLength)
	bufferNumber = make([]int, maxStrLength)
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := strings.TrimSpace(stdin.Text())
	Decode(s)
	Answer()
}

func Decode(s string) {
	isStart := false
	n := 1
	cnt := 0

	bufferCnt := 0
	startCnt := 0

	for _, r := range s {
		if r == startChar {
			startCnt++
			if isStart {
				bufferNumber[cnt] = n
				isStart = false
				n = 1
				cnt++
				bufferCnt = 0
			}
			continue
		}

		if r == endChar {
			tmp := map[rune]struct{}{}
			for _, ru := range buffer[cnt] {
				if _, ok := tmp[ru]; ok {
					continue
				}
				tmp[ru] = struct{}{}
				ret[ru][startCnt] *= bufferNumber[cnt-1]
				ret[ru][cnt-1] += ret[ru][cnt]
				ret[ru][cnt] = 0
			}

			buffer[cnt-1] = append(buffer[cnt-1], buffer[cnt]...)
			buffer[cnt] = []rune{}
			startCnt--
			cnt--
			continue
		}

		if '0' <= r && r <= '9' {
			if !isStart {
				n = int(r - '0')
			} else {
				n = n*10 + int(r-'0')
			}
			isStart = true
			continue
		}
		ret[r][startCnt] += n
		retCnt[r] = append(retCnt[r], startCnt)

		if isStart {
			isStart = false
			n = 1
			bufferCnt = 0
		}

		if bufferCnt == 0 && len(buffer[cnt]) == 0 {
			buffer[cnt] = make([]rune, 0, alphaLength)
		}
		buffer[cnt] = append(buffer[cnt], r)
		bufferCnt++
	}
}

func Answer() {
	for _, key := range alphas {
		fmt.Printf("%s %d\n", string(key), ret[key][0])
	}
}
