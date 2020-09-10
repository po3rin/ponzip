package main

import (
	"fmt"
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
	r := []rune(t)

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

	for r, c := range C {
		fmt.Printf("%+v : %+v\n", string(r), c)
	}
	// LF-mapping ----------

	lfm := make([]int, len(r))
	for i, c := range t {
		lfm[C[c]] = i
		C[c] = C[c] + 1
	}

	// inverse -------------

	p := strings.Index(t, "$")
	result := make([]rune, len(r))
	for i := range t {
		p = lfm[p]
		result[i] = r[p]
	}

	return string(result[:len(result)-1])
}
