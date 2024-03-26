package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(countDifferentBits(c2, c1))
}

func countDifferentBits(a1, a2 [32]uint8) int {
	count := 0
	for i := range a1 {
		if a1[i] != a2[i] {
			count++
		}
	}
	return count
}
