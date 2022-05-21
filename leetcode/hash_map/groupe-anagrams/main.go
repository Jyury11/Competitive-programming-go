package main

import (
	"sort"
	"strconv"
)

type hashMap struct {
	char rune
	num  int
}

func groupAnagrams(strs []string) [][]string {
	hashTable := map[string][]string{}
	for _, str := range strs {
		if _, exists := hashTable[hash(str)]; !exists {
			hashTable[hash(str)] = []string{str}
			continue
		}
		hashTable[hash(str)] = append(hashTable[hash(str)], str)
	}

	ret := [][]string{}
	for _, h := range hashTable {
		ret = append(ret, h)
	}
	return ret
}

func hash(str string) string {
	m := map[rune]int{}
	for _, char := range str {
		m[char]++
	}

	ret := ""
	maps := make([]hashMap, len(m))
	i := 0
	for key, value := range m {
		maps[i].char = key
		maps[i].num = value
		i++
	}

	sort.Slice(maps, func(i, j int) bool {
		return maps[i].char > maps[j].char
	})

	for _, ma := range maps {
		ret += string(ma.char)
		ret += strconv.Itoa(ma.num)
	}
	return ret
}
