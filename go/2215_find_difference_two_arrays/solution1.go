package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	solution := make([][]int, 2)
	solution[0] = []int{}
	solution[1] = []int{}

	// make hashmaps
	h1 := make(map[int]bool)
	for _, num := range nums1 {
		h1[num] = true
	}
	h2 := make(map[int]bool)
	for _, num := range nums2 {
		h2[num] = true
	}

	// iterate over nums1 and check for in nums2
	seen1 := make(map[int]bool)
	for _, num := range nums1 {
		if !h2[num] && !seen1[num] {
			seen1[num] = true
		}
	}
	seen2 := make(map[int]bool)
	for _, num := range nums2 {
		if !h1[num] && !seen2[num] {
			seen2[num] = true
		}
	}

	// convert back to array
	for num := range seen1 {
		solution[0] = append(solution[0], num)
	}
	for num := range seen2 {
		solution[1] = append(solution[1], num)
	}

	return solution
}
