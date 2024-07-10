package merge_string

import (
	"fmt"
	s "strings"
)

func MergeAlternately(word1 string, word2 string) string {
	var word1Ptr *string
	var word2Ptr *string

	word1Ptr = &word1
	word2Ptr = &word2

	var output s.Builder

	i := 0
	j := 0
	for i < len(*word1Ptr) || j < len(*word2Ptr) {
		if i < len(*word1Ptr) {
			output.WriteByte((*word1Ptr)[i])
			i++
		}

		if j < len(*word2Ptr) {
			output.WriteByte((*word2Ptr)[j])
			j++
		}
		fmt.Println(output.String())
	}

	return output.String()
}
