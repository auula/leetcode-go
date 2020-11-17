package sort

// SelectionSort 假定一个下标为最小的值所在的位置下标 然后循环比较 然后交换位置
func SelectionSort(nums []int) {
	var min int
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}
