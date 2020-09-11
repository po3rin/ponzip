package main

import (
	"fmt"
	// ...
)

func main() {
	t := "abracadabra"
	bwt := BWT(t)    // 後ほど実装
	fmt.Println(bwt) // ard$rcaaaabb

	bwtinv := BWTInverse(bwt) // 後ほど実装
	fmt.Println(bwtinv)       // abracadabra
}
