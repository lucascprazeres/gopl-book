// Ftoc exibe duas conversões de Fahrenheit para Celsius
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gºF = %gºC\n", freezingF, ftoC(freezingF))
	fmt.Printf("%gºF = %gºC\n", boilingF, ftoC(boilingF))
}

func ftoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
