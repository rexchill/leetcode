package main

// T55:跳跃游戏
func canJump(nums []int) bool {
	right := 0
	for i, v := range nums {
		if i > right {
			break
		}
		right = max(i+v, right)
	}
	return right >= len(nums)-1
}
