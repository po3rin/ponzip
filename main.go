package main

import (
	"fmt"
)

func main() {
	b := Bwt("abracadabra")
	fmt.Println(b)
	s := BwtInverse(b)
	fmt.Println(s)
}
