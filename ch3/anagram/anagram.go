// anagram recebe duas strings e diz se elas são anagramas.
package main

import (
	"fmt"
	"maps"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("anagram: not enough arguments")
	}

	first, second := os.Args[1], os.Args[2]
	result := isAnagram(first, second)
	fmt.Println(result)
}

func isAnagram(s1, s2 string) bool {
	return maps.Equal(charOccurences(s1), charOccurences(s2))
}

func charOccurences(s string) map[string]int {
	occurences := make(map[string]int)
	for _, r := range s {
		occurences[string(r)]++
	}
	return occurences
}
