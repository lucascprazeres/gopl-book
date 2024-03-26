package main

import (
	"fmt"
)

func main() {
	l := [...]int{1, 2, 3, 4, 5}
	r := [...]int{1, 2, 3, 4, 5}
	rotateLeft(&l, 1)
	rotateRight(&r, 4)
	fmt.Println(l, r)
}

func rotateLeft(s *[5]int, shift int) {
	copy(s[:], append(s[shift:], s[:shift]...))
}

func rotateRight(s *[5]int, shift int) {
	copy(s[:], append(s[len(s)-shift:], s[:len(s)-shift]...))
}
