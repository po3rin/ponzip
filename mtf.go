package main

import "bytes"

func MTF(t string) ([]byte, string) {
	var symbols string
	seq := make([]byte, len(t))
	pad := []byte(symbols)

	for i, c := range []byte(t) {
		x := bytes.IndexByte(pad, c)
		seq[i] = byte(x)
		copy(pad[1:], pad[:x])
		pad[0] = c
	}
	return seq, symbols
}
