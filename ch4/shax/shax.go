package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hash = flag.String("hash", "256", "version of SHA algorithm")

func main() {
	flag.Parse()

	outputStr := readUserInput()

	switch *hash {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(outputStr)))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(outputStr)))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(outputStr)))
	}
}

func readUserInput() string {
	input := bufio.NewScanner(os.Stdin)
	result := ""
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		result += input.Text()
	}
	return result
}
