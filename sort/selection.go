package sort

// SelectionSort !
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i

		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}

		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}
