package sort

// InsertionSort 把无序数组分成2组 一组有序一组无序 从无序拿数据和有序比较 然后插入到合适位置
func InsertionSort(nums []int) {
	var pervIndex, current int
	for i := 1; i < len(nums); i++ {
		current = nums[i]
		pervIndex = i - 1
		for pervIndex >= 0 && current < nums[pervIndex] {
			nums[pervIndex+1] = nums[pervIndex]
			pervIndex--
		}
		nums[pervIndex+1] = current
	}
}
