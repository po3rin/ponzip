package main

import (
	"sort"
	"strings"
	"unicode/utf8"
)

func BWT(t string) string {
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
func BWTInverse(t string) string {

	// Cの構築 ------------
	C := make(map[rune]int)
	for _, c := range t {
		v, ok := C[c]
		if !ok {
			C[c] = 1
			continue
		}
		C[c] = v + 1
	}

	sig := Deduplicate(strings.Split(t, ""))

	var sum int
	for _, c := range sig {
		r := []rune(c)[0]
		cur := C[r]
		C[r] = sum
		sum = sum + cur
	}

	// LF-mapping ----------

	lf := make([]int, utf8.RuneCountInString(t))
	for i, c := range t {
		lf[C[c]] = i
		C[c] = C[c] + 1
	}

	// inverse -------------
	r := []rune(t)
	p := strings.Index(t, "$")
	result := make([]rune, len(r))
	for i := range t {
		p = lf[p]
		result[i] = r[p]
	}

	// 最後の$を外して返す
	return string(result[:len(result)-1])
}
