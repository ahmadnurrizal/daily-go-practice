package functions

import "strings"

func LengthOfLongestSubstring(s string) int {
	n := len(s)
	seen := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < n; end++ {
		if idx, ok := seen[s[end]]; ok && idx >= start {
			start = idx + 1
		}
		seen[s[end]] = end
		maxLen = Max(maxLen, end-start+1)
	}

	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// removeDuplicates function is to modify a given slice of integers called nums in place to have only unique values, by removing any duplicates that exist within the slice. The function returns the number of unique values in the modified slice,
func RemoveDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}

	j := 0 // points to  the index of last filled position
	for i := 1; i < ln; i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

// remove multiple value on slice
func RemoveElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			copy(nums[i:], nums[i+1:]) // remove element but duplicate on last element
			nums = nums[:len(nums)-1]  // remove last element
			i--
		}
	}
	return len(nums)
}

// The LengthOfLastWord function takes a string as input and returns the length of the last word in the string.
func LengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}

// The plusOne function takes an input slice of non-negative integers representing a non-negative integer and returns a new slice representing that integer plus one.
func PlusOne(digits []int) []int {
	var n int = len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	var a = make([]int, n+1)
	a[0] = 1
	return a
}
