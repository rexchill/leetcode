package main

// T169:多数元素
func majorityElement(nums []int) int {
	num := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if num == nums[i] {
			count++
			continue
		}
		count--
		if count == 0 {
			i++
			num = nums[i]
			count = 1
		}
	}
	return num
}

/*
	func majorityElement(nums []int) (ans int) {
	    hp := 0
	    for _, x := range nums {
	        if hp == 0 { // x 是初始擂主，生命值为 1
	            ans, hp = x, 1
	        } else if x == ans { // 比武，同门加血，否则扣血
	            hp++
	        } else {
	            hp--
	        }
	    }
	    return
	}
*/
//func main() {
//	nums := []int{2, 2, 1, 1, 1, 2, 2}
//	fmt.Println(majorityElement(nums))
//}
