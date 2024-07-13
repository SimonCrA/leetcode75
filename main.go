package main

import (
	"fmt"

	gcd_strings "github.com/simoncra/leetcode75/1071"
)

func main() {
	// stringMerged := merge_string.MergeAlternately("ab", "pqrs")
	gcd := gcd_strings.GcdOfStrings("ABCABC", "ABC")

	// fmt.Println("merge_string", stringMerged)
	fmt.Println("gcd_strings", gcd)
}
