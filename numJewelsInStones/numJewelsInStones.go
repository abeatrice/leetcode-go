// https://leetcode.com/problems/jewels-and-stones/
// You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
// The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".
package numJewelsInStones

func NumJewelsInStones(J, S string) int {
	var result int
	for _, jewel := range J {
		for _, stone := range S {
			if stone == jewel {
				result++
			}
		}
	}
	return result
}
