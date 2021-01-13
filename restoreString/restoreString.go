package restoreString

// RestoreString ...
// https://leetcode.com/problems/shuffle-string/
func RestoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for i, pos := range indices {
		result[pos] = s[i]
	}
	return string(result)
}
