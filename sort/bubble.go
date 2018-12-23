package sort

// BubbleSort !
func BubbleSort(nums []int) {

	length := len(nums)

	for i := 0; i < length-1; i++ {

		isSwap := false

		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSwap = true
			}
		}

		if isSwap == false {
			break
		}
	}
}
