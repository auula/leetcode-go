package search

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	t.Log(nums)
	t.Log("4 index is :", BinarySearch(nums, 4))
}
