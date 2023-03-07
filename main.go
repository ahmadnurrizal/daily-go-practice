package main

import (
	"fmt"

	"github.com/ahmadnurrizal/daily-go-practice/functions"
)

func main() {
	fmt.Println(functions.LengthOfLongestSubstring("wadidaww"))
	fmt.Println(functions.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(functions.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	fmt.Println(functions.LengthOfLastWord("Hello World     "))
	fmt.Println(functions.PlusOne([]int{4, 3, 2, 1}))
	fmt.Println(functions.AddBinary("101", "110"))
	fmt.Println(functions.ClimbStairs(5))
	fmt.Println(functions.Merge([]int{1, 3, 0}, 2, []int{2}, 1))
	binaryTree := functions.InitBinaryTree()
	fmt.Println(functions.InorderTraversal(binaryTree))
	fmt.Println(functions.FindMedianSortedArrays([]int{1, 2, 3, 4}, []int{2, 3, 3}))
	fmt.Println(functions.LongestPalindrome("babad"))
	fmt.Println(functions.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(functions.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))

}
