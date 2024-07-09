package main

import (
	"fmt"

	merge_string "github.com/simoncra/leetcode75/1768"
)

func main() {
	stringMerged := merge_string.MergeAlternately("ab", "pqrs")

	fmt.Println("merge_string", stringMerged)
}
