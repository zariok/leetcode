package main

func twoSum2(nums []int, target int) []int {
	numsDone := make(map[int]int, len(nums))

	for idx, val := range nums {
		// what number do we need to make target
		diff := target - val

		// have we the number needed before?
		if vIdx, ok := numsDone[diff]; ok {
			return []int{idx, vIdx}
		}
		// save current number/idx for next loop
		numsDone[val] = idx
	}
	return nil
}
