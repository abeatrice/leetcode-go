// https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Given a string of allowed distinct characters and an array of strings words.
// A string is consistent if all characters in the string appear in the string allowed.
// Returns the number of consistent strings in the array words.
package countConsistentStrings

func CountConsistentStrings(allowed string, words []string) int {
	var result int
	for _, word := range words {
		if wordIsValid(allowed, word) {
			result++
		}
	}
	return result
}

func wordIsValid(allowed, word string) bool {
	var found bool
	for _, letter := range word {
		found = false
		for _, target := range allowed {
			if letter == target {
				found = true
				break
			}
		}
		if found != true {
			return false
		}
	}
	return true
}
