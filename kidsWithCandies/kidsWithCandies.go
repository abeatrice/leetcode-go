// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
// For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them.
// Notice that multiple kids can have the greatest number of candies.
package kidsWithCandies

func KidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, candie := range candies {
		if candie > max {
			max = candie
		}
	}

	var results []bool
	for _, candie := range candies {
		if (candie + extraCandies) >= max {
			results = append(results, true)
		} else {
			results = append(results, false)
		}
	}

	return results
}
