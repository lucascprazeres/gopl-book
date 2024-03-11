package main

import (
	"os"
	"strings"
)

func LoopEcho() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func JoinEcho() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
}
