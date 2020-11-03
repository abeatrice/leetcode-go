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
