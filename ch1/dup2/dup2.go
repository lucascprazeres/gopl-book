// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma
// vez na entrada. Ele lê de stdin ou de uma lista de arquivos nomeados.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for fileName, lineCount := range counts {
		for line, n := range lineCount {
			if n > 1 {
				fmt.Printf("%d\t%s\t%v\n", n, line, fileName)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[fileName] == nil {
			counts[fileName] = make(map[string]int)
		} else {
			counts[fileName][input.Text()]++
		}
	}
	// NOTA: ignorando errors em potencial de input.Err()
}
