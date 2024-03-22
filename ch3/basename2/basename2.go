package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 se "/" não for encontrado
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
