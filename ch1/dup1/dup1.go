// Dup1 exibe o texto de toda linha que aparece mais de
// uma vez na entrada-padrão, precedia por sua contagem.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	// NOTA: ignorando erros em potencial de input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
