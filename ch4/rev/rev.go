// reverse inverte uma fatia de ints in-place
package main

import "fmt"

func main() {
	s := [...]int{1, 2, 3, 4, 5}
	reverse(&s)
}

func reverse(ptr *[5]int) {
	for i, j := 0, len(*ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
	fmt.Println(*ptr)
}

func reverseString(s *[]byte) {

}
