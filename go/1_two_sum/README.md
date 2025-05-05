# 1. Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

**^^ test this?**

### Example 1
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

### Example 2
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

### Example 3
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

### Constraints:
* 2 <= nums.length <= $10^4$
* $-10^9$ <= nums[i] <= $10^9$
* $-10^9$ <= target <= $10^9$
* **Only one valid answer exists**


### Follow-up
Can you come up with an algorithm that is less than $0(n^2)$ time complexity?


### My Implementation

#### Solution 1: 
The quick nested range.

#### Solution 2: 
* Make a map and keep track of all the numbers we saw previously.
* Calculate what value we need to match and check the map
* If we didn't match, add that val to the numbers we've already covered for next round.
