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
}
