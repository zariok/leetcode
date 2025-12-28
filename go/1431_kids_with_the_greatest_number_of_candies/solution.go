package main

func kidsWithCandies(candies []int, extraCandies int) []bool {

	// need to test; if one of the kids gets the extra candies, would it be the greatest
	// of the kids without extras
	candiesPlus := make([]int, len(candies))
	copy(candiesPlus, candies)
	for idx := range candies {
		candiesPlus[idx] += extraCandies
	}

	kids := make([]bool, len(candies))
	for idx, v := range candiesPlus {
		greatest := true
		for _, v2 := range candies {
			if v < v2 {
				greatest = false
				break
			}
		}
		kids[idx] = greatest
	}

	return kids
}
