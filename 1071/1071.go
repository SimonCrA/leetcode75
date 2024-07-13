package gcd_strings

//   - Since the candidate string needs to evenly divide str1 and str2 we only need to consider candidate strings
//     with a length that is evenly divisible into the length of both str1 and str2. This quickly provides us with
//     a solution that significantly outperforms brute-force.
//   - As a result of the above, we only need to try string lengths up to the length of the shorter of our two input strings.
//     This gives us our bounds for determining candidate string lengths.
//   - Since we are looking for the longest candidate we can begin our iteration from the largest string length that
//     evenly divides both strings and work downward to smaller length strings. The first string that fits our criteria
//     is then our answer.
func GcdOfStrings(str1 string, str2 string) string {
	return ""
}
