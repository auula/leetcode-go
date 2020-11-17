package sort

// MergeSort 分治法 对无序数组进行拆分 然后排序
func MergeSort(nums []int) []int {
	var result []int
	if len(nums) < 2 {
		return nums
	}
	middle := len(nums) / 2
	left := nums[:middle]
	right := nums[middle:]
	return func(left, right []int) []int {
		for len(left) != 0 && len(right) != 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		}
		if len(left) != 0 {
			result = append(result, left[:]...)
		}
		if len(right) != 0 {
			result = append(result, right[:]...)
		}
		return result
	}(MergeSort(left), MergeSort(right))
}
