package functions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

// AddBinary function takes two binary strings a and b as input, adds them together, and returns their sum as a binary string.
func AddBinary(a string, b string) string {
	i, j, res, carry := len(a)-1, len(b)-1, "", 0
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2
	}
	if carry == 1 {
		return "1" + res
	}
	return res
}

func ClimbStairs(n int) int {

	n1, n2 := 1, 1
	var temp int
	for i := 0; i < n-1; i++ {
		temp = n1
		n1 = n1 + n2
		n2 = temp

	}
	return n1
}

// https://leetcode.com/problems/merge-sorted-array/
// I add return on this code, just little adjustment
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	n_num1 := m - 1
	n_num2 := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if (n_num1 >= 0 && n_num2 >= 0 && nums1[n_num1] > nums2[n_num2]) || (n_num2 < 0) {
			nums1[k] = nums1[n_num1]
			n_num1--
		} else {
			nums1[k] = nums2[n_num2]
			n_num2--
		}
	}

	return nums1
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//// Option 1
	// n_num1 := len(nums1) - 1
	// n_num2 := len(nums2) - 1

	// myArr := make([]int, len(nums1)+len(nums2))
	// copy(myArr[:], nums1[:])
	// fmt.Println(myArr)

	// for k := len(nums1) + len(nums2) - 1; k >= 0; k-- {
	// 	if (n_num1 >= 0 && n_num2 >= 0 && myArr[n_num1] > nums2[n_num2]) || (n_num2 < 0) {
	// 		myArr[k] = myArr[n_num1]
	// 		n_num1--
	// 	} else {
	// 		myArr[k] = nums2[n_num2]
	// 		n_num2--
	// 	}
	// }

	// mid := (len(nums1) + len(nums2)) / 2
	// if len(myArr)%2 == 0 {
	// 	return float64(myArr[mid-1]+myArr[mid]) / 2.0
	// } else {
	// 	return float64(myArr[mid])
	// }

	//// Option 2
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	return float64(nums1[len(nums1)>>1]+nums1[(len(nums1)-1)>>1]) / 2.0
}

// https://leetcode.com/problems/longest-palindromic-substring/description/
func LongestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		max = MaxPalindrome(s, i, i, max)   // check for odd palindrome
		max = MaxPalindrome(s, i, i+1, max) // check for even palindrome
	}
	return max
}

func MaxPalindrome(s string, i, j int, max string) string {
	var sub string
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(max) < len(sub) { // compare lasts max palindrome
		return sub
	}
	return max
}

// https://leetcode.com/problems/container-with-most-water/description/
func MaxArea(height []int) int {
	var max, tempMax int
	i, j := 0, len(height)-1
	for i < j {
		if height[i] > height[j] {
			tempMax = height[j] * (j - i)
			j--
		} else {
			tempMax = height[i] * (j - i)
			i++
		}
		if tempMax > max {
			max = tempMax
		}
	}
	return max
}

// https://leetcode.com/problems/3sum/description/
func ThreeSum(nums []int) [][]int {
	n := len(nums)

	// Sort the given array
	sort.Ints(nums)

	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// Skip all duplicates from left
		// num1Idx>0 ensures this check is made only from 2nd element onwards
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				num3Idx--
			} else {
				// Increment num2Idx to increase sum value
				num2Idx++
			}
		}
	}
	return result
}
