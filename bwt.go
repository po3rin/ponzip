package main

import (
	"fmt"
	"sort"
	"strings"
)

func Bwt(t string) string {
	t += "$"
	sa := make([]string, len(t))
	for i := 0; i < len(t); i++ {
		sa[i] = t[i:]
	}
	sort.Strings(sa)

	var result string
	for _, v := range sa {
		if len(v) < len(t) {
			result += GetChar(t, -len(v)-1)
			continue
		}
		result += GetChar(t, -1)
	}
	return result
}

// func BwtIncrement(t string) string {
// 	var (
// 		bwt   string
// 		index int
// 	)

// 	t += "$"
// 	r := []rune(t)

// 	for i := 0; i < len(r); i++ {
// 		c := r[len(r)-i-1]
// 		bwt = Insert(bwt, string(c), index)
// 		index = Rank(bwt, string(c), index) + RankLessThan(bwt, string(c), len([]rune(bwt)))
// 		fmt.Println("------")
// 		fmt.Println(string(c))
// 		fmt.Println(index)
// 	}
// 	return bwt
// }

func BwtInverse(t string) string {
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
