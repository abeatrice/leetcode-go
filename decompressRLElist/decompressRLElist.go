// https://leetcode.com/problems/decompress-run-length-encoded-list/
// We are given a list nums of integers representing a list compressed with run-length encoding.
// Consider each adjacent pair of elements [freq, val] = [nums[2*i], nums[2*i+1]] (with i >= 0).  For each such pair, there are freq elements with value val concatenated in a sublist. Concatenate all the sublists from left to right to generate the decompressed list.
// Return the decompressed list.
package decompressRLElist

func DecompressRLElist(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i += 2 {
		cur := nums[i : i+2]
		for j := 0; j < cur[0]; j++ {
			result = append(result, cur[1])
		}
	}
	return result
}
