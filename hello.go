package main

import (
	"fmt"
	"hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	reversed := "nimij"
	original := morestrings.ReverseRunes(reversed)
	fmt.Printf("Original: %v, Reversed: %v \n", reversed, original)
	fmt.Println(cmp.Diff("nimij", "jimin"))
}
