package main

func twoSum(nums []int, target int) []int {
	var ret []int

	for idx, v := range nums {
		for idx2, v2 := range nums {
			if idx == idx2 {
				continue // can't use the same idx
			}
			if v+v2 == target {
				ret = append(ret, idx, idx2)
				return ret
			}
		}

	}

	return ret
}
