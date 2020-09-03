package main

import (
	"fmt"
	"sort"
	"strings"
)

func getChar(t string, i int) string {
	rs := []rune(t)
	if i < 0 {
		return string(rs[len(rs)+i])
	}
	return string(rs[i])
}

func bwt(t string) string {
	t += "$"
	sa := make([]string, len(t))
	for i := 0; i < len(t); i++ {
		sa[i] = t[i:]
	}
	sort.Strings(sa)

	var result string
	for _, v := range sa {
		if len(v) < len(t) {
			result += getChar(t, -len(v)-1)
			continue
		}
		result += getChar(t, -1)
	}
	return result
}

func rank(t string, char string, index int) int {
	r := []rune(t)
	sr := r[:index]

	c := strings.Count(string(sr), char)
	return c
}

func rankLessThan(t string, char string, index int) int {
	r := []rune(t)
	ra := r[:index]

	base := []rune(char)[0]
	var counter int
	for _, c := range ra {
		if c < base {
			counter++
		}
	}
	return counter
}

func bwtInverse(t string) string {
	idx := strings.Index(t, "$")
	fmt.Println(idx)
	charCounter := make(map[rune]int)

	for _, c := range t {
		_, ok := charCounter[c]
		if !ok {
			charCounter[c] = 1
			continue
		}
		charCounter[c]++
	}
	return ""
}
