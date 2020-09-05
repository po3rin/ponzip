package main

import "strings"

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
