package sort

import (
	"testing"
)

// go test -v -run=TestMergeSort
func TestMergeSort(t *testing.T) {
	var nums []int = []int{11, 2, 66, 9, 12, 3, 6, 9}
	t.Log(nums)
	t.Log(MergeSort(nums))
}
