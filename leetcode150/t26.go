package main

// T26:删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	left := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[left] {
			left++
			nums[left] = nums[i]
		}
	}
	return left + 1
}

//func main() {
//	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
//	fmt.Println(removeDuplicates(nums), nums)
//
//}
