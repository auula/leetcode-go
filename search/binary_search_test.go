package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	t.Log(nums)
	t.Log("4 index is :", BinarySearch(nums, 4))
}

func Test_binarySerach(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	t.Log(nums)
	t.Log("2 of index :", _binarySerach(nums, 2, 0, len(nums)))
}
