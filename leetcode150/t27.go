package main

import "fmt"

// T27:移除元素
func removeElement(nums []int, val int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val), nums)
}
