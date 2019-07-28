package main

import (
	"fmt"

	"./popcountmap"
	"./popcountloop"
)

func main() {
	n := uint64(0x0F0F0F0F0F0F0F0F)
	fmt.Printf("%064b\n", n)
	fmt.Println("Number of set bits:")
	fmt.Printf("[popcountmap]: %d\n", popcountmap.PopCount(n))
	fmt.Printf("[popcountloop]: %d\n", popcountloop.PopCount(n))
}
