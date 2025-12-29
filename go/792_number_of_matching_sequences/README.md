# 792. Number of Matching Subsequences


Given a string `s` and an array of strings `words`, return the number of `words[i]` that is a subsequence of `s`.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
 
## Example 1:
```
Input: s = "abcde", words = ["a","bb","acd","ace"]
Output: 3
Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".
```

## Example 2:
```
Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
Output: 2
``` 

## Constraints:
```
1 <= s.length <= 5 * 104
1 <= words.length <= 5000
1 <= words[i].length <= 50
s and words[i] consist of only lowercase English letters.
```

## My work
### Solution 1:
This is similar to 392 but works over an array of words

### Solution 2: 
After seeing my submission be in the 1500ms response, I went to see how to make it more
efficient.  If we keep track of the words and what position letter we are at, we can 
iterate over the `s` and work any word waiting for that letter.  Now for each letter
in `s` we work any `word` waiting for that letter and increment it to the next letter
and store it in the "waiting for".

