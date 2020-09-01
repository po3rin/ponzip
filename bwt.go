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
	sa := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		sa[i] = s[i:]
	}
	sort.Strings(sa)

	var result string
	for _, v := range sa {
		if len(v) < len(s) {
			result += getChar(s, -len(v)-1)
			continue
		}
		result += getChar(s, -1)
	}
	return result
}

func rank(s string, char string, index int) int {
	r := []rune(s)
	sr := r[:index]

	c := strings.Count(string(sr), char)
	return c
}

func rankLessThan(s string, char string, index int) int {
	r := []rune(s)
	sr := r[:index]

	base := []rune(char)[0]
	var counter int
	for _, c := range sr {
		if c < base {
			counter++
		}
	}
	return counter
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
