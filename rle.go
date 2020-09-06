package main

import (
	"strconv"
)

func RLE(s string) string {
	var result []rune
	var count int
	for i, c := range s {
		if i == 0 {
			result = append(result, c)
			count++
			continue
		}

		if result[len(result)-1] == c {
			count++
			continue
		}

		s = strconv.Itoa(count)
		runeCount := []rune(s)
		result = append(result, runeCount...)
		result = append(result, c)
		count = 1
	}
	s = strconv.Itoa(count)
	runeCount := []rune(s)
	result = append(result, runeCount...)

	return string(result)
}
