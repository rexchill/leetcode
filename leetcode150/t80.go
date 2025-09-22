package main

import "fmt"

// T80:删除有序数组中的重复项II
func removeDuplicates2(nums []int) int {
	left := 0
	for i := 1; i < len(nums); i++ {
		if left > 0 && nums[i] == nums[left] && nums[left] == nums[left-1] {
			continue
		} else {
			left++
			nums[left] = nums[i]

		}
	}
	return left + 1
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k := removeDuplicates2(nums)
	fmt.Println(k, nums[:k])

}
