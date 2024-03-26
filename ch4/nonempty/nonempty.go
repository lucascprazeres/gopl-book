// Nonempty devolve uma fatia que armazena apenas as strings não vazias.
// O array subjacente é modificado durante a chamada
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonEmpty2(data))
	fmt.Printf("%q\n", data)
}

func nonEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonEmpty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
