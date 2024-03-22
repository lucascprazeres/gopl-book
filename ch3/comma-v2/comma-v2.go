package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	fmt.Println(commaV2(os.Args[1]))
}

func commaV2(s string) string {
	var buff bytes.Buffer
	commaIndexes := make([]int, len(s)/2)

	for i := len(s); i >= 0; i -= 3 {
		commaIndexes = append(commaIndexes, i)
	}

	for i, v := range s {
		if i != 0 && slices.Contains(commaIndexes, i) {
			buff.WriteByte(',')
		}
		buff.WriteRune(v)

	}
	return buff.String()
}
