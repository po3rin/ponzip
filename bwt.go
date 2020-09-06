package main

import (
	"sort"
	"strings"
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

	psi := make(map[int]int)
	for i, c := range t {
		psi[C[c]] = i
		C[c] = C[c] + 1
	}

	// inverse -------------

	var p int
	r := []rune(t)
	result := make([]rune, len(r))
	for i := range t {
		p = psi[p]
		result[i] = r[p]
	}

	return string(result)
}
