package main

import (
	"sort"
	"strings"
)

func GetChar(t string, i int) string {
	rs := []rune(t)
	if i < 0 {
		return string(rs[len(rs)+i])
	}
	return string(rs[i])
}

func Insert(t string, c string, index int) string {
	r := []rune(t)
	r = append(r, 0)
	copy(r[index+1:], r[index:])
	r[index] = []rune(c)[0]
	return string(r)
}

func Rank(t string, char string, index int) int {
	r := []rune(t)
	sr := r[:index]

	c := strings.Count(string(sr), char)
	return c
}

func RankLessThan(t string, char string, index int) int {
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

func Deduplicate(s []string) []string {
	var result []string
	sort.Strings(s)
	j := 0
	for i := 1; i < len(s); i++ {
		if s[j] == s[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		s[j] = s[i]
	}
	result = s[:j+1]
	return result
}

func MoveToFront(needle rune, haystack []rune) []rune {
	if len(haystack) == 0 || haystack[0] == needle {
		return haystack
	}
	var prev rune
	for i, elem := range haystack {
		switch {
		case i == 0:
			haystack[0] = needle
			prev = elem
		case elem == needle:
			haystack[i] = prev
			return haystack
		default:
			haystack[i] = prev
			prev = elem
		}
	}
	return append(haystack, prev)
}
