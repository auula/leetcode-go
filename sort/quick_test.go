package sort

import "testing"

func TestQuickSort(t *testing.T) {
	var nums []int = []int{11, 2, 66, 9, 12, 3, 6, 9}
	t.Log(nums)
	QuickSort(nums, 0, len(nums)-1)
	t.Log(nums)
}
