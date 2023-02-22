package functions

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
