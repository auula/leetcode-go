package sort

// BubbleSort 比较相邻的元素，如果第一个比第二个大，就交换他们两个。
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
}
