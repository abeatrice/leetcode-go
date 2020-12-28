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
