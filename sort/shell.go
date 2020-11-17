package sort

// ShellSort 缩小增量排序 对数据不断插入排序
func ShellSort(nums []int) {
	var pervIndex, current int
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			current = nums[i]
			pervIndex = i - gap
			for pervIndex >= 0 && nums[pervIndex] > current {
				nums[pervIndex+gap] = nums[pervIndex]
				pervIndex -= gap
			}
			nums[pervIndex+gap] = current
		}
	}
}
