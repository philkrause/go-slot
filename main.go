package main

import (
	"fmt"
)

func main() {

	reelSymbols := Symbol()

	r1 := RNGGen(9)

	fmt.Println(reelSymbols[r1])

}
