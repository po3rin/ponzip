package main

import (
	"fmt"
	"sort"
	"strings"
)

func getChar(s string, i int) string {
	rs := []rune(s)
	if i < 0 {
		return string(rs[len(rs)+i])
	}
	return string(rs[i])
}

func bwt(s string) string {
	s += "$"
	ss := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ss[i] = s[i:]
	}
	sort.Strings(ss)

	var result string
	for _, v := range ss {
		if len(v) < len(s) {
			result += getChar(s, -len(v)-1)
			continue
		}
		result += getChar(s, -1)
	}
	return result
}

func invBwt(s string) string {
	idx := strings.Index(s, "$")
	fmt.Println(idx)
	charCounter := make(map[rune]int)
	// 各文字の出現回数をカウント
	for _, c := range s {
		_, ok := charCounter[c]
		if !ok {
			charCounter[c] = 1
			continue
		}
		charCounter[c]++
	}
	fmt.Println(charCounter)
	return ""
}
