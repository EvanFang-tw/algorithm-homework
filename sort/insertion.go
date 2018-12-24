package sort

// InsertionSort !
func InsertionSort(nums []int) {

	for i := 1; i < len(nums); i++ {

		j := i

		for j > 0 {

			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			} else {
				break
			}

			j = j - 1
		}
	}
}
