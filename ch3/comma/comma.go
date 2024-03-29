// comma insere vírgulas em uma string que representa um inteiro decimal
// não negativo
package main

import "fmt"

func main() {
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
