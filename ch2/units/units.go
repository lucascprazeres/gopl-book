package main

import (
	"ch2/lengthconv"
	"ch2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// do something
		return
	}

	for _, arg := range args {
		f := parseArgumentIntoFloat(arg)
		convertUnits(f)
	}
}

func parseArgumentIntoFloat(arg string) float64 {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "units: %v\n", err)
		os.Exit(1)
	}
	return t
}

func convertUnits(arg float64) {
	fmt.Println("==== TEMPERATURAS ====")

	celsius := tempconv.Celsius(arg)
	fahrenheit := tempconv.Fahrenheit(arg)
	kelvin := tempconv.Kelvin(arg)

	fmt.Printf("%s = %s, %s = %s\n", celsius, tempconv.CToF(celsius), fahrenheit, tempconv.FToC(fahrenheit))
	fmt.Printf("%s = %s, %s = %s\n", celsius, tempconv.CToK(celsius), kelvin, tempconv.KToC(kelvin))
	fmt.Printf("%s = %s, %s = %s\n", fahrenheit, tempconv.FToK(fahrenheit), kelvin, tempconv.KToF(kelvin))

	fmt.Println("==== COMPRIMENTO ====")

	meters := lengthconv.Meter(arg)
	feet := lengthconv.Feet(arg)

	fmt.Printf("%s = %s, %s = %s\n", meters, lengthconv.MToF(meters), feet, lengthconv.FToM(feet))
}
