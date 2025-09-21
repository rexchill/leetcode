package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	if m == 0 {
		for idx := n - 1; idx >= 0; idx-- {
			nums1[idx] = nums2[idx]
		}
		return
	}
	if n == 0 {
		return
	}

	idx := len(nums1) - 1
	i := len(nums1) - n - 1
	j := n - 1
	for j >= 0 {
		if i == -1 {
			break
		}
		if nums1[i] > nums2[j] {
			nums1[idx] = nums1[i]
			nums1[i] = 0
			idx--
			i--
		} else {
			nums1[idx] = nums2[j]
			idx--
			j--
		}
	}

	for ; j >= 0; j-- {
		nums1[j] = nums2[j]
	}
}
func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
