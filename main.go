package main

import (
	"fmt"
)

func main() {
	b := bwt("abracadabra")
	fmt.Println(b)
	s := invBwt(b)
	fmt.Println(s)
}
