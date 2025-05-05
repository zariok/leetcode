2645. Minimum Additions to Make Valid String

Given a string word to which you can insert letters "a", "b" or "c" anywhere and any number of times, return the minimum number of letters that must be inserted so that word becomes valid.

A string is called valid if it can be formed by concatenating the string "abc" several times.

### Example 1
```
Input: word = "b"
Output: 2
Explanation: Insert the letter "a" right before "b", and the letter "c" right next to "b" to obtain the valid string "abc".
```


### Example 2
```
Input: word = "b"
Output: 2
Explanation: Insert the letter "a" right before "b", and the letter "c" right next to "b" to obtain the valid string "abc".
```


### Example 3
```
Input: word = "b"
Output: 2
Explanation: Insert the letter "a" right before "b", and the letter "c" right next to "b" to obtain the valid string "abc".
```


## My work
Initially I overthought this problem and figured any configured string could be the "valid word".  
**Solution 1** is overly complex with looking at the byte slices, ensuring I don't overrun the slice and check the next letter.  I was also building out the correct word.
**Solution 2** iterates over the word as runes and only calculates the changes